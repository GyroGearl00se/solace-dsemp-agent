package state

import (
	"os"
	"reflect"
	"regexp"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
	"github.com/ghodss/yaml"
	"github.com/sirupsen/logrus"

	"github.com/GyroGearl00se/solace-dsemp-agent/pkg/crypto"
)

// TargetState holds the entire declarative target specification.
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

// QueueWithSubscriptions pairs a queue with its topic subscriptions.
type QueueWithSubscriptions struct {
	swagger.MsgVpnQueue
	QueueSubscriptions []string `json:"queueSubscriptions,omitempty"`
}

// RestDeliveryPoints defines an RDP and its children.
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

// PublishException represents an ACL profile publish exception.
type PublishException struct {
	PublishExceptionTopic string `json:"publishExceptionTopic"`
	TopicSyntax           string `json:"topicSyntax"`
}

// SubscribeException represents an ACL profile subscribe exception.
type SubscribeException struct {
	SubscribeExceptionTopic string `json:"subscribeExceptionTopic"`
	TopicSyntax             string `json:"topicSyntax"`
}

// ACLProfileWithExceptions pairs an ACL Profile with publish and subscribe exceptions.
type ACLProfileWithExceptions struct {
	swagger.MsgVpnAclProfile
	PublishExceptions   []PublishException   `json:"publishExceptions,omitempty"`
	SubscribeExceptions []SubscribeException `json:"subscribeExceptions,omitempty"`
}

// WhitelistPatterns allows specifying patterns for whitelisted resources.
type WhitelistPatterns struct {
	QueuePatterns         []string `json:"queuePatterns"`
	TopicEndpointPatterns []string `json:"topicEndpointPatterns"`
}

var (
	aesPattern = regexp.MustCompile(`^\$aes\{(.+)\}$`)
	envPattern = regexp.MustCompile(`^\$env\{([A-Za-z_][A-Za-z0-9_]*)\}$`)
)

// ParseTargetState deserializes JSON/YAML into a TargetState and decrypts $env and $aes strings.
func ParseTargetState(data []byte, aesKey string) (*TargetState, error) {
	var state TargetState
	if err := yaml.Unmarshal(data, &state); err != nil {
		return nil, err
	}
	SubstituteEnvAndAesStrings(reflect.ValueOf(&state).Elem(), aesKey)
	return &state, nil
}

// SubstituteEnvAndAesStrings recursively replaces $env{VAR} and $aes{CIPHER} patterns.
func SubstituteEnvAndAesStrings(v reflect.Value, aesKey string) {
	switch v.Kind() {
	case reflect.Ptr:
		if !v.IsNil() {
			SubstituteEnvAndAesStrings(v.Elem(), aesKey)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			if field.CanSet() {
				SubstituteEnvAndAesStrings(field, aesKey)
			}
		}
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			SubstituteEnvAndAesStrings(v.Index(i), aesKey)
		}
	case reflect.String:
		str := v.String()
		if m := envPattern.FindStringSubmatch(str); m != nil {
			if envVal, ok := os.LookupEnv(m[1]); ok {
				v.SetString(envVal)
			}
		} else if m := aesPattern.FindStringSubmatch(str); m != nil {
			if aesKey == "" {
				logrus.WithField("category", "AES").Errorf("Found encrypted value '%s' but SOL_AES_KEY is not set.", str)
				return
			}
			if decrypted, err := crypto.DecryptAESCBCBase64(m[1], aesKey); err == nil {
				v.SetString(decrypted)
			} else {
				logrus.WithFields(logrus.Fields{
					"category": "AES",
					"value":    str,
					"error":    err,
				}).Errorf("Failed to decrypt value: %v", err)
			}
		}
	}
}

