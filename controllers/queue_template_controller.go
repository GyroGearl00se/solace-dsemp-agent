package controllers

import (
	"context"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
)

type QueueTemplateController struct{}

func (c *QueueTemplateController) Create(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	template := obj.(swagger.MsgVpnQueueTemplate)
	_, _, err := client.QueueTemplateApi.CreateMsgVpnQueueTemplate(ctx, template, msgVpn, nil)
	return err
}

func (c *QueueTemplateController) Get(ctx context.Context, client *swagger.APIClient, msgVpn string) ([]interface{}, error) {
	resp, _, err := client.QueueTemplateApi.GetMsgVpnQueueTemplates(ctx, msgVpn, nil)
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, len(resp.Data))
	for i, template := range resp.Data {
		result[i] = template
	}
	return result, nil
}

func (c *QueueTemplateController) Update(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	template := obj.(swagger.MsgVpnQueueTemplate)
	_, _, err := client.QueueTemplateApi.UpdateMsgVpnQueueTemplate(ctx, template, msgVpn, template.QueueTemplateName, nil)
	return err
}

func (c *QueueTemplateController) Delete(ctx context.Context, client *swagger.APIClient, msgVpn string, identifier string) error {
	_, _, err := client.QueueTemplateApi.DeleteMsgVpnQueueTemplate(ctx, msgVpn, identifier)
	return err
}

func (c *QueueTemplateController) GetIdentifier(obj interface{}) string {
	return obj.(swagger.MsgVpnQueueTemplate).QueueTemplateName
}

func (c *QueueTemplateController) ShouldManage(obj interface{}) bool {
	if template, ok := obj.(swagger.MsgVpnQueueTemplate); ok {
		return !isSystemResource(template.QueueTemplateName)
	}
	return false
}

func (c *QueueTemplateController) Equal(obj1, obj2 interface{}) bool {
	return !needsUpdate(obj1, obj2)
}
