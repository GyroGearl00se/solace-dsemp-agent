package controllers

import (
	"context"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
)

type ProxyController struct{}

func (c *ProxyController) Create(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	proxy := obj.(swagger.MsgVpnProxy)
	_, _, err := client.ProxyApi.CreateMsgVpnProxy(ctx, proxy, msgVpn, nil)
	return err
}

func (c *ProxyController) Get(ctx context.Context, client *swagger.APIClient, msgVpn string) ([]interface{}, error) {
	resp, _, err := client.ProxyApi.GetMsgVpnProxies(ctx, msgVpn, nil)
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, len(resp.Data))
	for i, proxy := range resp.Data {
		result[i] = proxy
	}
	return result, nil
}

func (c *ProxyController) Update(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	proxy := obj.(swagger.MsgVpnProxy)
	_, _, err := client.ProxyApi.UpdateMsgVpnProxy(ctx, proxy, msgVpn, proxy.ProxyName, nil)
	return err
}

func (c *ProxyController) Delete(ctx context.Context, client *swagger.APIClient, msgVpn string, identifier string) error {
	_, _, err := client.ProxyApi.DeleteMsgVpnProxy(ctx, msgVpn, identifier)
	return err
}

func (c *ProxyController) GetIdentifier(obj interface{}) string {
	return obj.(swagger.MsgVpnProxy).ProxyName
}

func (c *ProxyController) ShouldManage(obj interface{}) bool {
	if proxy, ok := obj.(swagger.MsgVpnProxy); ok {
		return !isSystemResource(proxy.ProxyName)
	}
	return false
}

func (c *ProxyController) Equal(obj1, obj2 interface{}) bool {
	return !needsUpdate(obj1, obj2)
}
