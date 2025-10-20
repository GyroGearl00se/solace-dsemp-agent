package controllers

import (
	"context"
	"fmt"
	"strings"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
)

type ACLProfileExceptionController struct {
	IsPublishException bool
	AclProfileName     string
}

func (c *ACLProfileExceptionController) ShouldManage(obj interface{}) bool {
	return true
}

func (c *ACLProfileExceptionController) Equal(obj1, obj2 interface{}) bool {
	return !needsUpdate(obj1, obj2)
}

func (c *ACLProfileExceptionController) Create(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	if c.IsPublishException {
		exception := obj.(swagger.MsgVpnAclProfilePublishException)
		_, _, err := client.AclProfileApi.CreateMsgVpnAclProfilePublishException(ctx, exception, msgVpn, c.AclProfileName, nil)
		return err
	} else {
		exception := obj.(swagger.MsgVpnAclProfileSubscribeException)
		_, _, err := client.AclProfileApi.CreateMsgVpnAclProfileSubscribeException(ctx, exception, msgVpn, c.AclProfileName, nil)
		return err
	}
}

func (c *ACLProfileExceptionController) Get(ctx context.Context, client *swagger.APIClient, msgVpn string) ([]interface{}, error) {
	if c.IsPublishException {
		resp, _, err := client.AclProfileApi.GetMsgVpnAclProfilePublishExceptions(ctx, msgVpn, c.AclProfileName, nil)
		if err != nil {
			return nil, err
		}
		result := make([]interface{}, len(resp.Data))
		for i, exception := range resp.Data {
			result[i] = exception
		}
		return result, nil
	} else {
		resp, _, err := client.AclProfileApi.GetMsgVpnAclProfileSubscribeExceptions(ctx, msgVpn, c.AclProfileName, nil)
		if err != nil {
			return nil, err
		}
		result := make([]interface{}, len(resp.Data))
		for i, exception := range resp.Data {
			result[i] = exception
		}
		return result, nil
	}
}

func (c *ACLProfileExceptionController) Update(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	// Exceptions can't be updated, they need to be deleted and recreated
	return nil
}

func (c *ACLProfileExceptionController) Delete(ctx context.Context, client *swagger.APIClient, msgVpn string, identifier string) error {
	// Split identifier back into topicSyntax and topic
	parts := strings.SplitN(identifier, ":", 2)
	if len(parts) != 2 {
		return fmt.Errorf("invalid identifier format, expected 'topicSyntax:topic', got '%s'", identifier)
	}
	topicSyntax, topic := parts[0], parts[1]

	if c.IsPublishException {
		_, _, err := client.AclProfileApi.DeleteMsgVpnAclProfilePublishException(ctx, msgVpn, c.AclProfileName, topicSyntax, topic)
		return err
	} else {
		_, _, err := client.AclProfileApi.DeleteMsgVpnAclProfileSubscribeException(ctx, msgVpn, c.AclProfileName, topicSyntax, topic)
		return err
	}
}

func (c *ACLProfileExceptionController) GetIdentifier(obj interface{}) string {
	if c.IsPublishException {
		exception := obj.(swagger.MsgVpnAclProfilePublishException)
		return exception.TopicSyntax + ":" + exception.PublishExceptionTopic
	} else {
		exception := obj.(swagger.MsgVpnAclProfileSubscribeException)
		return exception.TopicSyntax + ":" + exception.SubscribeExceptionTopic
	}
}
