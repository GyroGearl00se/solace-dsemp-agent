package controllers

import (
	"context"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
)

type TopicEndpointTemplateController struct{}

func (c *TopicEndpointTemplateController) Create(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	template := obj.(swagger.MsgVpnTopicEndpointTemplate)
	_, _, err := client.TopicEndpointTemplateApi.CreateMsgVpnTopicEndpointTemplate(ctx, template, msgVpn, nil)
	return err
}

func (c *TopicEndpointTemplateController) Get(ctx context.Context, client *swagger.APIClient, msgVpn string) ([]interface{}, error) {
	resp, _, err := client.TopicEndpointTemplateApi.GetMsgVpnTopicEndpointTemplates(ctx, msgVpn, nil)
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, len(resp.Data))
	for i, template := range resp.Data {
		result[i] = template
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
