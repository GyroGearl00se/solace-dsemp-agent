package controllers

import (
	"context"
	"net/url"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
)

type JndiQueueController struct{}

func (c *JndiQueueController) Create(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	queue := obj.(swagger.MsgVpnJndiQueue)
	_, _, err := client.JndiApi.CreateMsgVpnJndiQueue(ctx, queue, msgVpn, nil)
	return err
}

func (c *JndiQueueController) Get(ctx context.Context, client *swagger.APIClient, msgVpn string) ([]interface{}, error) {
	resp, _, err := client.JndiApi.GetMsgVpnJndiQueues(ctx, msgVpn, nil)
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, len(resp.Data))
	for i, queue := range resp.Data {
		result[i] = queue
	}
	return result, nil
}

func (c *JndiQueueController) Update(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	queue := obj.(swagger.MsgVpnJndiQueue)
	encodedQueueName := url.PathEscape(queue.QueueName)
	_, _, err := client.JndiApi.UpdateMsgVpnJndiQueue(ctx, queue, msgVpn, encodedQueueName, nil)
	return err
}

func (c *JndiQueueController) Delete(ctx context.Context, client *swagger.APIClient, msgVpn string, identifier string) error {
	encodedIdentifier := url.PathEscape(identifier)
	_, _, err := client.JndiApi.DeleteMsgVpnJndiQueue(ctx, msgVpn, encodedIdentifier)
	return err
}

func (c *JndiQueueController) GetIdentifier(obj interface{}) string {
	return obj.(swagger.MsgVpnJndiQueue).QueueName
}

func (c *JndiQueueController) ShouldManage(obj interface{}) bool {
	if queue, ok := obj.(swagger.MsgVpnJndiQueue); ok {
		return !isSystemResource(queue.QueueName)
	}
	return false
}

func (c *JndiQueueController) Equal(obj1, obj2 interface{}) bool {
	return !needsUpdate(obj1, obj2)
}
