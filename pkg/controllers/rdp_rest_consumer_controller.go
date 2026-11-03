package controllers

import (
	"context"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
)

type RDPRestConsumerController struct {
	RDPName string
}

func (c *RDPRestConsumerController) Create(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	consumer := obj.(swagger.MsgVpnRestDeliveryPointRestConsumer)
	_, _, err := client.RestDeliveryPointApi.CreateMsgVpnRestDeliveryPointRestConsumer(ctx, consumer, msgVpn, c.RDPName, nil)
	return err
}

func (c *RDPRestConsumerController) Get(ctx context.Context, client *swagger.APIClient, msgVpn string) ([]interface{}, error) {
	resp, _, err := client.RestDeliveryPointApi.GetMsgVpnRestDeliveryPointRestConsumers(ctx, msgVpn, c.RDPName, nil)
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, len(resp.Data))
	for i, consumer := range resp.Data {
		result[i] = consumer
	}
	return result, nil
}

func (c *RDPRestConsumerController) Update(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	consumer := obj.(swagger.MsgVpnRestDeliveryPointRestConsumer)
	_, _, err := client.RestDeliveryPointApi.UpdateMsgVpnRestDeliveryPointRestConsumer(ctx, consumer, msgVpn, c.RDPName, consumer.RestConsumerName, nil)
	return err
}

func (c *RDPRestConsumerController) Delete(ctx context.Context, client *swagger.APIClient, msgVpn string, identifier string) error {
	_, _, err := client.RestDeliveryPointApi.DeleteMsgVpnRestDeliveryPointRestConsumer(ctx, msgVpn, c.RDPName, identifier)
	return err
}

func (c *RDPRestConsumerController) GetIdentifier(obj interface{}) string {
	return obj.(swagger.MsgVpnRestDeliveryPointRestConsumer).RestConsumerName
}

func (c *RDPRestConsumerController) ShouldManage(obj interface{}) bool {
	return !isSystemResource(c.RDPName)
}

func (c *RDPRestConsumerController) Equal(obj1, obj2 interface{}) bool {
	return !needsUpdate(obj1, obj2)
}
