package controllers

import (
	"context"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
)

type ClientProfileController struct{}

func (c *ClientProfileController) Create(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	profile := obj.(swagger.MsgVpnClientProfile)
	_, _, err := client.ClientProfileApi.CreateMsgVpnClientProfile(ctx, profile, msgVpn, nil)
	return err
}

func (c *ClientProfileController) Get(ctx context.Context, client *swagger.APIClient, msgVpn string) ([]interface{}, error) {
	resp, _, err := client.ClientProfileApi.GetMsgVpnClientProfiles(ctx, msgVpn, nil)
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, len(resp.Data))
	for i, profile := range resp.Data {
		result[i] = profile
	}
	return result, nil
}

func (c *ClientProfileController) Update(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	profile := obj.(swagger.MsgVpnClientProfile)
	profile.MsgVpnName = msgVpn
	_, _, err := client.ClientProfileApi.UpdateMsgVpnClientProfile(ctx, profile, msgVpn, profile.ClientProfileName, nil)
	return err
}

func (c *ClientProfileController) Delete(ctx context.Context, client *swagger.APIClient, msgVpn string, identifier string) error {
	_, _, err := client.ClientProfileApi.DeleteMsgVpnClientProfile(ctx, msgVpn, identifier)
	return err
}

func (c *ClientProfileController) GetIdentifier(obj interface{}) string {
	return obj.(swagger.MsgVpnClientProfile).ClientProfileName
}

func (c *ClientProfileController) ShouldManage(obj interface{}) bool {
	if profile, ok := obj.(swagger.MsgVpnClientProfile); ok {
		return !isSystemResource(profile.ClientProfileName) && profile.ClientProfileName != "default"
	}
	return false
}

func (c *ClientProfileController) Equal(obj1, obj2 interface{}) bool {
	return !needsUpdate(obj1, obj2)
}
