package controllers

import (
	"context"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
)

type ClientUsernameController struct{}

func (c *ClientUsernameController) Create(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	username := obj.(swagger.MsgVpnClientUsername)
	_, _, err := client.ClientUsernameApi.CreateMsgVpnClientUsername(ctx, username, msgVpn, nil)
	return err
}

func (c *ClientUsernameController) Get(ctx context.Context, client *swagger.APIClient, msgVpn string) ([]interface{}, error) {
	resp, _, err := client.ClientUsernameApi.GetMsgVpnClientUsernames(ctx, msgVpn, nil)
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, len(resp.Data))
	for i, username := range resp.Data {
		result[i] = username
	}
	return result, nil
}

func (c *ClientUsernameController) Update(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	username := obj.(swagger.MsgVpnClientUsername)
	username.MsgVpnName = msgVpn
	_, _, err := client.ClientUsernameApi.UpdateMsgVpnClientUsername(ctx, username, msgVpn, username.ClientUsername, nil)
	return err
}

func (c *ClientUsernameController) Delete(ctx context.Context, client *swagger.APIClient, msgVpn string, identifier string) error {
	_, _, err := client.ClientUsernameApi.DeleteMsgVpnClientUsername(ctx, msgVpn, identifier)
	return err
}

func (c *ClientUsernameController) GetIdentifier(obj interface{}) string {
	return obj.(swagger.MsgVpnClientUsername).ClientUsername
}

func (c *ClientUsernameController) ShouldManage(obj interface{}) bool {
	if username, ok := obj.(swagger.MsgVpnClientUsername); ok {
		return !isSystemResource(username.ClientUsername) && username.ClientUsername != "default"
	}
	return false
}

func (c *ClientUsernameController) Equal(obj1, obj2 interface{}) bool {
	return !needsUpdate(obj1, obj2)
}
