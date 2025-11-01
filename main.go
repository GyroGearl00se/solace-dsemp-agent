package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/GyroGearl00se/solace-dsemp-agent/config"
	"github.com/GyroGearl00se/solace-dsemp-agent/controllers"
	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
)

func main() {
	//Viper config loading
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	_ = viper.ReadInConfig()
	viper.AutomaticEnv()

	// Configure logrus
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	viper.SetDefault("logLevel", 4)
	logLevel := viper.GetInt("LOGLEVEL")
	logrus.SetLevel(logrus.Level(logLevel))

	logrus.WithField("category", "Info").Info("Solace DSEMP Agent started...")

	brokerURL := viper.GetString("SOL_SEMP_BROKER_URL")
	sempUser := viper.GetString("SOL_SEMP_USER")
	sempPass := viper.GetString("SOL_SEMP_PASS")
	viper.SetDefault("SOL_SEMP_MSG_VPN", "default")
	msgVPN := viper.GetString("SOL_SEMP_MSG_VPN")

	// New Solace messaging configuration for state consumption
	solaceConfig := config.NewDefaultSolaceConfig()
	solaceConfig.URL = viper.GetString("SOL_STATE_BROKER_URL")
	solaceConfig.VPN = viper.GetString("SOL_STATE_MSG_VPN")
	solaceConfig.Username = viper.GetString("SOL_STATE_USERNAME")
	solaceConfig.Password = viper.GetString("SOL_STATE_PASSWORD")
	solaceConfig.Topic = viper.GetString("SOL_STATE_TOPIC")
	viper.SetDefault("SOL_RECONNECT_INTERVAL", "5")
	reconnectInterval, _ := time.ParseDuration(viper.GetString("SOL_RECONNECT_INTERVAL"))

	// Viper default for resource management
	viper.SetDefault("SOL_MANAGE_QUEUES", false)
	viper.SetDefault("SOL_MANAGE_ACL_PROFILES", false)
	viper.SetDefault("SOL_MANAGE_CLIENT_USERNAMES", false)
	viper.SetDefault("SOL_MANAGE_CLIENT_PROFILES", false)
	viper.SetDefault("SOL_MANAGE_BRIDGES", false)
	viper.SetDefault("SOL_MANAGE_DMR_BRIDGES", false)
	viper.SetDefault("SOL_MANAGE_BRIDGE_REMOTE_MSG_VPNS", false)
	viper.SetDefault("SOL_MANAGE_JNDI_CONNECTION_FACTORIES", false)
	viper.SetDefault("SOL_MANAGE_JNDI_QUEUES", false)
	viper.SetDefault("SOL_MANAGE_JNDI_TOPICS", false)
	viper.SetDefault("SOL_MANAGE_PROXIES", false)
	viper.SetDefault("SOL_MANAGE_QUEUE_TEMPLATES", false)
	viper.SetDefault("SOL_MANAGE_TOPIC_ENDPOINTS", false)
	viper.SetDefault("SOL_MANAGE_TOPIC_ENDPOINT_TEMPLATES", false)
	viper.SetDefault("SOL_MANAGE_MSG_VPNS", false)

	// Get Management options
	manageQueues := viper.GetBool("SOL_MANAGE_QUEUES")
	manageACLProfiles := viper.GetBool("SOL_MANAGE_ACL_PROFILES")
	manageClientUsernames := viper.GetBool("SOL_MANAGE_CLIENT_USERNAMES")
	manageClientProfiles := viper.GetBool("SOL_MANAGE_CLIENT_PROFILES")
	manageBridges := viper.GetBool("SOL_MANAGE_BRIDGES")
	manageDMRBridges := viper.GetBool("SOL_MANAGE_DMR_BRIDGES")
	manageBridgeRemoteMsgVpns := viper.GetBool("SOL_MANAGE_BRIDGE_REMOTE_MSG_VPNS")
	manageJndiConnectionFactories := viper.GetBool("SOL_MANAGE_JNDI_CONNECTION_FACTORIES")
	manageJndiQueues := viper.GetBool("SOL_MANAGE_JNDI_QUEUES")
	manageJndiTopics := viper.GetBool("SOL_MANAGE_JNDI_TOPICS")
	manageProxies := viper.GetBool("SOL_MANAGE_PROXIES")
	manageQueueTemplates := viper.GetBool("SOL_MANAGE_QUEUE_TEMPLATES")
	manageTopicEndpoints := viper.GetBool("SOL_MANAGE_TOPIC_ENDPOINTS")
	manageTopicEndpointTemplates := viper.GetBool("SOL_MANAGE_TOPIC_ENDPOINT_TEMPLATES")
	manageMsgVpns := viper.GetBool("SOL_MANAGE_MSG_VPNS")

	// Set default for dry run and get the value
	viper.SetDefault("SOL_DRYRUN", false)
	dryRun := viper.GetBool("SOL_DRYRUN")

	if brokerURL == "" {
		logrus.WithField("category", "Config").Fatal("SOL_SEMP_BROKER_URL must be set!")
	}
	if sempUser == "" || sempPass == "" {
		logrus.WithField("category", "Config").Fatal("SOL_SEMP_USER and SOL_SEMP_PASS must be set!")
	}

	if solaceConfig.URL == "" {
		logrus.WithField("category", "Config").Fatal("SOL_STATE_BROKER_URL must be set!")
	}
	if solaceConfig.Username == "" || solaceConfig.Password == "" {
		logrus.WithField("category", "Config").Fatal("SOL_STATE_USERNAME and SOL_STATE_PASSWORD must be set!")
	}
	if solaceConfig.Topic == "" {
		logrus.WithField("category", "Config").Fatal("SOL_STATE_TOPIC must be set!")
	}

	swaggerConfHost, _ := url.Parse(brokerURL)

	swaggerConf := swagger.NewConfiguration()
	swaggerConf.Host = swaggerConfHost.Hostname()
	swaggerConf.BasePath = brokerURL + "/SEMP/v2/config"
	swaggerConf.Scheme = swaggerConfHost.Scheme

	ctx := context.WithValue(context.Background(), swagger.ContextBasicAuth, swagger.BasicAuth{
		UserName: sempUser,
		Password: sempPass,
	})

	// Create Solace client for consuming state
	solaceClient := config.NewSolaceClient(solaceConfig)
	solaceClient.OnStateReceived(func(state *config.TargetState) {
		//logrus.WithField("category", "State").Infof("Received state: %+v", state)
		var hadError bool
		var errorList []config.Error

		if state != nil {
			if dryRun {
				logrus.WithField("category", "RunMode").Warn("DRY RUN mode is enabled. No changes will be applied.")
			}

			logrus.WithField("category", "Version").Infof("Found version tag: %s", state.Version)
			sempClient := swagger.NewAPIClient(swaggerConf)

			collect := func(categoryLabel string, errs []config.Error) {
				if len(errs) == 0 {
					return
				}
				hadError = true
				for _, e := range errs {
					logrus.WithFields(logrus.Fields{
						"category": e.Category,
						"id":       e.ResourceID,
						"action":   e.Action,
						"message":  e.Message,
					}).Errorf("%s error: %s", categoryLabel, e.Message)
				}
				errorList = append(errorList, errs...)
			}

			var queueHandler, queueSubscriptionHandler, aclHandler, aclPublishHandler, aclSubscribeHandler, clientUsernameHandler, clientProfileHandler, bridgeHandler, bridgeRemoteMsgVpnHandler, dmrBridgeHandler, jndiConnectionFactoryHandler, jndiQueueHandler, jndiTopicHandler, proxyHandler, queueTemplateHandler, topicEndpointHandler, topicEndpointTemplateHandler, msgVpnHandler *controllers.GenericCRUDHandler

			if manageQueues {
				queueHandler = &controllers.GenericCRUDHandler{
					ResourceType: "Queue",
					Controller: &controllers.QueueController{
						WhitelistPatterns: state.WhitelistPatterns.QueuePatterns,
					},
					GetState: func() []interface{} {
						// Extract the swagger.MsgVpnQueue from our custom struct
						result := make([]interface{}, len(state.Queues))
						for i, q := range state.Queues {
							result[i] = q.MsgVpnQueue
						}
						return result
					},
				}
			}

			if manageACLProfiles {
				aclHandler = &controllers.GenericCRUDHandler{
					ResourceType: "ACL Profile",
					Controller:   &controllers.ACLProfileController{},
					GetState: func() []interface{} {
						result := make([]interface{}, len(state.ACLProfiles))
						for i, acl := range state.ACLProfiles {
							result[i] = acl.MsgVpnAclProfile
						}
						return result
					},
				}
			}

			if manageClientUsernames {
				clientUsernameHandler = &controllers.GenericCRUDHandler{
					ResourceType: "Client Username",
					Controller:   &controllers.ClientUsernameController{},
					GetState: func() []interface{} {
						result := make([]interface{}, len(state.ClientUsernames))
						for i, cu := range state.ClientUsernames {
							result[i] = cu
						}
						return result
					},
				}
			}

			if manageClientProfiles {
				clientProfileHandler = &controllers.GenericCRUDHandler{
					ResourceType: "Client Profile",
					Controller:   &controllers.ClientProfileController{},
					GetState: func() []interface{} {
						result := make([]interface{}, len(state.ClientProfiles))
						for i, cp := range state.ClientProfiles {
							result[i] = cp
						}
						return result
					},
				}
			}

			if manageBridges {
				bridgeHandler = &controllers.GenericCRUDHandler{
					ResourceType: "Bridge",
					Controller:   &controllers.BridgeController{},
					GetState: func() []interface{} {
						result := make([]interface{}, len(state.Bridges))
						for i, b := range state.Bridges {
							result[i] = b
						}
						return result
					},
				}
			}

			if manageBridgeRemoteMsgVpns {
				bridgeRemoteMsgVpnHandler = &controllers.GenericCRUDHandler{
					ResourceType: "Bridge Remote Msg VPN",
					Controller:   &controllers.BridgeRemoteMsgVpnController{},
					GetState: func() []interface{} {
						result := make([]interface{}, len(state.BridgeRemoteMsgVpns))
						for i, b := range state.BridgeRemoteMsgVpns {
							result[i] = b
						}
						return result
					},
				}
			}

			if manageDMRBridges {
				dmrBridgeHandler = &controllers.GenericCRUDHandler{
					ResourceType: "DMR Bridge",
					Controller:   &controllers.DmrBridgeController{},
					GetState: func() []interface{} {
						result := make([]interface{}, len(state.DmrBridges))
						for i, b := range state.DmrBridges {
							result[i] = b
						}
						return result
					},
				}
			}

			if manageJndiConnectionFactories {
				jndiConnectionFactoryHandler = &controllers.GenericCRUDHandler{
					ResourceType: "JNDI Connection Factory",
					Controller:   &controllers.JndiConnectionFactoryController{},
					GetState: func() []interface{} {
						result := make([]interface{}, len(state.JndiConnectionFactories))
						for i, jcf := range state.JndiConnectionFactories {
							result[i] = jcf
						}
						return result
					},
				}
			}

			if manageJndiQueues {
				jndiQueueHandler = &controllers.GenericCRUDHandler{
					ResourceType: "JNDI Queue",
					Controller:   &controllers.JndiQueueController{},
					GetState: func() []interface{} {
						result := make([]interface{}, len(state.JndiQueues))
						for i, jq := range state.JndiQueues {
							result[i] = jq
						}
						return result
					},
				}
			}

			if manageJndiTopics {
				jndiTopicHandler = &controllers.GenericCRUDHandler{
					ResourceType: "JNDI Topic",
					Controller:   &controllers.JndiTopicController{},
					GetState: func() []interface{} {
						result := make([]interface{}, len(state.JndiTopics))
						for i, jt := range state.JndiTopics {
							result[i] = jt
						}
						return result
					},
				}
			}

			if manageProxies {
				proxyHandler = &controllers.GenericCRUDHandler{
					ResourceType: "Proxy",
					Controller:   &controllers.ProxyController{},
					GetState: func() []interface{} {
						result := make([]interface{}, len(state.Proxies))
						for i, p := range state.Proxies {
							result[i] = p
						}
						return result
					},
				}
			}

			if manageQueueTemplates {
				queueTemplateHandler = &controllers.GenericCRUDHandler{
					ResourceType: "Queue Template",
					Controller:   &controllers.QueueTemplateController{},
					GetState: func() []interface{} {
						result := make([]interface{}, len(state.QueueTemplates))
						for i, qt := range state.QueueTemplates {
							result[i] = qt
						}
						return result
					},
				}
			}

			if manageTopicEndpoints {
				topicEndpointHandler = &controllers.GenericCRUDHandler{
					ResourceType: "Topic Endpoint",
					Controller: &controllers.TopicEndpointController{
						WhitelistPatterns: state.WhitelistPatterns.TopicEndpointPatterns,
					},
					GetState: func() []interface{} {
						result := make([]interface{}, len(state.TopicEndpoints))
						for i, te := range state.TopicEndpoints {
							result[i] = te
						}
						return result
					},
				}
			}

			if manageTopicEndpointTemplates {
				topicEndpointTemplateHandler = &controllers.GenericCRUDHandler{
					ResourceType: "Topic Endpoint Template",
					Controller:   &controllers.TopicEndpointTemplateController{},
					GetState: func() []interface{} {
						result := make([]interface{}, len(state.TopicEndpointsTemplates))
						for i, tet := range state.TopicEndpointsTemplates {
							result[i] = tet
						}
						return result
					},
				}
			}

			if manageMsgVpns {
				msgVpnHandler = &controllers.GenericCRUDHandler{
					ResourceType: "MsgVpn",
					Controller:   &controllers.MsgVpnController{},
					GetState: func() []interface{} {
						result := make([]interface{}, len(state.MsgVpns))
						for i, tet := range state.MsgVpns {
							result[i] = tet
						}
						return result
					},
				}
			}

			// Delete resources that are present on the broker but not in the desired state
			if manageJndiConnectionFactories {
				collect("jndi-connection-factories", jndiConnectionFactoryHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "delete"))
			}
			if manageJndiQueues {
				collect("jndi-queues", jndiQueueHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "delete"))
			}
			if manageJndiTopics {
				collect("jndi-topics", jndiTopicHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "delete"))
			}
			if manageBridgeRemoteMsgVpns {
				collect("bridge-remote-msg-vpns", bridgeRemoteMsgVpnHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "delete"))
			}
			if manageBridges {
				collect("bridges", bridgeHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "delete"))
			}
			if manageDMRBridges {
				collect("dmr-bridges", dmrBridgeHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "delete"))
			}
			if manageQueues {
				collect("queues", queueHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "delete"))
			}
			if manageTopicEndpoints {
				collect("topic-endpoints", topicEndpointHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "delete"))
			}
			if manageQueueTemplates {
				collect("queue-templates", queueTemplateHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "delete"))
			}
			if manageTopicEndpointTemplates {
				collect("topic-endpoint-templates", topicEndpointTemplateHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "delete"))
			}
			if manageClientUsernames {
				collect("client-usernames", clientUsernameHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "delete"))
			}
			if manageACLProfiles {
				collect("acl-profiles", aclHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "delete"))
			}
			if manageClientProfiles {
				collect("client-profiles", clientProfileHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "delete"))
			}
			if manageProxies {
				collect("proxies", proxyHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "delete"))
			}
			if manageMsgVpns {
				collect("msg-vpns", msgVpnHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "delete"))
			}

			// Upsert resources that are in the desired state
			if manageMsgVpns {
				collect("msg-vpns", msgVpnHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "upsert"))
			}
			if manageACLProfiles {
				collect("acl-profiles", aclHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "upsert"))
				// Process ACL Profile Publish and Subscribe Exceptions for each ACL Profile
				for _, aclProfile := range state.ACLProfiles {
					currentProfile := aclProfile

					// Handle Publish Exceptions
					if len(currentProfile.PublishExceptions) > 0 {
						aclPublishHandler = &controllers.GenericCRUDHandler{
							ResourceType: "ACL Profile Publish Exception",
							Controller:   &controllers.ACLProfileExceptionController{IsPublishException: true, AclProfileName: currentProfile.AclProfileName},
							GetState: func() []interface{} {
								exceptions := make([]interface{}, len(currentProfile.PublishExceptions))
								for i, ex := range currentProfile.PublishExceptions {
									exceptions[i] = swagger.MsgVpnAclProfilePublishException{
										AclProfileName:        currentProfile.AclProfileName,
										PublishExceptionTopic: ex.PublishExceptionTopic,
										TopicSyntax:           ex.TopicSyntax,
									}
								}
								return exceptions
							},
						}
						collect(fmt.Sprintf("acl-profile publish exceptions for %s", currentProfile.AclProfileName), aclPublishHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "crud"))
					}

					// Handle Subscribe Exceptions
					if len(currentProfile.SubscribeExceptions) > 0 {
						aclSubscribeHandler = &controllers.GenericCRUDHandler{
							ResourceType: "ACL Profile Subscribe Exception",
							Controller:   &controllers.ACLProfileExceptionController{IsPublishException: false, AclProfileName: currentProfile.AclProfileName},
							GetState: func() []interface{} {
								exceptions := make([]interface{}, len(currentProfile.SubscribeExceptions))
								for i, ex := range currentProfile.SubscribeExceptions {
									exceptions[i] = swagger.MsgVpnAclProfileSubscribeException{
										AclProfileName:          currentProfile.AclProfileName,
										SubscribeExceptionTopic: ex.SubscribeExceptionTopic,
										TopicSyntax:             ex.TopicSyntax,
									}
								}
								return exceptions
							},
						}
						collect(fmt.Sprintf("acl-profile subscribe exceptions for %s", currentProfile.AclProfileName), aclSubscribeHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "crud"))
					}
				}
			}
			if manageClientProfiles {
				collect("client-profiles", clientProfileHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "upsert"))
			}
			if manageClientUsernames {
				collect("client-usernames", clientUsernameHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "upsert"))
			}
			if manageQueueTemplates {
				collect("queue-templates", queueTemplateHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "upsert"))
			}
			if manageTopicEndpointTemplates {
				collect("topic-endpoint-templates", topicEndpointTemplateHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "upsert"))
			}
			if manageQueues {
				collect("queues", queueHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "upsert"))
				// Process subscriptions for each queue
				for _, queue := range state.Queues {
					currentQueue := queue

					if len(currentQueue.QueueSubscriptions) > 0 {
						queueSubscriptionHandler = &controllers.GenericCRUDHandler{
							ResourceType: "QueueSubscription",
							Controller:   &controllers.QueueSubscriptionController{QueueName: currentQueue.QueueName},
							GetState: func() []interface{} {
								subscriptions := make([]interface{}, len(currentQueue.QueueSubscriptions))
								for i, subTopic := range currentQueue.QueueSubscriptions {
									subscriptions[i] = swagger.MsgVpnQueueSubscription{
										SubscriptionTopic: subTopic,
										QueueName:         currentQueue.QueueName,
									}
								}
								return subscriptions
							},
						}
						collect(fmt.Sprintf("queue subscriptions for %s", currentQueue.QueueName), queueSubscriptionHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "crud"))
					}
				}
			}
			if manageTopicEndpoints {
				collect("topic-endpoints", topicEndpointHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "upsert"))
			}
			if manageBridges {
				collect("bridges", bridgeHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "upsert"))
			}
			if manageBridgeRemoteMsgVpns {
				collect("bridge-remote-msg-vpns", bridgeRemoteMsgVpnHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "upsert"))
			}
			if manageDMRBridges {
				collect("dmr-bridges", dmrBridgeHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "upsert"))
			}
			if manageJndiConnectionFactories {
				collect("jndi-connection-factories", jndiConnectionFactoryHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "upsert"))
			}
			if manageJndiQueues {
				collect("jndi-queues", jndiQueueHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "upsert"))
			}
			if manageJndiTopics {
				collect("jndi-topics", jndiTopicHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "upsert"))
			}
			if manageProxies {
				collect("proxies", proxyHandler.ProcessState(ctx, sempClient, msgVPN, dryRun, "upsert"))
			}

			// Send status report

			if viper.GetBool("SOL_STATUS_WEBHOOK_ENABLED") || viper.GetBool("SOL_STATUS_MESSAGE_ENABLED") {
				logrus.WithField("category", "Statusreport").Infof("collected errors: %+v", errorList)
				sendStatusReport(!hadError, errorList, state)
			}
		}

	})

	// Connect to Solace with retry
	for {
		logrus.WithField("category", "Solace state consumer").Info("Attempting to connect to Solace broker...")

		if err := solaceClient.Connect(); err != nil {
			logrus.WithField("category", "Solace state consumer").Errorf("Failed to connect to Solace broker: %v", err)
			logrus.WithField("category", "Solace state consumer").Infof("Retrying in %v...", reconnectInterval)
			time.Sleep(reconnectInterval)
			continue
		}

		logrus.WithField("category", "Solace state consumer").Info("Successfully connected to Solace broker")
		break
	}
	defer solaceClient.Disconnect()

	select {}
}

func sendStatusReport(success bool, errors []config.Error, state *config.TargetState) {
	const category = "Statusreport"
	extraFields := viper.GetStringMapString("SOL_STATUS_EXTRA_FIELDS")
	logrus.WithField("category", category).Infof("sending status, success=%v, errors=%d", success, len(errors))
	msg := map[string]interface{}{
		"timestamp":     time.Now().UTC().Format(time.RFC3339),
		"success":       success,
		"errors":        nil,
		"configVersion": state.Version,
	}
	if len(errors) > 0 {
		msg["errors"] = errors
	}
	for k, v := range extraFields {
		msg[k] = v
	}
	body, _ := json.Marshal(msg)

	if viper.GetBool("SOL_STATUS_WEBHOOK_ENABLED") {
		webhookURL := viper.GetString("SOL_STATUS_WEBHOOK_URL")
		if webhookURL == "" {
			return
		}

		user := viper.GetString("SOL_STATUS_WEBHOOK_USER")
		pass := viper.GetString("SOL_STATUS_WEBHOOK_PASS")

		logrus.WithField("category", category).Infof("Starting WEBHOOK")

		client := &http.Client{Timeout: 10 * time.Second}

		req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(body))
		if err != nil {
			logrus.WithField("category", category).Errorf("Failed to create webhook request: %v", err)
			return
		}

		req.Header.Set("Content-Type", "application/json")
		if user != "" && pass != "" {

			req.SetBasicAuth(user, pass)
		}

		resp, err := client.Do(req)
		if err != nil {
			logrus.WithField("category", category).Errorf("Failed to send status webhook: %v", err)
			return
		}

		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			logrus.WithField("category", category).Warnf("Status webhook returned HTTP %d", resp.StatusCode)
		}

		logrus.WithField("category", category).Infof("WEBHOOK successfully sent to %s", webhookURL)
	}
	if viper.GetBool("SOL_STATUS_MESSAGE_ENABLED") {
		brokerURL := viper.GetString("SOL_STATUS_MESSAGE_BROKER_URL")
		topic := viper.GetString("SOL_STATUS_MESSAGE_TOPIC")
		brokerUser := viper.GetString("SOL_STATUS_MESSAGE_BROKER_USER")
		brokerPass := viper.GetString("SOL_STATUS_MESSAGE_BROKER_PASS")
		msgVpn := viper.GetString("SOL_STATUS_MESSAGE_MSG_VPN")

		if brokerURL == "" || topic == "" || brokerUser == "" || brokerPass == "" || msgVpn == "" {
			logrus.WithField("category", category).Error("Status message configuration is incomplete")
			return
		}

		if err := config.SendStatusMessage(success, body, brokerURL, topic, brokerUser, brokerPass, msgVpn); err != nil {
			logrus.WithField("category", category).Errorf("Failed to send status message: %v", err)
			return
		}

		logrus.WithField("category", category).Infof("Status message sent successfully to %s on topic %s", brokerURL, topic)
	}

}
