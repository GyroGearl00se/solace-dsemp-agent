package controllers

import (
	"context"
	"encoding/json"
	"fmt"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
	"github.com/sirupsen/logrus"
)

type RDPQueueBindingProtectedRequestHeaderController struct {
	RDPName          string
	QueueBindingName string
}

func (c *RDPQueueBindingProtectedRequestHeaderController) Create(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	header := obj.(swagger.MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader)
	if header.HeaderName == "" {
		return fmt.Errorf("protected request header missing headerName")
	}
	if header.HeaderValue == "" {
		return fmt.Errorf("protected request header missing headerValue")
	}

	header.RestDeliveryPointName = ""
	header.QueueBindingName = ""
	header.MsgVpnName = ""

	bodyBytes, _ := json.Marshal(header)
	logrus.WithFields(logrus.Fields{
		"category":          "RDP Queue Binding Protected Request Header",
		"msgVpn":            msgVpn,
		"restDeliveryPoint": c.RDPName,
		"queueBinding":      c.QueueBindingName,
		"payload":           string(bodyBytes),
	}).Debug("Creating protected request header payload")

	_, _, err := client.RestDeliveryPointApi.CreateMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader(ctx, header, msgVpn, c.RDPName, c.QueueBindingName, nil)
	if err != nil {
		fields := logrus.Fields{
			"category":          "RDP Queue Binding Protected Request Header",
			"msgVpn":            msgVpn,
			"restDeliveryPoint": c.RDPName,
			"queueBinding":      c.QueueBindingName,
			"headerName":        header.HeaderName,
		}
		if gse, ok := err.(swagger.GenericSwaggerError); ok {
			fields["responseBody"] = string(gse.Body())
		}
		logrus.WithFields(fields).Error("Protected request header create failed")
	}
	return err
}

func (c *RDPQueueBindingProtectedRequestHeaderController) Get(ctx context.Context, client *swagger.APIClient, msgVpn string) ([]interface{}, error) {
	resp, _, err := client.RestDeliveryPointApi.GetMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaders(ctx, msgVpn, c.RDPName, c.QueueBindingName, nil)
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, len(resp.Data))
	for i, header := range resp.Data {
		result[i] = header
	}
	return result, nil
}

func (c *RDPQueueBindingProtectedRequestHeaderController) Update(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error {
	header := obj.(swagger.MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader)
	header.RestDeliveryPointName = ""
	header.QueueBindingName = ""
	header.MsgVpnName = ""
	_, _, err := client.RestDeliveryPointApi.UpdateMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader(ctx, header, msgVpn, c.RDPName, c.QueueBindingName, header.HeaderName, nil)
	return err
}

func (c *RDPQueueBindingProtectedRequestHeaderController) Delete(ctx context.Context, client *swagger.APIClient, msgVpn string, identifier string) error {
	_, _, err := client.RestDeliveryPointApi.DeleteMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader(ctx, msgVpn, c.RDPName, c.QueueBindingName, identifier)
	return err
}

func (c *RDPQueueBindingProtectedRequestHeaderController) GetIdentifier(obj interface{}) string {
	return obj.(swagger.MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader).HeaderName
}

func (c *RDPQueueBindingProtectedRequestHeaderController) ShouldManage(obj interface{}) bool {
	return !isSystemResource(c.RDPName)
}

func (c *RDPQueueBindingProtectedRequestHeaderController) Equal(obj1, obj2 interface{}) bool {
	return !needsUpdate(obj1, obj2)
}
