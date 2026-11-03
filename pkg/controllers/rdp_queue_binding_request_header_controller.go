package controllers

import (
	"context"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
)

type RDPQueueBindingRequestHeaderController struct {
	RDPName          string
	QueueBindingName string
}

func (c *RDPQueueBindingRequestHeaderController) Create(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	header := obj.(swagger.MsgVpnRestDeliveryPointQueueBindingRequestHeader)
	header.RestDeliveryPointName = ""
	header.QueueBindingName = ""
	header.MsgVpnName = ""
	_, _, err := client.RestDeliveryPointApi.CreateMsgVpnRestDeliveryPointQueueBindingRequestHeader(ctx, header, msgVpn, c.RDPName, c.QueueBindingName, nil)
	return err
}

func (c *RDPQueueBindingRequestHeaderController) Get(ctx context.Context, client *swagger.APIClient, msgVpn string) ([]interface{}, error) {
	resp, _, err := client.RestDeliveryPointApi.GetMsgVpnRestDeliveryPointQueueBindingRequestHeaders(ctx, msgVpn, c.RDPName, c.QueueBindingName, nil)
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, len(resp.Data))
	for i, header := range resp.Data {
		result[i] = header
	}
	return result, nil
}

func (c *RDPQueueBindingRequestHeaderController) Update(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	header := obj.(swagger.MsgVpnRestDeliveryPointQueueBindingRequestHeader)
	header.RestDeliveryPointName = ""
	header.QueueBindingName = ""
	header.MsgVpnName = ""
	_, _, err := client.RestDeliveryPointApi.UpdateMsgVpnRestDeliveryPointQueueBindingRequestHeader(ctx, header, msgVpn, c.RDPName, c.QueueBindingName, header.HeaderName, nil)
	return err
}

func (c *RDPQueueBindingRequestHeaderController) Delete(ctx context.Context, client *swagger.APIClient, msgVpn string, identifier string) error {
	_, _, err := client.RestDeliveryPointApi.DeleteMsgVpnRestDeliveryPointQueueBindingRequestHeader(ctx, msgVpn, c.RDPName, c.QueueBindingName, identifier)
	return err
}

func (c *RDPQueueBindingRequestHeaderController) GetIdentifier(obj interface{}) string {
	return obj.(swagger.MsgVpnRestDeliveryPointQueueBindingRequestHeader).HeaderName
}

func (c *RDPQueueBindingRequestHeaderController) ShouldManage(obj interface{}) bool {
	return !isSystemResource(c.RDPName)
}

func (c *RDPQueueBindingRequestHeaderController) Equal(obj1, obj2 interface{}) bool {
	return !needsUpdate(obj1, obj2)
}
