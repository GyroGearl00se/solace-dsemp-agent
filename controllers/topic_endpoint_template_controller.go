package controllers

import (
	"context"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
	"github.com/antihax/optional"
)

type TopicEndpointTemplateController struct{}

func (c *TopicEndpointTemplateController) Create(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	template := obj.(swagger.MsgVpnTopicEndpointTemplate)
	_, _, err := client.TopicEndpointTemplateApi.CreateMsgVpnTopicEndpointTemplate(ctx, template, msgVpn, nil)
	return err
}

func (c *TopicEndpointTemplateController) Get(ctx context.Context, client *swagger.APIClient, msgVpn string) ([]interface{}, error) {
	var result []interface{}
	var cursor string

	for {
		opts := &swagger.TopicEndpointTemplateApiGetMsgVpnTopicEndpointTemplatesOpts{
			Count: optional.NewInt32(100),
		}
		if cursor != "" {
			opts.Cursor = optional.NewString(cursor)
		}

		resp, _, err := client.TopicEndpointTemplateApi.GetMsgVpnTopicEndpointTemplates(ctx, msgVpn, opts)
		if err != nil {
			return nil, err
		}

		for _, template := range resp.Data {
			result = append(result, template)
		}

		if resp.Meta == nil || resp.Meta.Paging == nil || resp.Meta.Paging.CursorQuery == "" {
			break
		}
		cursor = resp.Meta.Paging.CursorQuery
	}

	return result, nil
}

func (c *TopicEndpointTemplateController) Update(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	template := obj.(swagger.MsgVpnTopicEndpointTemplate)
	_, _, err := client.TopicEndpointTemplateApi.UpdateMsgVpnTopicEndpointTemplate(ctx, template, msgVpn, template.TopicEndpointTemplateName, nil)
	return err
}

func (c *TopicEndpointTemplateController) Delete(ctx context.Context, client *swagger.APIClient, msgVpn string, identifier string) error {
	_, _, err := client.TopicEndpointTemplateApi.DeleteMsgVpnTopicEndpointTemplate(ctx, msgVpn, identifier)
	return err
}

func (c *TopicEndpointTemplateController) GetIdentifier(obj interface{}) string {
	return obj.(swagger.MsgVpnTopicEndpointTemplate).TopicEndpointTemplateName
}

func (c *TopicEndpointTemplateController) ShouldManage(obj interface{}) bool {
	if template, ok := obj.(swagger.MsgVpnTopicEndpointTemplate); ok {
		return !isSystemResource(template.TopicEndpointTemplateName)
	}
	return false
}

func (c *TopicEndpointTemplateController) Equal(obj1, obj2 interface{}) bool {
	return !needsUpdate(obj1, obj2)
}
