package controllers

import (
	"context"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
)

type JndiConnectionFactoryController struct{}

func (c *JndiConnectionFactoryController) Create(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	cf := obj.(swagger.MsgVpnJndiConnectionFactory)
	_, _, err := client.JndiApi.CreateMsgVpnJndiConnectionFactory(ctx, cf, msgVpn, nil)
	return err
}

func (c *JndiConnectionFactoryController) Get(ctx context.Context, client *swagger.APIClient, msgVpn string) ([]interface{}, error) {
	resp, _, err := client.JndiApi.GetMsgVpnJndiConnectionFactories(ctx, msgVpn, nil)
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, len(resp.Data))
	for i, cf := range resp.Data {
		result[i] = cf
	}
	return result, nil
}

func (c *JndiConnectionFactoryController) Update(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	cf := obj.(swagger.MsgVpnJndiConnectionFactory)
	_, _, err := client.JndiApi.UpdateMsgVpnJndiConnectionFactory(ctx, cf, msgVpn, cf.ConnectionFactoryName, nil)
	return err
}

func (c *JndiConnectionFactoryController) Delete(ctx context.Context, client *swagger.APIClient, msgVpn string, identifier string) error {
	_, _, err := client.JndiApi.DeleteMsgVpnJndiConnectionFactory(ctx, msgVpn, identifier)
	return err
}

func (c *JndiConnectionFactoryController) GetIdentifier(obj interface{}) string {
	return obj.(swagger.MsgVpnJndiConnectionFactory).ConnectionFactoryName
}

func (c *JndiConnectionFactoryController) ShouldManage(obj interface{}) bool {
	if cf, ok := obj.(swagger.MsgVpnJndiConnectionFactory); ok {
		return !isSystemResource(cf.ConnectionFactoryName) && cf.ConnectionFactoryName != "/jms/cf/default"
	}
	return false
}

func (c *JndiConnectionFactoryController) Equal(obj1, obj2 interface{}) bool {
	return !needsUpdate(obj1, obj2)
}
