package controllers

import (
	"context"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
	"github.com/antihax/optional"
)

type RDPController struct{}

func (c *RDPController) Create(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	rdp := obj.(swagger.MsgVpnRestDeliveryPoint)
	_, _, err := client.RestDeliveryPointApi.CreateMsgVpnRestDeliveryPoint(ctx, rdp, msgVpn, nil)
	return err
}

func (c *RDPController) Get(ctx context.Context, client *swagger.APIClient, msgVpn string) ([]interface{}, error) {
	var result []interface{}
	var cursor string

	for {
		opts := &swagger.RestDeliveryPointApiGetMsgVpnRestDeliveryPointsOpts{
			Count: optional.NewInt32(100),
		}
		if cursor != "" {
			opts.Cursor = optional.NewString(cursor)
		}

		resp, _, err := client.RestDeliveryPointApi.GetMsgVpnRestDeliveryPoints(ctx, msgVpn, opts)
		if err != nil {
			return nil, err
		}

		for _, rdp := range resp.Data {
			result = append(result, rdp)
		}

		if resp.Meta == nil || resp.Meta.Paging == nil || resp.Meta.Paging.CursorQuery == "" {
			break
		}
		cursor = resp.Meta.Paging.CursorQuery
	}

	return result, nil
}

func (c *RDPController) Update(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	rdp := obj.(swagger.MsgVpnRestDeliveryPoint)
	_, _, err := client.RestDeliveryPointApi.UpdateMsgVpnRestDeliveryPoint(ctx, rdp, msgVpn, rdp.RestDeliveryPointName, nil)
	return err
}

func (c *RDPController) Delete(ctx context.Context, client *swagger.APIClient, msgVpn string, identifier string) error {
	_, _, err := client.RestDeliveryPointApi.DeleteMsgVpnRestDeliveryPoint(ctx, msgVpn, identifier)
	return err
}

func (c *RDPController) GetIdentifier(obj interface{}) string {
	return obj.(swagger.MsgVpnRestDeliveryPoint).RestDeliveryPointName
}

func (c *RDPController) ShouldManage(obj interface{}) bool {
	if rdp, ok := obj.(swagger.MsgVpnRestDeliveryPoint); ok {
		return !isSystemResource(rdp.RestDeliveryPointName)
	}
	return false
}

func (c *RDPController) Equal(obj1, obj2 interface{}) bool {
	return !needsUpdate(obj1, obj2)
}
