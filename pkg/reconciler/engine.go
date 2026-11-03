package reconciler

import (
	"context"
	"fmt"
	"sync"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
	"github.com/sirupsen/logrus"

	"github.com/GyroGearl00se/solace-dsemp-agent/pkg/controllers"
	"github.com/GyroGearl00se/solace-dsemp-agent/pkg/semputil"
	"github.com/GyroGearl00se/solace-dsemp-agent/pkg/state"
)

// ManagementFlags controls which resource families the agent manages.
type ManagementFlags struct {
	MsgVpns                 bool
	Queues                  bool
	ACLProfiles             bool
	ClientUsernames         bool
	ClientProfiles          bool
	Bridges                 bool
	DMRBridges              bool
	BridgeRemoteMsgVpns     bool
	JndiConnectionFactories bool
	JndiQueues              bool
	JndiTopics              bool
	Proxies                 bool
	RestDeliveryPoints      bool
	QueueTemplates          bool
	TopicEndpoints          bool
	TopicEndpointTemplates  bool
}

// Engine coordinates the declarative state reconciliation on a Solace broker.
type Engine struct {
	client     *swagger.APIClient
	msgVpn     string
	dryRun     bool
	flags      ManagementFlags
	mutex      sync.Mutex
}

// NewEngine constructs a reconciliation engine.
func NewEngine(client *swagger.APIClient, msgVpn string, dryRun bool, flags ManagementFlags) *Engine {
	return &Engine{
		client: client,
		msgVpn: msgVpn,
		dryRun: dryRun,
		flags:  flags,
	}
}

// Reconcile executes a synchronized reconciliation loop against the broker.
func (e *Engine) Reconcile(ctx context.Context, target *state.TargetState) (bool, []state.Error) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	var errorList []state.Error
	var hadError bool

	if target == nil {
		return true, nil
	}

	if e.dryRun {
		logrus.WithField("category", "RunMode").Warn("DRY RUN mode is enabled. No changes will be applied.")
	}
	logrus.WithField("category", "Version").Infof("Reconciling desired state version: %s", target.Version)

	collect := func(categoryLabel string, errs []state.Error) {
		if len(errs) == 0 {
			return
		}
		hadError = true
		for _, err := range errs {
			logrus.WithFields(logrus.Fields{
				"category": err.Category,
				"id":       err.ResourceID,
				"action":   err.Action,
				"message":  err.Message,
			}).Errorf("%s error: %s", categoryLabel, err.Message)
		}
		errorList = append(errorList, errs...)
	}

	// Instantiate primary handlers
	var queueHandler, aclHandler, clientUsernameHandler, clientProfileHandler, bridgeHandler,
		bridgeRemoteMsgVpnHandler, dmrBridgeHandler, jndiConnectionFactoryHandler, jndiQueueHandler,
		jndiTopicHandler, proxyHandler, rdpHandler, queueTemplateHandler, topicEndpointHandler,
		topicEndpointTemplateHandler, msgVpnHandler *controllers.GenericCRUDHandler

	if e.flags.Queues {
		queueHandler = &controllers.GenericCRUDHandler{
			ResourceType: "Queue",
			Controller:   &controllers.QueueController{WhitelistPatterns: target.WhitelistPatterns.QueuePatterns},
			GetState: func() []interface{} {
				res := make([]interface{}, len(target.Queues))
				for i, q := range target.Queues {
					res[i] = q.MsgVpnQueue
				}
				return res
			},
		}
	}

	if e.flags.ACLProfiles {
		aclHandler = &controllers.GenericCRUDHandler{
			ResourceType: "ACL Profile",
			Controller:   &controllers.ACLProfileController{},
			GetState: func() []interface{} {
				res := make([]interface{}, len(target.ACLProfiles))
				for i, acl := range target.ACLProfiles {
					res[i] = acl.MsgVpnAclProfile
				}
				return res
			},
		}
	}

	if e.flags.ClientUsernames {
		clientUsernameHandler = &controllers.GenericCRUDHandler{
			ResourceType: "Client Username",
			Controller:   &controllers.ClientUsernameController{},
			GetState: func() []interface{} {
				res := make([]interface{}, len(target.ClientUsernames))
				for i, cu := range target.ClientUsernames {
					res[i] = cu
				}
				return res
			},
		}
	}

	if e.flags.ClientProfiles {
		clientProfileHandler = &controllers.GenericCRUDHandler{
			ResourceType: "Client Profile",
			Controller:   &controllers.ClientProfileController{},
			GetState: func() []interface{} {
				res := make([]interface{}, len(target.ClientProfiles))
				for i, cp := range target.ClientProfiles {
					res[i] = cp
				}
				return res
			},
		}
	}

	if e.flags.Bridges {
		bridgeHandler = &controllers.GenericCRUDHandler{
			ResourceType: "Bridge",
			Controller:   &controllers.BridgeController{},
			GetState: func() []interface{} {
				res := make([]interface{}, len(target.Bridges))
				for i, b := range target.Bridges {
					res[i] = b
				}
				return res
			},
		}
	}

	if e.flags.BridgeRemoteMsgVpns {
		bridgeRemoteMsgVpnHandler = &controllers.GenericCRUDHandler{
			ResourceType: "Bridge Remote Msg VPN",
			Controller:   &controllers.BridgeRemoteMsgVpnController{},
			GetState: func() []interface{} {
				res := make([]interface{}, len(target.BridgeRemoteMsgVpns))
				for i, b := range target.BridgeRemoteMsgVpns {
					res[i] = b
				}
				return res
			},
		}
	}

	if e.flags.DMRBridges {
		dmrBridgeHandler = &controllers.GenericCRUDHandler{
			ResourceType: "DMR Bridge",
			Controller:   &controllers.DmrBridgeController{},
			GetState: func() []interface{} {
				res := make([]interface{}, len(target.DmrBridges))
				for i, b := range target.DmrBridges {
					res[i] = b
				}
				return res
			},
		}
	}

	if e.flags.JndiConnectionFactories {
		jndiConnectionFactoryHandler = &controllers.GenericCRUDHandler{
			ResourceType: "JNDI Connection Factory",
			Controller:   &controllers.JndiConnectionFactoryController{},
			GetState: func() []interface{} {
				res := make([]interface{}, len(target.JndiConnectionFactories))
				for i, jcf := range target.JndiConnectionFactories {
					res[i] = jcf
				}
				return res
			},
		}
	}

	if e.flags.JndiQueues {
		jndiQueueHandler = &controllers.GenericCRUDHandler{
			ResourceType: "JNDI Queue",
			Controller:   &controllers.JndiQueueController{},
			GetState: func() []interface{} {
				res := make([]interface{}, len(target.JndiQueues))
				for i, jq := range target.JndiQueues {
					res[i] = jq
				}
				return res
			},
		}
	}

	if e.flags.JndiTopics {
		jndiTopicHandler = &controllers.GenericCRUDHandler{
			ResourceType: "JNDI Topic",
			Controller:   &controllers.JndiTopicController{},
			GetState: func() []interface{} {
				res := make([]interface{}, len(target.JndiTopics))
				for i, jt := range target.JndiTopics {
					res[i] = jt
				}
				return res
			},
		}
	}

	if e.flags.Proxies {
		proxyHandler = &controllers.GenericCRUDHandler{
			ResourceType: "Proxy",
			Controller:   &controllers.ProxyController{},
			GetState: func() []interface{} {
				res := make([]interface{}, len(target.Proxies))
				for i, p := range target.Proxies {
					res[i] = p
				}
				return res
			},
		}
	}

	if e.flags.QueueTemplates {
		queueTemplateHandler = &controllers.GenericCRUDHandler{
			ResourceType: "Queue Template",
			Controller:   &controllers.QueueTemplateController{},
			GetState: func() []interface{} {
				res := make([]interface{}, len(target.QueueTemplates))
				for i, qt := range target.QueueTemplates {
					res[i] = qt
				}
				return res
			},
		}
	}

	if e.flags.TopicEndpoints {
		topicEndpointHandler = &controllers.GenericCRUDHandler{
			ResourceType: "Topic Endpoint",
			Controller:   &controllers.TopicEndpointController{WhitelistPatterns: target.WhitelistPatterns.TopicEndpointPatterns},
			GetState: func() []interface{} {
				res := make([]interface{}, len(target.TopicEndpoints))
				for i, te := range target.TopicEndpoints {
					res[i] = te
				}
				return res
			},
		}
	}

	if e.flags.TopicEndpointTemplates {
		topicEndpointTemplateHandler = &controllers.GenericCRUDHandler{
			ResourceType: "Topic Endpoint Template",
			Controller:   &controllers.TopicEndpointTemplateController{},
			GetState: func() []interface{} {
				res := make([]interface{}, len(target.TopicEndpointsTemplates))
				for i, tet := range target.TopicEndpointsTemplates {
					res[i] = tet
				}
				return res
			},
		}
	}

	if e.flags.MsgVpns {
		msgVpnHandler = &controllers.GenericCRUDHandler{
			ResourceType: "MsgVpn",
			Controller:   &controllers.MsgVpnController{},
			GetState: func() []interface{} {
				res := make([]interface{}, len(target.MsgVpns))
				for i, vpn := range target.MsgVpns {
					res[i] = vpn
				}
				return res
			},
		}
	}

	if e.flags.RestDeliveryPoints {
		rdpHandler = &controllers.GenericCRUDHandler{
			ResourceType: "Rest Delivery Point",
			Controller:   &controllers.RDPController{},
			GetState: func() []interface{} {
				res := make([]interface{}, len(target.RestDeliveryPoints))
				for i, rdp := range target.RestDeliveryPoints {
					res[i] = rdp.MsgVpnRestDeliveryPoint
				}
				return res
			},
		}
	}

	// PHASE 1: DELETE (Executed in reverse dependency order)
	if e.flags.RestDeliveryPoints {
		collect("rest-delivery-points", rdpHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "delete"))
	}
	if e.flags.JndiConnectionFactories {
		collect("jndi-connection-factories", jndiConnectionFactoryHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "delete"))
	}
	if e.flags.JndiQueues {
		collect("jndi-queues", jndiQueueHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "delete"))
	}
	if e.flags.JndiTopics {
		collect("jndi-topics", jndiTopicHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "delete"))
	}
	if e.flags.BridgeRemoteMsgVpns {
		collect("bridge-remote-msg-vpns", bridgeRemoteMsgVpnHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "delete"))
	}
	if e.flags.Bridges {
		collect("bridges", bridgeHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "delete"))
	}
	if e.flags.DMRBridges {
		collect("dmr-bridges", dmrBridgeHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "delete"))
	}
	if e.flags.Queues {
		collect("queues", queueHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "delete"))
	}
	if e.flags.TopicEndpoints {
		collect("topic-endpoints", topicEndpointHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "delete"))
	}
	if e.flags.QueueTemplates {
		collect("queue-templates", queueTemplateHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "delete"))
	}
	if e.flags.TopicEndpointTemplates {
		collect("topic-endpoint-templates", topicEndpointTemplateHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "delete"))
	}
	if e.flags.ClientUsernames {
		collect("client-usernames", clientUsernameHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "delete"))
	}
	if e.flags.ACLProfiles {
		collect("acl-profiles", aclHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "delete"))
	}
	if e.flags.ClientProfiles {
		collect("client-profiles", clientProfileHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "delete"))
	}
	if e.flags.Proxies {
		collect("proxies", proxyHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "delete"))
	}
	if e.flags.MsgVpns {
		collect("msg-vpns", msgVpnHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "delete"))
	}

	// PHASE 2: UPSERT (Executed in forward dependency order)
	if e.flags.MsgVpns {
		collect("msg-vpns", msgVpnHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "upsert"))
	}

	if e.flags.ACLProfiles {
		collect("acl-profiles", aclHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "upsert"))

		for _, aclProfile := range target.ACLProfiles {
			currentProfile := aclProfile

			aclPublishHandler := &controllers.GenericCRUDHandler{
				ResourceType: "ACL Profile Publish Exception",
				Controller:   &controllers.ACLProfileExceptionController{IsPublishException: true, AclProfileName: currentProfile.AclProfileName},
				GetState: func() []interface{} {
					res := make([]interface{}, len(currentProfile.PublishExceptions))
					for i, ex := range currentProfile.PublishExceptions {
						res[i] = swagger.MsgVpnAclProfilePublishException{
							AclProfileName:        currentProfile.AclProfileName,
							PublishExceptionTopic: ex.PublishExceptionTopic,
							TopicSyntax:           ex.TopicSyntax,
						}
					}
					return res
				},
			}
			collect(fmt.Sprintf("acl-profile publish exceptions for %s", currentProfile.AclProfileName), aclPublishHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "crud"))

			aclSubscribeHandler := &controllers.GenericCRUDHandler{
				ResourceType: "ACL Profile Subscribe Exception",
				Controller:   &controllers.ACLProfileExceptionController{IsPublishException: false, AclProfileName: currentProfile.AclProfileName},
				GetState: func() []interface{} {
					res := make([]interface{}, len(currentProfile.SubscribeExceptions))
					for i, ex := range currentProfile.SubscribeExceptions {
						res[i] = swagger.MsgVpnAclProfileSubscribeException{
							AclProfileName:          currentProfile.AclProfileName,
							SubscribeExceptionTopic: ex.SubscribeExceptionTopic,
							TopicSyntax:             ex.TopicSyntax,
						}
					}
					return res
				},
			}
			collect(fmt.Sprintf("acl-profile subscribe exceptions for %s", currentProfile.AclProfileName), aclSubscribeHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "crud"))
		}
	}

	if e.flags.ClientProfiles {
		collect("client-profiles", clientProfileHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "upsert"))
	}
	if e.flags.ClientUsernames {
		collect("client-usernames", clientUsernameHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "upsert"))
	}
	if e.flags.QueueTemplates {
		collect("queue-templates", queueTemplateHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "upsert"))
	}
	if e.flags.TopicEndpointTemplates {
		collect("topic-endpoint-templates", topicEndpointTemplateHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "upsert"))
	}

	if e.flags.Queues {
		collect("queues", queueHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "upsert"))

		for _, queue := range target.Queues {
			currentQueue := queue

			queueSubHandler := &controllers.GenericCRUDHandler{
				ResourceType: "QueueSubscription",
				Controller:   &controllers.QueueSubscriptionController{QueueName: currentQueue.QueueName},
				GetState: func() []interface{} {
					res := make([]interface{}, len(currentQueue.QueueSubscriptions))
					for i, subTopic := range currentQueue.QueueSubscriptions {
						res[i] = swagger.MsgVpnQueueSubscription{
							SubscriptionTopic: subTopic,
							QueueName:         currentQueue.QueueName,
						}
					}
					return res
				},
			}
			collect(fmt.Sprintf("queue subscriptions for %s", currentQueue.QueueName), queueSubHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "crud"))
		}
	}

	if e.flags.TopicEndpoints {
		collect("topic-endpoints", topicEndpointHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "upsert"))
	}
	if e.flags.Bridges {
		collect("bridges", bridgeHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "upsert"))
	}
	if e.flags.BridgeRemoteMsgVpns {
		collect("bridge-remote-msg-vpns", bridgeRemoteMsgVpnHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "upsert"))
	}
	if e.flags.DMRBridges {
		collect("dmr-bridges", dmrBridgeHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "upsert"))
	}
	if e.flags.JndiConnectionFactories {
		collect("jndi-connection-factories", jndiConnectionFactoryHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "upsert"))
	}
	if e.flags.JndiQueues {
		collect("jndi-queues", jndiQueueHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "upsert"))
	}
	if e.flags.JndiTopics {
		collect("jndi-topics", jndiTopicHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "upsert"))
	}
	if e.flags.Proxies {
		collect("proxies", proxyHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "upsert"))
	}

	if e.flags.RestDeliveryPoints {
		collect("rest-delivery-points", rdpHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "upsert"))

		for _, rdp := range target.RestDeliveryPoints {
			currentRdp := rdp
			rdpName := currentRdp.MsgVpnRestDeliveryPoint.RestDeliveryPointName

			// Queue Binding
			rdpQueueBindingHandler := &controllers.GenericCRUDHandler{
				ResourceType: "RDP Queue Binding",
				Controller:   &controllers.RDPQueueBindingController{RDPName: rdpName},
				GetState: func() []interface{} {
					if currentRdp.MsgVpnRestDeliveryPointQueueBinding.QueueBindingName == "" {
						return []interface{}{}
					}
					binding := currentRdp.MsgVpnRestDeliveryPointQueueBinding
					binding.RestDeliveryPointName = rdpName
					return []interface{}{binding}
				},
			}
			collect(fmt.Sprintf("rdp queue bindings for %s", rdpName), rdpQueueBindingHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "crud"))

			if currentRdp.MsgVpnRestDeliveryPointQueueBinding.QueueBindingName != "" {
				queueBindingName := currentRdp.MsgVpnRestDeliveryPointQueueBinding.QueueBindingName

				headerHandler := &controllers.GenericCRUDHandler{
					ResourceType: "RDP Queue Binding Request Header",
					Controller:   &controllers.RDPQueueBindingRequestHeaderController{RDPName: rdpName, QueueBindingName: queueBindingName},
					GetState: func() []interface{} {
						res := make([]interface{}, len(currentRdp.RequestHeaders))
						for i, h := range currentRdp.RequestHeaders {
							h.RestDeliveryPointName = rdpName
							h.QueueBindingName = queueBindingName
							res[i] = h
						}
						return res
					},
				}
				collect(fmt.Sprintf("rdp queue binding request headers for %s/%s", rdpName, queueBindingName), headerHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "crud"))

				protectedHeaderHandler := &controllers.GenericCRUDHandler{
					ResourceType: "RDP Queue Binding Protected Request Header",
					Controller:   &controllers.RDPQueueBindingProtectedRequestHeaderController{RDPName: rdpName, QueueBindingName: queueBindingName},
					GetState: func() []interface{} {
						res := make([]interface{}, len(currentRdp.ProtectedRequestHeaders))
						for i, h := range currentRdp.ProtectedRequestHeaders {
							h.RestDeliveryPointName = rdpName
							h.QueueBindingName = queueBindingName
							res[i] = h
						}
						return res
					},
				}
				collect(fmt.Sprintf("rdp queue binding protected request headers for %s/%s", rdpName, queueBindingName), protectedHeaderHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "crud"))
			}

			// Rest Consumer
			rdpRestConsumerHandler := &controllers.GenericCRUDHandler{
				ResourceType: "RDP Rest Consumer",
				Controller:   &controllers.RDPRestConsumerController{RDPName: rdpName},
				GetState: func() []interface{} {
					if currentRdp.MsgVpnRestDeliveryPointRestConsumer.RestConsumerName == "" {
						return []interface{}{}
					}
					consumer := currentRdp.MsgVpnRestDeliveryPointRestConsumer
					consumer.RestDeliveryPointName = rdpName
					return []interface{}{consumer}
				},
			}
			collect(fmt.Sprintf("rdp rest consumers for %s", rdpName), rdpRestConsumerHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "crud"))

			if currentRdp.MsgVpnRestDeliveryPointRestConsumer.RestConsumerName != "" {
				restConsumerName := currentRdp.MsgVpnRestDeliveryPointRestConsumer.RestConsumerName

				tlsTrustedCNHandler := &controllers.GenericCRUDHandler{
					ResourceType: "RDP Rest Consumer TLS Trusted Common Name",
					Controller:   &controllers.RDPRestConsumerTlsTrustedCommonNameController{RDPName: rdpName, RestConsumerName: restConsumerName},
					GetState: func() []interface{} {
						res := make([]interface{}, len(currentRdp.TlsTrustedCommonNames))
						for i, cn := range currentRdp.TlsTrustedCommonNames {
							cn.RestDeliveryPointName = rdpName
							cn.RestConsumerName = restConsumerName
							res[i] = cn
						}
						return res
					},
				}

				cnErrors := tlsTrustedCNHandler.ProcessState(ctx, e.client, e.msgVpn, true, "crud")
				hasChanges := len(cnErrors) > 0 || len(currentRdp.TlsTrustedCommonNames) > 0

				if hasChanges && !e.dryRun {
					disableBody := swagger.MsgVpnRestDeliveryPointRestConsumer{Enabled: semputil.BoolPtrPtr(false)}
					_, _, err := e.client.RestDeliveryPointApi.UpdateMsgVpnRestDeliveryPointRestConsumer(ctx, disableBody, e.msgVpn, rdpName, restConsumerName, nil)
					if err != nil {
						collect(fmt.Sprintf("rdp rest consumer tls trusted common names for %s/%s", rdpName, restConsumerName), []state.Error{{
							Category:   "RDP Rest Consumer TLS Trusted Common Name",
							ResourceID: restConsumerName,
							Action:     "DISABLE_CONSUMER",
							Message:    err.Error(),
						}})
					} else {
						collect(fmt.Sprintf("rdp rest consumer tls trusted common names for %s/%s", rdpName, restConsumerName), tlsTrustedCNHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "crud"))

						if currentRdp.MsgVpnRestDeliveryPointRestConsumer.Enabled != nil && *currentRdp.MsgVpnRestDeliveryPointRestConsumer.Enabled != nil && **currentRdp.MsgVpnRestDeliveryPointRestConsumer.Enabled {
							enableBody := swagger.MsgVpnRestDeliveryPointRestConsumer{Enabled: semputil.BoolPtrPtr(true)}
							_, _, err := e.client.RestDeliveryPointApi.UpdateMsgVpnRestDeliveryPointRestConsumer(ctx, enableBody, e.msgVpn, rdpName, restConsumerName, nil)
							if err != nil {
								collect(fmt.Sprintf("rdp rest consumer tls trusted common names for %s/%s", rdpName, restConsumerName), []state.Error{{
									Category:   "RDP Rest Consumer TLS Trusted Common Name",
									ResourceID: restConsumerName,
									Action:     "ENABLE_CONSUMER",
									Message:    err.Error(),
								}})
							}
						}
					}
				} else {
					collect(fmt.Sprintf("rdp rest consumer tls trusted common names for %s/%s", rdpName, restConsumerName), tlsTrustedCNHandler.ProcessState(ctx, e.client, e.msgVpn, e.dryRun, "crud"))
				}
			}
		}
	}

	return !hadError, errorList
}

