package controllers

import (
	"context"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
)

type QueueController struct{}

func (c *QueueController) Create(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	queue := obj.(swagger.MsgVpnQueue)
	_, _, err := client.QueueApi.CreateMsgVpnQueue(ctx, queue, msgVpn, nil)
	return err
}

func (c *QueueController) Get(ctx context.Context, client *swagger.APIClient, msgVpn string) ([]interface{}, error) {
	resp, _, err := client.QueueApi.GetMsgVpnQueues(ctx, msgVpn, nil)
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, len(resp.Data))
	for i, queue := range resp.Data {
		result[i] = queue
	}
	return result, nil
}

func (c *QueueController) Update(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	queue := obj.(swagger.MsgVpnQueue)
	queue.MsgVpnName = msgVpn
	_, _, err := client.QueueApi.UpdateMsgVpnQueue(ctx, queue, msgVpn, queue.QueueName, nil)
	return err
}

func (c *QueueController) Delete(ctx context.Context, client *swagger.APIClient, msgVpn string, identifier string) error {
	_, _, err := client.QueueApi.DeleteMsgVpnQueue(ctx, msgVpn, identifier)
	return err
}

func (c *QueueController) GetIdentifier(obj interface{}) string {
	return obj.(swagger.MsgVpnQueue).QueueName
}

func (c *QueueController) ShouldManage(obj interface{}) bool {
	if queue, ok := obj.(swagger.MsgVpnQueue); ok {
		return !isSystemResource(queue.QueueName)
	}
	return false
}

func (c *QueueController) Equal(obj1, obj2 interface{}) bool {
	return !needsUpdate(obj1, obj2)
}
