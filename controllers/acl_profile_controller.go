package controllers

import (
	"context"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
)

type ACLProfileController struct{}

func (c *ACLProfileController) Create(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	profile := obj.(swagger.MsgVpnAclProfile)
	_, _, err := client.AclProfileApi.CreateMsgVpnAclProfile(ctx, profile, msgVpn, nil)
	return err
}

func (c *ACLProfileController) Get(ctx context.Context, client *swagger.APIClient, msgVpn string) ([]interface{}, error) {
	resp, _, err := client.AclProfileApi.GetMsgVpnAclProfiles(ctx, msgVpn, nil)
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, len(resp.Data))
	for i, profile := range resp.Data {
		result[i] = profile
	}
	return result, nil
}

func (c *ACLProfileController) Update(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	profile := obj.(swagger.MsgVpnAclProfile)
	profile.MsgVpnName = msgVpn
	_, _, err := client.AclProfileApi.UpdateMsgVpnAclProfile(ctx, profile, msgVpn, profile.AclProfileName, nil)
	return err
}

func (c *ACLProfileController) Delete(ctx context.Context, client *swagger.APIClient, msgVpn string, identifier string) error {
	_, _, err := client.AclProfileApi.DeleteMsgVpnAclProfile(ctx, msgVpn, identifier)
	return err
}

func (c *ACLProfileController) GetIdentifier(obj interface{}) string {
	return obj.(swagger.MsgVpnAclProfile).AclProfileName
}

func (c *ACLProfileController) ShouldManage(obj interface{}) bool {
	if profile, ok := obj.(swagger.MsgVpnAclProfile); ok {
		return !isSystemResource(profile.AclProfileName) && profile.AclProfileName != "default"
	}
	return false
}

func (c *ACLProfileController) Equal(obj1, obj2 interface{}) bool {
	return !needsUpdate(obj1, obj2)
}
