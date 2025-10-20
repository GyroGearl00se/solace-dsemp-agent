package controllers

import (
	"context"
	"fmt"
	"strings"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
)

type BridgeRemoteMsgVpnController struct{}

func (c *BridgeRemoteMsgVpnController) Create(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	remoteMsgVpn := obj.(swagger.MsgVpnBridgeRemoteMsgVpn)
	_, _, err := client.BridgeApi.CreateMsgVpnBridgeRemoteMsgVpn(ctx, remoteMsgVpn, msgVpn, remoteMsgVpn.BridgeName, remoteMsgVpn.BridgeVirtualRouter, nil)
	return err
}

func (c *BridgeRemoteMsgVpnController) Get(ctx context.Context, client *swagger.APIClient, msgVpn string) ([]interface{}, error) {
	bridges, _, err := client.BridgeApi.GetMsgVpnBridges(ctx, msgVpn, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get bridges: %v", err)
	}

	var result []interface{}

	for _, bridge := range bridges.Data {
		resp, _, err := client.BridgeApi.GetMsgVpnBridgeRemoteMsgVpns(ctx, msgVpn, bridge.BridgeName, bridge.BridgeVirtualRouter, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to get remote VPNs for bridge %s/%s: %v", bridge.BridgeName, bridge.BridgeVirtualRouter, err)
		}
		for _, rmv := range resp.Data {
			result = append(result, rmv)
		}
	}
	return result, nil
}

func (c *BridgeRemoteMsgVpnController) Update(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	remoteMsgVpn := obj.(swagger.MsgVpnBridgeRemoteMsgVpn)
	_, _, err := client.BridgeApi.UpdateMsgVpnBridgeRemoteMsgVpn(
		ctx,
		remoteMsgVpn,
		msgVpn,
		remoteMsgVpn.BridgeName,
		remoteMsgVpn.BridgeVirtualRouter,
		remoteMsgVpn.RemoteMsgVpnName,
		remoteMsgVpn.RemoteMsgVpnLocation,
		remoteMsgVpn.RemoteMsgVpnInterface,
		nil,
	)
	return err
}

func (c *BridgeRemoteMsgVpnController) Delete(ctx context.Context, client *swagger.APIClient, msgVpn string, identifier string) error {
	// Identifier format: bridgeName/virtualRouter/location/interface/remoteMsgVpnName
	parts := strings.Split(identifier, "/")
	if len(parts) != 5 {
		return fmt.Errorf("invalid identifier format for Bridge Remote Message VPN, expected 5 parts but got %d", len(parts))
	}
	_, _, err := client.BridgeApi.DeleteMsgVpnBridgeRemoteMsgVpn(
		ctx,
		msgVpn,
		parts[0], // bridgeName
		parts[1], // bridgeVirtualRouter
		parts[2], // remoteMsgVpnLocation
		parts[3], // remoteMsgVpnInterface
		parts[4], // remoteMsgVpnName
	)
	return err
}

func (c *BridgeRemoteMsgVpnController) GetIdentifier(obj interface{}) string {
	rmv := obj.(swagger.MsgVpnBridgeRemoteMsgVpn)
	return fmt.Sprintf("%s/%s/%s/%s/%s",
		rmv.BridgeName,
		rmv.BridgeVirtualRouter,
		rmv.RemoteMsgVpnLocation,
		rmv.RemoteMsgVpnInterface,
		rmv.RemoteMsgVpnName,
	)
}

func (c *BridgeRemoteMsgVpnController) ShouldManage(obj interface{}) bool {
	return true
}

func (c *BridgeRemoteMsgVpnController) Equal(obj1, obj2 interface{}) bool {
	return !needsUpdate(obj1, obj2)
}
