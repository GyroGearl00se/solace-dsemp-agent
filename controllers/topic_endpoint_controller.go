package controllers

import (
	"context"
	"net/url"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
	"github.com/antihax/optional"
)

type TopicEndpointController struct {
	WhitelistPatterns []string
}

func (c *TopicEndpointController) Create(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	endpoint := obj.(swagger.MsgVpnTopicEndpoint)
	_, _, err := client.TopicEndpointApi.CreateMsgVpnTopicEndpoint(ctx, endpoint, msgVpn, nil)
	return err
}

func (c *TopicEndpointController) Get(ctx context.Context, client *swagger.APIClient, msgVpn string) ([]interface{}, error) {
	var result []interface{}
	var cursor string

	for {
		opts := &swagger.TopicEndpointApiGetMsgVpnTopicEndpointsOpts{
			Count: optional.NewInt32(100),
		}
		if cursor != "" {
			opts.Cursor = optional.NewString(cursor)
		}

		resp, _, err := client.TopicEndpointApi.GetMsgVpnTopicEndpoints(ctx, msgVpn, opts)
		if err != nil {
			return nil, err
		}

		for _, endpoint := range resp.Data {
			result = append(result, endpoint)
		}

		if resp.Meta == nil || resp.Meta.Paging == nil || resp.Meta.Paging.CursorQuery == "" {
			break
		}
		cursor = resp.Meta.Paging.CursorQuery
	}

	return result, nil
}

func (c *TopicEndpointController) Update(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	endpoint := obj.(swagger.MsgVpnTopicEndpoint)
	encodedName := url.PathEscape(endpoint.TopicEndpointName)
	_, _, err := client.TopicEndpointApi.UpdateMsgVpnTopicEndpoint(ctx, endpoint, msgVpn, encodedName, nil)
	return err
}

func (c *TopicEndpointController) Delete(ctx context.Context, client *swagger.APIClient, msgVpn string, identifier string) error {
	encodedIdentifier := url.PathEscape(identifier)
	_, _, err := client.TopicEndpointApi.DeleteMsgVpnTopicEndpoint(ctx, msgVpn, encodedIdentifier)
	return err
}

func (c *TopicEndpointController) GetIdentifier(obj interface{}) string {
	return obj.(swagger.MsgVpnTopicEndpoint).TopicEndpointName
}

func (c *TopicEndpointController) ShouldManage(obj interface{}) bool {
	if endpoint, ok := obj.(swagger.MsgVpnTopicEndpoint); ok {
		if isSystemResource(endpoint.TopicEndpointName) || isWhitelisted(endpoint.TopicEndpointName, c.WhitelistPatterns) {
			return false
		}
		return true
	}
	return false
}

func (c *TopicEndpointController) Equal(obj1, obj2 interface{}) bool {
	return !needsUpdate(obj1, obj2)
}
