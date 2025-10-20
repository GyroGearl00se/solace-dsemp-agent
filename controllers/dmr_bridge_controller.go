package controllers

import (
	"context"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
)

type DmrBridgeController struct{}

func (c *DmrBridgeController) Create(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	bridge := obj.(swagger.MsgVpnDmrBridge)
	_, _, err := client.DmrBridgeApi.CreateMsgVpnDmrBridge(ctx, bridge, msgVpn, nil)
	return err
}

func (c *DmrBridgeController) Get(ctx context.Context, client *swagger.APIClient, msgVpn string) ([]interface{}, error) {
	resp, _, err := client.DmrBridgeApi.GetMsgVpnDmrBridges(ctx, msgVpn, nil)
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, len(resp.Data))
	for i, bridge := range resp.Data {
		result[i] = bridge
	}
	return result, nil
}

func (c *DmrBridgeController) Update(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	bridge := obj.(swagger.MsgVpnDmrBridge)
	bridge.MsgVpnName = msgVpn
	_, _, err := client.DmrBridgeApi.UpdateMsgVpnDmrBridge(ctx, bridge, msgVpn, bridge.RemoteNodeName, nil)
	return err
}

func (c *DmrBridgeController) Delete(ctx context.Context, client *swagger.APIClient, msgVpn string, identifier string) error {
	_, _, err := client.DmrBridgeApi.DeleteMsgVpnDmrBridge(ctx, msgVpn, identifier)
	return err
}

func (c *DmrBridgeController) GetIdentifier(obj interface{}) string {
	return obj.(swagger.MsgVpnDmrBridge).RemoteNodeName
}

func (c *DmrBridgeController) ShouldManage(obj interface{}) bool {
	if bridge, ok := obj.(swagger.MsgVpnDmrBridge); ok {
		return !isSystemResource(bridge.RemoteNodeName)
	}
	return false
}

func (c *DmrBridgeController) Equal(obj1, obj2 interface{}) bool {
	return !needsUpdate(obj1, obj2)
}
