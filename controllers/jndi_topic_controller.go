package controllers

import (
	"context"
	"net/url"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
)

type JndiTopicController struct{}

func (c *JndiTopicController) Create(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	topic := obj.(swagger.MsgVpnJndiTopic)
	_, _, err := client.JndiApi.CreateMsgVpnJndiTopic(ctx, topic, msgVpn, nil)
	return err
}

func (c *JndiTopicController) Get(ctx context.Context, client *swagger.APIClient, msgVpn string) ([]interface{}, error) {
	resp, _, err := client.JndiApi.GetMsgVpnJndiTopics(ctx, msgVpn, nil)
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, len(resp.Data))
	for i, topic := range resp.Data {
		result[i] = topic
	}
	return result, nil
}

func (c *JndiTopicController) Update(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	topic := obj.(swagger.MsgVpnJndiTopic)
	encodedTopicName := url.PathEscape(topic.TopicName)
	_, _, err := client.JndiApi.UpdateMsgVpnJndiTopic(ctx, topic, msgVpn, encodedTopicName, nil)
	return err
}

func (c *JndiTopicController) Delete(ctx context.Context, client *swagger.APIClient, msgVpn string, identifier string) error {
	encodedIdentifier := url.PathEscape(identifier)
	_, _, err := client.JndiApi.DeleteMsgVpnJndiTopic(ctx, msgVpn, encodedIdentifier)
	return err
}

func (c *JndiTopicController) GetIdentifier(obj interface{}) string {
	return obj.(swagger.MsgVpnJndiTopic).TopicName
}

func (c *JndiTopicController) ShouldManage(obj interface{}) bool {
	if topic, ok := obj.(swagger.MsgVpnJndiTopic); ok {
		return !isSystemResource(topic.TopicName)
	}
	return false
}

func (c *JndiTopicController) Equal(obj1, obj2 interface{}) bool {
	return !needsUpdate(obj1, obj2)
}
