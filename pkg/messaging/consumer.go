package messaging

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
	"solace.dev/go/messaging"
	"solace.dev/go/messaging/pkg/solace"
	"solace.dev/go/messaging/pkg/solace/config"
	"solace.dev/go/messaging/pkg/solace/message"
	"solace.dev/go/messaging/pkg/solace/resource"

	"github.com/GyroGearl00se/solace-dsemp-agent/pkg/state"
)

// ConsumerConfig specifies configuration for the Solace topic state consumer.
type ConsumerConfig struct {
	URL            string
	VPN            string
	Username       string
	Password       string
	Topic          string
	TrustStorePath string
	ValidateCert   bool
	AESKey         string
}

// Consumer subscribes to a Solace topic and decodes inbound TargetState payloads.
type Consumer struct {
	config    ConsumerConfig
	messaging solace.MessagingService
	receiver  solace.DirectMessageReceiver
	callback  func(*state.TargetState)
}

// NewConsumer creates a new Solace topic state consumer.
func NewConsumer(cfg ConsumerConfig) *Consumer {
	return &Consumer{config: cfg}
}

// OnStateReceived registers a handler callback invoked upon inbound target states.
func (c *Consumer) OnStateReceived(cb func(*state.TargetState)) {
	c.callback = cb
}

const logCategory = "Solace state consumer"

// Connect connects to the broker and starts the direct receiver.
func (c *Consumer) Connect() error {
	path := c.config.TrustStorePath
	if c.config.ValidateCert && c.config.TrustStorePath != "" {
		fileInfo, err := os.Stat(c.config.TrustStorePath)
		if err != nil {
			return fmt.Errorf("certificate path does not exist: %w", err)
		}
		if !fileInfo.IsDir() {
			path = filepath.Dir(path)
		}
	}

	brokerConfig := config.ServicePropertyMap{
		config.TransportLayerPropertyHost:                             c.config.URL,
		config.ServicePropertyVPNName:                                 c.config.VPN,
		config.AuthenticationPropertySchemeBasicPassword:              c.config.Password,
		config.AuthenticationPropertySchemeBasicUserName:              c.config.Username,
		config.TransportLayerPropertyReconnectionAttempts:             -1,
		config.TransportLayerPropertyReconnectionAttemptsWaitInterval: 500,
		config.ClientPropertyName:                                     "solace-dsemp-agent",
	}

	transportSecurityStrategy := config.NewTransportSecurityStrategy().
		WithExcludedProtocols(
			config.TransportSecurityProtocolTLSv1,
			config.TransportSecurityProtocolTLSv1_1,
			config.TransportSecurityProtocolSSLv3)

	if c.config.ValidateCert {
		if c.config.TrustStorePath != "" {
			brokerConfig[config.TransportLayerSecurityPropertyTrustStorePath] = path
			transportSecurityStrategy = transportSecurityStrategy.
				WithCertificateValidation(true, true, path, "")
		}
	} else {
		transportSecurityStrategy = transportSecurityStrategy.
			WithoutCertificateValidation()
	}

	messagingService, err := messaging.NewMessagingServiceBuilder().
		FromConfigurationProvider(brokerConfig).
		WithTransportSecurityStrategy(transportSecurityStrategy).
		Build()
	if err != nil {
		return fmt.Errorf("failed to build messaging service: %w", err)
	}

	c.messaging = messagingService

	messagingService.AddReconnectionAttemptListener(func(e solace.ServiceEvent) {
		logrus.WithField("category", logCategory).Infof("isConnected: %v", messagingService.IsConnected())
		if err := e.GetCause(); err != nil {
			logrus.WithField("category", logCategory).Errorf("Reconnection attempt error: %v", err)
		}
	})
	messagingService.AddReconnectionListener(func(e solace.ServiceEvent) {
		if err := e.GetCause(); err != nil {
			logrus.WithField("category", logCategory).Infof("Reconnection successful after error: %v", err)
		}
	})

	if err := messagingService.Connect(); err != nil {
		return fmt.Errorf("failed to connect to broker: %w", err)
	}

	for !messagingService.IsConnected() {
		logrus.WithField("category", logCategory).Info("Currently not connected, attempting to reconnect...")
		time.Sleep(500 * time.Millisecond)
	}
	logrus.WithField("category", logCategory).Info("Connected to broker")

	topicSub := resource.TopicSubscriptionOf(c.config.Topic)
	logrus.WithField("category", logCategory).Infof("Subscribing to: %s", c.config.Topic)

	directReceiver, err := messagingService.CreateDirectMessageReceiverBuilder().
		WithSubscriptions(topicSub).
		Build()
	if err != nil {
		c.messaging.Disconnect()
		return fmt.Errorf("failed to create message receiver: %w", err)
	}

	c.receiver = directReceiver

	if err := directReceiver.Start(); err != nil {
		c.messaging.Disconnect()
		return fmt.Errorf("failed to start message receiver: %w", err)
	}

	logrus.WithField("category", logCategory).Info("Message receiver started")

	if err := directReceiver.ReceiveAsync(c.handleMessage); err != nil {
		c.messaging.Disconnect()
		return fmt.Errorf("failed to register message handler: %w", err)
	}

	return nil
}

// Disconnect terminates receiver and broker connections gracefully.
func (c *Consumer) Disconnect() error {
	if c.receiver != nil {
		c.receiver.Terminate(10)
	}
	if c.messaging != nil {
		c.messaging.Disconnect()
	}
	return nil
}

func (c *Consumer) handleMessage(msg message.InboundMessage) {
	var payload []byte
	if payloadStr, ok := msg.GetPayloadAsString(); ok {
		payload = []byte(payloadStr)
	} else if payloadBytes, ok := msg.GetPayloadAsBytes(); ok {
		payload = payloadBytes
	} else {
		logrus.WithField("category", logCategory).Error("Failed to get payload as string or bytes")
		return
	}

	st, err := state.ParseTargetState(payload, c.config.AESKey)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"category": logCategory,
			"error":    err,
			"payload":  string(payload),
		}).Error("Failed to parse inbound target state payload")
		return
	}

	if c.callback != nil {
		c.callback(st)
	}
}

