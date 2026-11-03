package controllers

import (
	"context"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
	"github.com/antihax/optional"
)

type DmrBridgeController struct{}

func (c *DmrBridgeController) Create(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	bridge := obj.(swagger.MsgVpnDmrBridge)
	_, _, err := client.DmrBridgeApi.CreateMsgVpnDmrBridge(ctx, bridge, msgVpn, nil)
	return err
}

func (c *DmrBridgeController) Get(ctx context.Context, client *swagger.APIClient, msgVpn string) ([]interface{}, error) {
	var result []interface{}
	var cursor string

	for {
		opts := &swagger.DmrBridgeApiGetMsgVpnDmrBridgesOpts{
			Count: optional.NewInt32(100),
		}
		if cursor != "" {
			opts.Cursor = optional.NewString(cursor)
		}

		resp, _, err := client.DmrBridgeApi.GetMsgVpnDmrBridges(ctx, msgVpn, opts)
		if err != nil {
			return nil, err
		}

		for _, bridge := range resp.Data {
			result = append(result, bridge)
		}

		if resp.Meta == nil || resp.Meta.Paging == nil || resp.Meta.Paging.CursorQuery == "" {
			break
		}
		cursor = resp.Meta.Paging.CursorQuery
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
