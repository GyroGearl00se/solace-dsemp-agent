package controllers

import (
	"context"
	"net/url"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
)

type QueueSubscriptionController struct {
	QueueName string // Parent queue name
}

func (c *QueueSubscriptionController) Create(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	sub := obj.(swagger.MsgVpnQueueSubscription)
	_, _, err := client.QueueApi.CreateMsgVpnQueueSubscription(ctx, sub, msgVpn, c.QueueName, nil)
	return err
}

func (c *QueueSubscriptionController) Get(ctx context.Context, client *swagger.APIClient, msgVpn string) ([]interface{}, error) {
	resp, _, err := client.QueueApi.GetMsgVpnQueueSubscriptions(ctx, msgVpn, c.QueueName, nil)
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, len(resp.Data))
	for i, sub := range resp.Data {
		result[i] = sub
	}
	return result, nil
}

func (c *QueueSubscriptionController) Update(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	// Subscriptions can't be updated, they need to be deleted and recreated
	return nil
}

func (c *QueueSubscriptionController) Delete(ctx context.Context, client *swagger.APIClient, msgVpn string, identifier string) error {
	_, _, err := client.QueueApi.DeleteMsgVpnQueueSubscription(ctx, msgVpn, c.QueueName, identifier)
	return err
}

func (c *QueueSubscriptionController) GetIdentifier(obj interface{}) string {
	return url.QueryEscape(obj.(swagger.MsgVpnQueueSubscription).SubscriptionTopic)
}

func (c *QueueSubscriptionController) ShouldManage(obj interface{}) bool {
	return true
}

func (c *QueueSubscriptionController) Equal(obj1, obj2 interface{}) bool {
	return !needsUpdate(obj1, obj2)
}
