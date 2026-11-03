package controllers

import (
	"context"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
)

type RDPQueueBindingController struct {
	RDPName string
}

func (c *RDPQueueBindingController) Create(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	binding := obj.(swagger.MsgVpnRestDeliveryPointQueueBinding)
	binding.RestDeliveryPointName = ""
	binding.MsgVpnName = ""
	_, _, err := client.RestDeliveryPointApi.CreateMsgVpnRestDeliveryPointQueueBinding(ctx, binding, msgVpn, c.RDPName, nil)
	return err
}

func (c *RDPQueueBindingController) Get(ctx context.Context, client *swagger.APIClient, msgVpn string) ([]interface{}, error) {
	resp, _, err := client.RestDeliveryPointApi.GetMsgVpnRestDeliveryPointQueueBindings(ctx, msgVpn, c.RDPName, nil)
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, len(resp.Data))
	for i, rdp := range resp.Data {
		result[i] = rdp
	}
	return result, nil
}

func (c *RDPQueueBindingController) Update(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	binding := obj.(swagger.MsgVpnRestDeliveryPointQueueBinding)
	binding.RestDeliveryPointName = ""
	binding.MsgVpnName = ""
	_, _, err := client.RestDeliveryPointApi.UpdateMsgVpnRestDeliveryPointQueueBinding(ctx, binding, msgVpn, c.RDPName, binding.QueueBindingName, nil)
	return err
}

func (c *RDPQueueBindingController) Delete(ctx context.Context, client *swagger.APIClient, msgVpn string, identifier string) error {
	_, _, err := client.RestDeliveryPointApi.DeleteMsgVpnRestDeliveryPointQueueBinding(ctx, msgVpn, c.RDPName, identifier)
	return err
}

func (c *RDPQueueBindingController) GetIdentifier(obj interface{}) string {
	return obj.(swagger.MsgVpnRestDeliveryPointQueueBinding).QueueBindingName
}

func (c *RDPQueueBindingController) ShouldManage(obj interface{}) bool {
	return !isSystemResource(c.RDPName)
}

func (c *RDPQueueBindingController) Equal(obj1, obj2 interface{}) bool {
	return !needsUpdate(obj1, obj2)
}
