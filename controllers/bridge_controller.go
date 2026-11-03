package controllers

import (
	"context"
	"fmt"
	"strings"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
	"github.com/antihax/optional"
)

type BridgeController struct{}

func (c *BridgeController) Create(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	bridge := obj.(swagger.MsgVpnBridge)
	_, _, err := client.BridgeApi.CreateMsgVpnBridge(ctx, bridge, msgVpn, nil)
	return err
}

func (c *BridgeController) Get(ctx context.Context, client *swagger.APIClient, msgVpn string) ([]interface{}, error) {
	var result []interface{}
	var cursor string

	for {
		opts := &swagger.BridgeApiGetMsgVpnBridgesOpts{
			Count: optional.NewInt32(100),
		}
		if cursor != "" {
			opts.Cursor = optional.NewString(cursor)
		}

		resp, _, err := client.BridgeApi.GetMsgVpnBridges(ctx, msgVpn, opts)
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

func (c *BridgeController) Update(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	bridge := obj.(swagger.MsgVpnBridge)
	bridge.MsgVpnName = msgVpn
	_, _, err := client.BridgeApi.UpdateMsgVpnBridge(ctx, bridge, msgVpn, bridge.BridgeName, bridge.BridgeVirtualRouter, nil)
	return err
}

func (c *BridgeController) Delete(ctx context.Context, client *swagger.APIClient, msgVpn string, identifier string) error {
	parts := strings.SplitN(identifier, "/", 2)
	if len(parts) != 2 {
		return fmt.Errorf("invalid bridge identifier format, expected 'bridgeName/virtualRouter', got '%s'", identifier)
	}
	bridgeName, virtualRouter := parts[0], parts[1]

	_, _, err := client.BridgeApi.DeleteMsgVpnBridge(ctx, msgVpn, bridgeName, virtualRouter)
	return err
}

func (c *BridgeController) GetIdentifier(obj interface{}) string {
	bridge := obj.(swagger.MsgVpnBridge)
	return bridge.BridgeName + "/" + bridge.BridgeVirtualRouter
}

func (c *BridgeController) ShouldManage(obj interface{}) bool {
	if bridge, ok := obj.(swagger.MsgVpnBridge); ok {
		return !isSystemResource(bridge.BridgeName)
	}
	return false
}

func (c *BridgeController) Equal(obj1, obj2 interface{}) bool {
	return !needsUpdate(obj1, obj2)
}
