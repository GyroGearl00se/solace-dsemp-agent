package controllers

import (
	"context"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
	"github.com/antihax/optional"
)

type ACLProfileController struct{}

func (c *ACLProfileController) Create(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	profile := obj.(swagger.MsgVpnAclProfile)
	_, _, err := client.AclProfileApi.CreateMsgVpnAclProfile(ctx, profile, msgVpn, nil)
	return err
}

func (c *ACLProfileController) Get(ctx context.Context, client *swagger.APIClient, msgVpn string) ([]interface{}, error) {
	var result []interface{}
	var cursor string

	for {
		opts := &swagger.AclProfileApiGetMsgVpnAclProfilesOpts{
			Count: optional.NewInt32(100),
		}
		if cursor != "" {
			opts.Cursor = optional.NewString(cursor)
		}

		resp, _, err := client.AclProfileApi.GetMsgVpnAclProfiles(ctx, msgVpn, opts)
		if err != nil {
			return nil, err
		}

		for _, profile := range resp.Data {
			result = append(result, profile)
		}

		if resp.Meta == nil || resp.Meta.Paging == nil || resp.Meta.Paging.CursorQuery == "" {
			break
		}
		cursor = resp.Meta.Paging.CursorQuery
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
