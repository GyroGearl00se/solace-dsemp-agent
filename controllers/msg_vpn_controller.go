package controllers

import (
	"context"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
)

type MsgVpnController struct{}

func (c *MsgVpnController) Create(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	vpn := obj.(swagger.MsgVpn)
	_, _, err := client.MsgVpnApi.CreateMsgVpn(ctx, vpn, nil)
	return err
}

func (c *MsgVpnController) Get(ctx context.Context, client *swagger.APIClient, msgVpn string) ([]interface{}, error) {
	resp, _, err := client.MsgVpnApi.GetMsgVpns(ctx, nil)
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, len(resp.Data))
	for i, vpn := range resp.Data {
		result[i] = vpn
	}
	return result, nil
}

func (c *MsgVpnController) Update(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	vpn := obj.(swagger.MsgVpn)
	_, _, err := client.MsgVpnApi.UpdateMsgVpn(ctx, vpn, vpn.MsgVpnName, nil)
	return err
}

func (c *MsgVpnController) Delete(ctx context.Context, client *swagger.APIClient, msgVpn string, identifier string) error {
	_, _, err := client.MsgVpnApi.DeleteMsgVpn(ctx, identifier)
	return err
}

func (c *MsgVpnController) GetIdentifier(obj interface{}) string {
	return obj.(swagger.MsgVpn).MsgVpnName
}

func (c *MsgVpnController) ShouldManage(obj interface{}) bool {
	if vpn, ok := obj.(swagger.MsgVpn); ok {
		return !isSystemResource(vpn.MsgVpnName)
	}
	return false
}

func (c *MsgVpnController) Equal(obj1, obj2 interface{}) bool {
	return !needsUpdate(obj1, obj2)
}
