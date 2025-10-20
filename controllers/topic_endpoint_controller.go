package controllers

import (
	"context"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
)

type TopicEndpointController struct{}

func (c *TopicEndpointController) Create(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	endpoint := obj.(swagger.MsgVpnTopicEndpoint)
	_, _, err := client.TopicEndpointApi.CreateMsgVpnTopicEndpoint(ctx, endpoint, msgVpn, nil)
	return err
}

func (c *TopicEndpointController) Get(ctx context.Context, client *swagger.APIClient, msgVpn string) ([]interface{}, error) {
	resp, _, err := client.TopicEndpointApi.GetMsgVpnTopicEndpoints(ctx, msgVpn, nil)
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, len(resp.Data))
	for i, endpoint := range resp.Data {
		result[i] = endpoint
	}
	return result, nil
}

func (c *TopicEndpointController) Update(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	endpoint := obj.(swagger.MsgVpnTopicEndpoint)
	_, _, err := client.TopicEndpointApi.UpdateMsgVpnTopicEndpoint(ctx, endpoint, msgVpn, endpoint.TopicEndpointName, nil)
	return err
}

func (c *TopicEndpointController) Delete(ctx context.Context, client *swagger.APIClient, msgVpn string, identifier string) error {
	_, _, err := client.TopicEndpointApi.DeleteMsgVpnTopicEndpoint(ctx, msgVpn, identifier)
	return err
}

func (c *TopicEndpointController) GetIdentifier(obj interface{}) string {
	return obj.(swagger.MsgVpnTopicEndpoint).TopicEndpointName
}

func (c *TopicEndpointController) ShouldManage(obj interface{}) bool {
	if endpoint, ok := obj.(swagger.MsgVpnTopicEndpoint); ok {
		return !isSystemResource(endpoint.TopicEndpointName)
	}
	return false
}

func (c *TopicEndpointController) Equal(obj1, obj2 interface{}) bool {
	return !needsUpdate(obj1, obj2)
}
