package config

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"time"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
	"github.com/ghodss/yaml"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"solace.dev/go/messaging"
	"solace.dev/go/messaging/pkg/solace"
	"solace.dev/go/messaging/pkg/solace/config"
	"solace.dev/go/messaging/pkg/solace/message"
	"solace.dev/go/messaging/pkg/solace/resource"
)

type TargetState struct {
	MsgVpns                 []swagger.MsgVpn                      `json:"msgVpns"`
	Queues                  []QueueWithSubscriptions              `json:"queues"`
	ACLProfiles             []ACLProfileWithExceptions            `json:"aclProfiles"`
	ClientUsernames         []swagger.MsgVpnClientUsername        `json:"clientUsernames"`
	ClientProfiles          []swagger.MsgVpnClientProfile         `json:"clientProfiles"`
	Bridges                 []swagger.MsgVpnBridge                `json:"bridges"`
	BridgeRemoteMsgVpns     []swagger.MsgVpnBridgeRemoteMsgVpn    `json:"bridgeRemoteMsgVpns"`
	DmrBridges              []swagger.MsgVpnDmrBridge             `json:"dmrBridges"`
	JndiConnectionFactories []swagger.MsgVpnJndiConnectionFactory `json:"jndiConnectionFactories"`
	JndiQueues              []swagger.MsgVpnJndiQueue             `json:"jndiQueues"`
	JndiTopics              []swagger.MsgVpnJndiTopic             `json:"jndiTopics"`
	Proxies                 []swagger.MsgVpnProxy                 `json:"proxies"`
	RestDeliveryPoints      []RestDeliveryPoints                  `json:"restDeliveryPoints"`
	QueueTemplates          []swagger.MsgVpnQueueTemplate         `json:"queueTemplates"`
	TopicEndpoints          []swagger.MsgVpnTopicEndpoint         `json:"topicEndpoints"`
	TopicEndpointsTemplates []swagger.MsgVpnTopicEndpointTemplate `json:"topicEndpointsTemplates"`
	Version                 string                                `json:"version"`
	WhitelistPatterns       WhitelistPatterns                     `json:"whitelistPatterns"`
}

// QueueWithSubscriptions is a custom struct that combines an Queue and its subscriptions.
type QueueWithSubscriptions struct {
	swagger.MsgVpnQueue
	QueueSubscriptions []string `json:"queueSubscriptions,omitempty"`
}

type RestDeliveryPoints struct {
	swagger.MsgVpnRestDeliveryPoint
	RestDeliveryPointQueueBinding `json:"queueBinding,omitempty"`
	RestDeliveryPointRestConsumer `json:"restConsumer,omitempty"`
}

type RestDeliveryPointQueueBinding struct {
	swagger.MsgVpnRestDeliveryPointQueueBinding
	RequestHeaders          []swagger.MsgVpnRestDeliveryPointQueueBindingRequestHeader          `json:"requestHeaders,omitempty"`
	ProtectedRequestHeaders []swagger.MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader `json:"protectedRequestHeaders,omitempty"`
}

type RestDeliveryPointRestConsumer struct {
	swagger.MsgVpnRestDeliveryPointRestConsumer
	TlsTrustedCommonNames []swagger.MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName `json:"tlsTrustedCommonNames,omitempty"`
}

// PublishException represents a publish exception without the parent ACL profile name.
type PublishException struct {
	PublishExceptionTopic string `json:"publishExceptionTopic"`
	TopicSyntax           string `json:"topicSyntax"`
}

// SubscribeException represents a subscribe exception without the parent ACL profile name.
type SubscribeException struct {
	SubscribeExceptionTopic string `json:"subscribeExceptionTopic"`
	TopicSyntax             string `json:"topicSyntax"`
}

// ACLProfileWithExceptions is a custom struct that combines an ACL Profile and its exceptions.
type ACLProfileWithExceptions struct {
	swagger.MsgVpnAclProfile
	PublishExceptions   []PublishException   `json:"publishExceptions,omitempty"`
	SubscribeExceptions []SubscribeException `json:"subscribeExceptions,omitempty"`
}

// SolaceConfig stores the configuration for the Solace messaging client
type SolaceConfig struct {
	URL            string
	VPN            string
	Username       string
	Password       string
	Topic          string
	TrustStorePath string
	ValidateCert   bool
}

// WhitelistPatterns allows specifying patterns for whitelisting resources
type WhitelistPatterns struct {
	QueuePatterns         []string `json:"queuePatterns"`
	TopicEndpointPatterns []string `json:"topicEndpointPatterns"`
}

// NewDefaultSolaceConfig creates a new SolaceConfig with default values
func NewDefaultSolaceConfig() *SolaceConfig {
	return &SolaceConfig{}
}

// SolaceClient represents a client that connects to Solace
type SolaceClient struct {
	config          *SolaceConfig
	messaging       solace.MessagingService
	receiver        solace.DirectMessageReceiver
	messageCallback func(*TargetState)
}

// NewSolaceClient creates a new SolaceClient
func NewSolaceClient(config *SolaceConfig) *SolaceClient {
	return &SolaceClient{
		config: config,
	}
}

const category = "Solace state consumer"

// Connect establishes a connection to Solace broker
func (sc *SolaceClient) Connect() error {
	path := sc.config.TrustStorePath
	if sc.config.ValidateCert && sc.config.TrustStorePath != "" {
		fileInfo, err := os.Stat(sc.config.TrustStorePath)
		if err != nil {
			return fmt.Errorf("certificate path does not exist: %w", err)
		}
		if !fileInfo.IsDir() {
			path = filepath.Dir(path)
		} else {
			path = sc.config.TrustStorePath
		}
	}
	brokerConfig := config.ServicePropertyMap{
		config.TransportLayerPropertyHost:                             sc.config.URL,
		config.ServicePropertyVPNName:                                 sc.config.VPN,
		config.AuthenticationPropertySchemeBasicPassword:              sc.config.Password,
		config.AuthenticationPropertySchemeBasicUserName:              sc.config.Username,
		config.TransportLayerPropertyReconnectionAttempts:             -1,
		config.TransportLayerPropertyReconnectionAttemptsWaitInterval: 500,
		config.ClientPropertyName:                                     "solace-dsemp-agent",
	}

	transportSecurityStrategy := config.NewTransportSecurityStrategy().
		WithExcludedProtocols(
			config.TransportSecurityProtocolTLSv1,
			config.TransportSecurityProtocolTLSv1_1,
			config.TransportSecurityProtocolSSLv3)

	if sc.config.ValidateCert {
		if sc.config.TrustStorePath != "" {
			brokerConfig[config.TransportLayerSecurityPropertyTrustStorePath] = path

			transportSecurityStrategy = transportSecurityStrategy.
				WithCertificateValidation(true, true, path, "")
		}
	} else {
		transportSecurityStrategy = transportSecurityStrategy.
			WithoutCertificateValidation()
	}

	messagingService, err := messaging.NewMessagingServiceBuilder().FromConfigurationProvider(brokerConfig).
		WithTransportSecurityStrategy(transportSecurityStrategy).
		Build()
	if err != nil {
		return fmt.Errorf("failed to build messaging service: %v", err)
	}

	sc.messaging = messagingService

	messagingService.AddReconnectionAttemptListener(func(e solace.ServiceEvent) {
		logrus.WithField("category", category).Infof("isConnected: %v", messagingService.IsConnected())
		if err := e.GetCause(); err != nil {
			logrus.WithField("category", category).Errorf("Reconnection attempt error: %v", err)
		}
	})
	messagingService.AddReconnectionListener(func(e solace.ServiceEvent) {
		if err := e.GetCause(); err != nil {
			logrus.WithField("category", category).Infof("Reconnection successful after error: %v", err)
		}
	})

	if err := messagingService.Connect(); err != nil {
		return fmt.Errorf("failed to connect to broker: %v", err)
	}

	for !messagingService.IsConnected() {
		logrus.WithField("category", category).Info("Currently not connected, attempting to reconnect...")
		time.Sleep(500 * time.Millisecond)
	}
	logrus.WithField("category", category).Info("Connected to broker")

	topicSub := resource.TopicSubscriptionOf(sc.config.Topic)
	logrus.WithField("category", category).Infof("Subscribing to: %s", sc.config.Topic)

	directReceiver, err := messagingService.CreateDirectMessageReceiverBuilder().
		WithSubscriptions(topicSub).
		Build()

	if err != nil {
		sc.messaging.Disconnect()
		return fmt.Errorf("failed to create message receiver: %v", err)
	}

	sc.receiver = directReceiver

	if err := directReceiver.Start(); err != nil {
		sc.messaging.Disconnect()
		return fmt.Errorf("failed to start message receiver: %v", err)
	}

	logrus.WithField("category", category).Info("Message receiver started")

	if err := directReceiver.ReceiveAsync(sc.handleMessage); err != nil {
		sc.messaging.Disconnect()
		return fmt.Errorf("failed to register message handler: %v", err)
	}

	return nil
}

func (sc *SolaceClient) Disconnect() error {
	if sc.receiver != nil {
		sc.receiver.Terminate(10)
	}
	if sc.messaging != nil {
		sc.messaging.Disconnect()
	}
	return nil
}

// handleMessage processes incoming messages
func (sc *SolaceClient) handleMessage(msg message.InboundMessage) {
	var payload []byte

	// Try to get payload as string first, fallback to bytes
	if payloadStr, ok := msg.GetPayloadAsString(); ok {
		payload = []byte(payloadStr)
	} else if payloadBytes, ok := msg.GetPayloadAsBytes(); ok {
		payload = payloadBytes
	} else {
		logrus.WithField("category", category).Error("Failed to get payload as string or bytes")
		return
	}

	// Unmarshal directly into TargetState (supporting both YAML and JSON)
	var state TargetState
	if err := yaml.Unmarshal(payload, &state); err != nil {
		logrus.WithFields(logrus.Fields{
			"category": category,
			"error":    err,
			"payload":  string(payload),
		}).Error("Failed to unmarshal payload (YAML/JSON)")
		return
	}

	// Process environment variables in the state
	SubstituteEnvAndAesStrings(reflect.ValueOf(&state).Elem())

	if sc.messageCallback != nil {
		sc.messageCallback(&state)
	}
}

func (sc *SolaceClient) OnStateReceived(callback func(*TargetState)) {
	sc.messageCallback = callback
}

var aesPattern = regexp.MustCompile(`^\$aes\{(.+)\}$`)
var envPattern = regexp.MustCompile(`^\$env\{([A-Za-z_][A-Za-z0-9_]*)\}$`)

func SubstituteEnvAndAesStrings(v reflect.Value) {
	switch v.Kind() {
	case reflect.Ptr:
		if !v.IsNil() {
			SubstituteEnvAndAesStrings(v.Elem())
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			if field.CanSet() {
				SubstituteEnvAndAesStrings(field)
			}
		}
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			SubstituteEnvAndAesStrings(v.Index(i))
		}
	case reflect.String:
		str := v.String()
		if m := envPattern.FindStringSubmatch(str); m != nil {
			if envVal, ok := os.LookupEnv(m[1]); ok {
				v.SetString(envVal)
			}
		} else if m := aesPattern.FindStringSubmatch(str); m != nil {
			aesKey := viper.GetString("SOL_AES_KEY")
			if aesKey == "" {
				logrus.WithField("category", "AES").Errorf("Found encrypted value '%s' but SOL_AES_KEY environment variable is not set. Aborting.", str)
			}
			if decrypted, err := decryptAESCBCBase64(m[1], aesKey); err == nil {
				v.SetString(decrypted)
			} else {
				logrus.WithFields(logrus.Fields{
					"category": "AES",
					"value":    str,
					"error":    err,
				}).Errorf("Failed to decrypt value. Aborting.")
			}
		}
	}
}

func pkcs7Unpad(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, fmt.Errorf("data is empty")
	}
	padLen := int(data[length-1])
	if padLen > length {
		return nil, fmt.Errorf("invalid padding")
	}
	return data[:length-padLen], nil
}

func decryptAESCBCBase64(cipherBase64, keyText string) (string, error) {
	// Decode Base64
	cipherData, err := base64.StdEncoding.DecodeString(cipherBase64)
	if err != nil {
		return "", err
	}

	// Key (must be 16, 24, or 32 bytes for AES-128, AES-192, or AES-256)
	key := []byte(keyText)
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", fmt.Errorf("invalid AES key size: %d bytes. Must be 16, 24, or 32", len(key))
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(cipherData) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	// Extract IV from the beginning of the ciphertext
	iv := cipherData[:aes.BlockSize]
	ciphertext := cipherData[aes.BlockSize:]

	if len(ciphertext)%aes.BlockSize != 0 {
		return "", fmt.Errorf("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	plaintext, err = pkcs7Unpad(plaintext)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
