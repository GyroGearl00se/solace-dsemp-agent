package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/GyroGearl00se/solace-dsemp-agent/config"
	"github.com/GyroGearl00se/solace-dsemp-agent/controllers"
	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
	"github.com/GyroGearl00se/solace-dsemp-agent/semplegacy"
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
	solaceConfig.TrustStorePath = viper.GetString("SOL_TRUST_STORE_PATH")
	viper.SetDefault("SOL_VALIDATE_CERT", true)
	solaceConfig.ValidateCert = viper.GetBool("SOL_VALIDATE_CERT")
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

	// Set default for dry run and get the value
	viper.SetDefault("SOL_DRYRUN", false)
	dryRun := viper.GetBool("SOL_DRYRUN")

	if brokerURL == "" {
		logrus.WithField("category", "Config").Fatal("SOL_SEMP_BROKER_URL must be set!")
	}
	if sempUser == "" || sempPass == "" {
		logrus.WithField("category", "Config").Fatal("SOL_SEMP_USER and SOL_SEMP_PASS must be set!")
	}

	viper.SetDefault("SOL_STATE_CONSUMER_ENABLED", true)
	stateConsumerEnabled := viper.GetBool("SOL_STATE_CONSUMER_ENABLED")

	viper.SetDefault("SOL_INITIAL_STATE_ENABLED", true)
	initialStateEnabled := viper.GetBool("SOL_INITIAL_STATE_ENABLED")
	initialStateFile := viper.GetString("SOL_INITIAL_STATE_FILE")

	// We must have at least one mode configured
	hasMessagingURL := solaceConfig.URL != ""
	hasMessagingCreds := solaceConfig.Username != "" || solaceConfig.Password != ""
	hasMessagingTopic := solaceConfig.Topic != ""

	hasAnyMessaging := hasMessagingURL || hasMessagingCreds || hasMessagingTopic
	messagingConfigured := stateConsumerEnabled && hasMessagingURL && solaceConfig.Username != "" && solaceConfig.Password != "" && hasMessagingTopic
	fileConfigured := initialStateEnabled && initialStateFile != ""

	if stateConsumerEnabled && hasAnyMessaging && !messagingConfigured {
		logrus.WithField("category", "Config").Fatal("Solace messaging state consumer is enabled but partially configured! You must set ALL of: SOL_STATE_BROKER_URL, SOL_STATE_USERNAME, SOL_STATE_PASSWORD, and SOL_STATE_TOPIC.")
	}

	if initialStateEnabled && initialStateFile == "" && !messagingConfigured {
		logrus.WithField("category", "Config").Fatal("Initial state file loader is enabled but SOL_INITIAL_STATE_FILE path is not set!")
	}

	if !messagingConfigured && !fileConfigured {
		logrus.WithField("category", "Config").Fatal("No active configuration source was provided! Either enable and configure the messaging state consumer (SOL_STATE_CONSUMER_ENABLED=true, SOL_STATE_BROKER_URL, SOL_STATE_USERNAME, SOL_STATE_PASSWORD, SOL_STATE_TOPIC) or enable and configure the file-based state loader (SOL_INITIAL_STATE_ENABLED=true, SOL_INITIAL_STATE_FILE).")
	}

	swaggerConfHost, _ := url.Parse(brokerURL)

	swaggerConf := swagger.NewConfiguration()
	swaggerConf.Host = swaggerConfHost.Hostname()
	swaggerConf.BasePath = brokerURL + "/SEMP/v2/config"
	swaggerConf.Scheme = swaggerConfHost.Scheme
	certPool := x509.NewCertPool()
	if solaceConfig.TrustStorePath != "" {
		if err := LoadCertificates(solaceConfig.TrustStorePath, certPool); err != nil {
			logrus.WithField("category", "Config").Fatalf("Failed to load certificates: %v", err)
		}
	}
	swaggerConf.HTTPClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: !solaceConfig.ValidateCert,
				RootCAs:            certPool,
			},
		},
		Timeout: 10 * time.Second,
	}

	ctx := context.WithValue(context.Background(), swagger.ContextBasicAuth, swagger.BasicAuth{
		UserName: sempUser,
		Password: sempPass,
	})

	// 1. Run initial file-based bootstrapping synchronously first (if enabled)
	if fileConfigured {
		logrus.WithField("category", "Bootstrap").Infof("SOL_INITIAL_STATE_FILE is configured and enabled: %s", initialStateFile)

		sempClient := swagger.NewAPIClient(swaggerConf)

		// Wait for the SEMP API to become reachable
		logrus.WithField("category", "Bootstrap").Info("Waiting for SEMP API to become reachable...")
		for {
			_, err := getAppliedVersionFromBroker(ctx, sempClient, msgVPN)
			if err == nil {
				break
			}
			if isSempNotFoundError(err) {
				break
			}
			logrus.WithField("category", "Bootstrap").Warnf("SEMP API not reachable yet, retrying in 5s... (error: %v)", err)
			time.Sleep(5 * time.Second)
		}

		// Initial execution (synchronous)
		processInitialStateFile(ctx, initialStateFile, swaggerConf, msgVPN, dryRun, certPool)

		// Start watching the file for changes (asynchronous in background)
		go watchInitialStateFile(ctx, initialStateFile, swaggerConf, msgVPN, dryRun, certPool)
	} else {
		logrus.WithField("category", "Config").Info("Initial state file loader is disabled or not configured.")
	}

	// 2. Start the messaging state consumer (if enabled)
	if messagingConfigured {
		// Create Solace client for consuming state
		solaceClient := config.NewSolaceClient(solaceConfig)
		solaceClient.OnStateReceived(func(state *config.TargetState) {
			if state == nil {
				return
			}

			logrus.WithField("category", "State").Infof("Received target state message via topic. Version: %s", state.Version)

			sempClient := swagger.NewAPIClient(swaggerConf)
			brokerVersion, err := getAppliedVersionFromBroker(ctx, sempClient, msgVPN)
			if err != nil {
				logrus.WithField("category", "VersionCheck").Errorf("Failed to check broker version: %v", err)
			}

			if brokerVersion != "" && !isVersionNewer(state.Version, brokerVersion) {
				logrus.WithFields(logrus.Fields{
					"category": "VersionCheck",
					"state":    state.Version,
					"broker":   brokerVersion,
				}).Infof("Incoming messaging state version is not newer than currently applied state on broker. Skipping.")
				return
			}

			success, _ := reconcileState(ctx, state, swaggerConf, msgVPN, dryRun, certPool)
			if success && !dryRun {
				if err := saveAppliedVersionToBroker(ctx, sempClient, msgVPN, state.Version); err != nil {
					logrus.WithField("category", "VersionCheck").Errorf("Failed to save applied version to broker: %v", err)
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
	} else {
		logrus.WithField("category", "Config").Info("Solace messaging state consumer is not configured. Running in file-only mode.")
	}

	select {}
}

func reconcileState(ctx context.Context, state *config.TargetState, swaggerConf *swagger.Configuration, msgVPN string, dryRun bool, certPool *x509.CertPool) (bool, []config.Error) {
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

		// Get Management options dynamically from viper inside the function
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
			sendStatusReport(!hadError, errorList, state, certPool)
		}
	}
	return !hadError, errorList
}

func getAppliedVersionFromBroker(ctx context.Context, client *swagger.APIClient, msgVpn string) (string, error) {
	queue, resp, err := client.QueueApi.GetMsgVpnQueue(ctx, msgVpn, "solace-dsemp-agent-metadata", nil)
	if err != nil {
		if (resp != nil && resp.StatusCode == 404) || isSempNotFoundError(err) {
			return "", nil // Queue doesn't exist yet
		}
		return "", err
	}
	return queue.Data.Owner, nil
}

func saveAppliedVersionToBroker(ctx context.Context, client *swagger.APIClient, msgVpn string, version string) error {
	_, resp, err := client.QueueApi.GetMsgVpnQueue(ctx, msgVpn, "solace-dsemp-agent-metadata", nil)
	if err != nil {
		if (resp != nil && resp.StatusCode == 404) || isSempNotFoundError(err) {
			// Create metadata queue
			queue := swagger.MsgVpnQueue{
				QueueName:      "solace-dsemp-agent-metadata",
				Owner:          version,
				IngressEnabled: boolPtrPtr(false),
				EgressEnabled:  boolPtrPtr(false),
			}
			_, _, err = client.QueueApi.CreateMsgVpnQueue(ctx, queue, msgVpn, nil)
			return err
		}
		return err
	}

	// Update queue Owner field
	queue := swagger.MsgVpnQueue{
		Owner: version,
	}
	_, _, err = client.QueueApi.UpdateMsgVpnQueue(ctx, queue, msgVpn, "solace-dsemp-agent-metadata", nil)
	return err
}

func isVersionNewer(newVer, currentVer string) bool {
	if newVer == "" {
		return false
	}
	if currentVer == "" {
		return true
	}

	// Fast path
	if newVer == currentVer {
		return false
	}

	newParts := strings.Split(newVer, ".")
	currParts := strings.Split(currentVer, ".")

	for i := 0; i < len(newParts) && i < len(currParts); i++ {
		n, errN := strconv.Atoi(newParts[i])
		c, errC := strconv.Atoi(currParts[i])

		if errN == nil && errC == nil {
			if n > c {
				return true
			} else if n < c {
				return false
			}
		} else {
			if newParts[i] > currParts[i] {
				return true
			} else if newParts[i] < currParts[i] {
				return false
			}
		}
	}

	return len(newParts) > len(currParts)
}

func processInitialStateFile(ctx context.Context, filePath string, swaggerConf *swagger.Configuration, msgVPN string, dryRun bool, certPool *x509.CertPool) {
	const category = "Bootstrap"

	data, err := os.ReadFile(filePath)
	if err != nil {
		logrus.WithField("category", category).Errorf("Failed to read initial state file %s: %v", filePath, err)
		return
	}

	var state config.TargetState
	if err := json.Unmarshal(data, &state); err != nil {
		logrus.WithField("category", category).Errorf("Failed to parse initial state JSON: %v", err)
		return
	}

	// Substitute env vars and AES encrypted strings
	config.SubstituteEnvAndAesStrings(reflect.ValueOf(&state).Elem())

	// Check currently applied version on the broker
	sempClient := swagger.NewAPIClient(swaggerConf)
	brokerVersion, err := getAppliedVersionFromBroker(ctx, sempClient, msgVPN)
	if err != nil {
		logrus.WithField("category", category).Errorf("Failed to check broker version: %v", err)
		return
	}

	if brokerVersion != "" && !isVersionNewer(state.Version, brokerVersion) {
		logrus.WithFields(logrus.Fields{
			"category": category,
			"file":     state.Version,
			"broker":   brokerVersion,
		}).Infof("Initial state file version is not newer than currently applied state on broker. Skipping bootstrap.")
		return
	}

	logrus.WithFields(logrus.Fields{
		"category": category,
		"version":  state.Version,
	}).Info("Applying initial bootstrap state file...")

	success, _ := reconcileState(ctx, &state, swaggerConf, msgVPN, dryRun, certPool)
	if success && !dryRun {
		if err := saveAppliedVersionToBroker(ctx, sempClient, msgVPN, state.Version); err != nil {
			logrus.WithField("category", category).Errorf("Failed to save applied version to broker: %v", err)
		}
	}
}

func watchInitialStateFile(ctx context.Context, filePath string, swaggerConf *swagger.Configuration, msgVPN string, dryRun bool, certPool *x509.CertPool) {
	const category = "FileWatcher"
	logrus.WithField("category", category).Infof("Starting watcher for %s (polling every 10s)", filePath)

	var lastModTime time.Time

	// Initial stat to get starting mod time
	if info, err := os.Stat(filePath); err == nil {
		lastModTime = info.ModTime()
	}

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			info, err := os.Stat(filePath)
			if err != nil {
				// File might be temporarily missing or rotated by K8s ConfigMap
				continue
			}

			if info.ModTime().After(lastModTime) {
				logrus.WithFields(logrus.Fields{
					"category": category,
					"path":     filePath,
				}).Info("Detected modification to initial state file. Re-applying...")

				lastModTime = info.ModTime()
				processInitialStateFile(ctx, filePath, swaggerConf, msgVPN, dryRun, certPool)
			}
		}
	}
}

// loadCertificates loads certificates from a file or directory
func LoadCertificates(path string, certPool *x509.CertPool) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("certificate path does not exist: %w", err)
	}

	if fileInfo.IsDir() {
		// Load all .pem and .crt files from the directory
		files, err := os.ReadDir(path)
		if err != nil {
			return fmt.Errorf("failed to read directory: %w", err)
		}

		certsLoaded := 0
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			ext := filepath.Ext(file.Name())
			if ext == ".pem" || ext == ".crt" {
				certPath := filepath.Join(path, file.Name())
				certData, err := os.ReadFile(certPath)
				if err != nil {
					logrus.WithField("category", "Config").Warnf("Failed to read certificate file %s: %v", certPath, err)
					continue
				}
				if !certPool.AppendCertsFromPEM(certData) {
					logrus.WithField("category", "Config").Warnf("Failed to parse certificates from %s", certPath)
					continue
				}
				certsLoaded++
			}
		}
		if certsLoaded == 0 {
			logrus.WithField("category", "Config").Warnf("No .pem or .crt files found in directory: %s", path)
		} else {
			logrus.WithField("category", "Config").Infof("Loaded %d certificate(s) from %s", certsLoaded, path)
		}
	} else {
		// Load single file
		certData, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read certificate file: %w", err)
		}
		if !certPool.AppendCertsFromPEM(certData) {
			return fmt.Errorf("failed to parse certificates from file")
		}
		logrus.WithField("category", "Config").Infof("Loaded certificate(s) from %s", path)
	}

	return nil
}

// createHTTPClient creates an HTTP client with optional TLS certificate validation
func createHTTPClient(validateCert bool, certPool *x509.CertPool) (*http.Client, error) {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: !validateCert,
	}

	// If certificate validation is enabled and a trust store path is provided
	if validateCert && certPool == nil {
		tlsConfig.RootCAs = certPool
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	return &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}, nil
}

func sendStatusReport(success bool, errors []config.Error, state *config.TargetState, certPool *x509.CertPool) {
	const category = "Statusreport"

	brokerURL := viper.GetString("SOL_SEMP_BROKER_URL")
	sempUser := viper.GetString("SOL_SEMP_USER")
	sempPass := viper.GetString("SOL_SEMP_PASS")
	validateCert := viper.GetBool("SOL_VALIDATE_CERT")
	trustStorePath := viper.GetString("SOL_TRUST_STORE_PATH")
	brokerVersion := semplegacy.GetBrokerVersion(brokerURL, sempUser, sempPass, validateCert, certPool)

	extraFields := viper.GetStringMapString("SOL_STATUS_EXTRA_FIELDS")
	logrus.WithField("category", "debug").Infof("brokerVersion=%v", brokerVersion)
	logrus.WithField("category", category).Infof("sending status, success=%v, errors=%d", success, len(errors))
	msg := map[string]interface{}{
		"timestamp":          time.Now().UTC().Format(time.RFC3339),
		"success":            success,
		"errors":             nil,
		"targetstateversion": state.Version,
	}
	if brokerVersion != "" {
		msg["brokerVersion"] = brokerVersion
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

		client, err := createHTTPClient(validateCert, certPool)
		if err != nil {
			logrus.WithField("category", category).Errorf("Failed to create HTTP client: %v", err)
			return
		}

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

		if err := config.SendStatusMessage(success, body, brokerURL, topic, brokerUser, brokerPass, msgVpn, validateCert, trustStorePath); err != nil {
			logrus.WithField("category", category).Errorf("Failed to send status message: %v", err)
			return
		}

		logrus.WithField("category", category).Infof("Status message sent successfully to %s on topic %s", brokerURL, topic)
	}

}

func boolPtrPtr(b bool) **bool {
	p := &b
	return &p
}

type sempErrorResponse struct {
	Meta struct {
		Error struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
			Status      string `json:"status"`
		} `json:"error"`
		ResponseCode int `json:"responseCode"`
	} `json:"meta"`
}

func isSempNotFoundError(err error) bool {
	if err == nil {
		return false
	}

	if swgErr, ok := err.(swagger.GenericSwaggerError); ok {
		var sempHdr sempErrorResponse
		if json.Unmarshal(swgErr.Body(), &sempHdr) == nil {
			if sempHdr.Meta.Error.Status == "NOT_FOUND" || sempHdr.Meta.Error.Code == 6 {
				return true
			}
		}
	}

	errStr := err.Error()
	if strings.Contains(errStr, "NOT_FOUND") || strings.Contains(errStr, "Could not find match") {
		return true
	}

	return false
}
