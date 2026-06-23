package controllers

import (
	"context"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
)

type RDPRestConsumerTlsTrustedCommonNameController struct {
	RDPName          string
	RestConsumerName string
}

func (c *RDPRestConsumerTlsTrustedCommonNameController) Create(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	cn := obj.(swagger.MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName)
	cn.RestDeliveryPointName = ""
	cn.RestConsumerName = ""
	cn.MsgVpnName = ""
	_, _, err := client.RestDeliveryPointApi.CreateMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName(ctx, cn, msgVpn, c.RDPName, c.RestConsumerName, nil)
	return err
}

func (c *RDPRestConsumerTlsTrustedCommonNameController) Get(ctx context.Context, client *swagger.APIClient, msgVpn string) ([]interface{}, error) {
	resp, _, err := client.RestDeliveryPointApi.GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames(ctx, msgVpn, c.RDPName, c.RestConsumerName, nil)
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, len(resp.Data))
	for i, cn := range resp.Data {
		result[i] = cn
	}
	return result, nil
}

// Update is a no-op because TlsTrustedCommonName has no Update/Replace API.
// The only meaningful field is the name itself (the identifier), so changes require delete+create.
func (c *RDPRestConsumerTlsTrustedCommonNameController) Update(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	return nil
}

func (c *RDPRestConsumerTlsTrustedCommonNameController) Delete(ctx context.Context, client *swagger.APIClient, msgVpn string, identifier string) error {
	_, _, err := client.RestDeliveryPointApi.DeleteMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName(ctx, msgVpn, c.RDPName, c.RestConsumerName, identifier)
	return err
}

func (c *RDPRestConsumerTlsTrustedCommonNameController) GetIdentifier(obj interface{}) string {
	return obj.(swagger.MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName).TlsTrustedCommonName
}

func (c *RDPRestConsumerTlsTrustedCommonNameController) ShouldManage(obj interface{}) bool {
	return !isSystemResource(c.RDPName)
}

func (c *RDPRestConsumerTlsTrustedCommonNameController) Equal(obj1, obj2 interface{}) bool {
	return !needsUpdate(obj1, obj2)
}
