package main

import (
	"context"
	"crypto/x509"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/GyroGearl00se/solace-dsemp-agent/pkg/httplistener"
	"github.com/GyroGearl00se/solace-dsemp-agent/pkg/messaging"
	"github.com/GyroGearl00se/solace-dsemp-agent/pkg/reconciler"
	"github.com/GyroGearl00se/solace-dsemp-agent/pkg/semputil"
	"github.com/GyroGearl00se/solace-dsemp-agent/pkg/state"
	"github.com/GyroGearl00se/solace-dsemp-agent/pkg/status"
	"github.com/GyroGearl00se/solace-dsemp-agent/pkg/watcher"
	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
)

func main() {
	// Initialize configuration
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	_ = viper.ReadInConfig()
	viper.AutomaticEnv()

	// Configure logger
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	viper.SetDefault("LOGLEVEL", 4)
	logrus.SetLevel(logrus.Level(viper.GetInt("LOGLEVEL")))

	logrus.WithField("category", "Info").Info("Solace DSEMP Agent starting...")

	// Broker SEMP Configuration
	brokerURL := viper.GetString("SOL_SEMP_BROKER_URL")
	sempUser := viper.GetString("SOL_SEMP_USER")
	sempPass := viper.GetString("SOL_SEMP_PASS")
	if brokerURL == "" {
		logrus.WithField("category", "Config").Fatal("SOL_SEMP_BROKER_URL must be set!")
	}
	if sempUser == "" || sempPass == "" {
		logrus.WithField("category", "Config").Fatal("SOL_SEMP_USER and SOL_SEMP_PASS must be set!")
	}

	viper.SetDefault("SOL_SEMP_MSG_VPN", "default")
	msgVPN := viper.GetString("SOL_SEMP_MSG_VPN")

	viper.SetDefault("SOL_VALIDATE_CERT", true)
	validateCert := viper.GetBool("SOL_VALIDATE_CERT")
	trustStorePath := viper.GetString("SOL_TRUST_STORE_PATH")

	viper.SetDefault("SOL_VERSION_CHECK_ENABLED", true)
	versionCheckEnabled := viper.GetBool("SOL_VERSION_CHECK_ENABLED")

	viper.SetDefault("SOL_DRYRUN", false)
	dryRun := viper.GetBool("SOL_DRYRUN")

	aesKey := viper.GetString("SOL_AES_KEY")

	// Load TLS certificates
	certPool := x509.NewCertPool()
	if trustStorePath != "" {
		if err := semputil.LoadCertificates(trustStorePath, certPool); err != nil {
			logrus.WithField("category", "Config").Fatalf("Failed to load certificates: %v", err)
		}
	}

	// SEMP Swagger client
	swaggerConf, err := semputil.NewSwaggerConfig(brokerURL, validateCert, certPool)
	if err != nil {
		logrus.WithField("category", "Config").Fatalf("Failed to initialize Swagger configuration: %v", err)
	}
	sempClient := swagger.NewAPIClient(swaggerConf)

	ctx, cancel := context.WithCancel(context.WithValue(context.Background(), swagger.ContextBasicAuth, swagger.BasicAuth{
		UserName: sempUser,
		Password: sempPass,
	}))
	defer cancel()

	// Setup graceful shutdown handler
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigChan
		logrus.WithField("category", "Lifecycle").Infof("Received signal %v. Initiating shutdown...", sig)
		cancel()
	}()

	// Status reporter setup
	statusReporterCfg := status.ReporterConfig{
		WebhookEnabled:   viper.GetBool("SOL_STATUS_WEBHOOK_ENABLED"),
		WebhookURL:       viper.GetString("SOL_STATUS_WEBHOOK_URL"),
		WebhookUser:      viper.GetString("SOL_STATUS_WEBHOOK_USER"),
		WebhookPass:      viper.GetString("SOL_STATUS_WEBHOOK_PASS"),
		MessageEnabled:   viper.GetBool("SOL_STATUS_MESSAGE_ENABLED"),
		MessageBrokerURL: viper.GetString("SOL_STATUS_MESSAGE_BROKER_URL"),
		MessageTopic:     viper.GetString("SOL_STATUS_MESSAGE_TOPIC"),
		MessageUser:      viper.GetString("SOL_STATUS_MESSAGE_BROKER_USER"),
		MessagePass:      viper.GetString("SOL_STATUS_MESSAGE_BROKER_PASS"),
		MessageMsgVpn:    viper.GetString("SOL_STATUS_MESSAGE_MSG_VPN"),
		ValidateCert:     validateCert,
		TrustStorePath:   trustStorePath,
		SempBrokerURL:    brokerURL,
		SempUser:         sempUser,
		SempPass:         sempPass,
		ExtraFields:      viper.GetStringMapString("SOL_STATUS_EXTRA_FIELDS"),
	}

	// Management flags
	flags := reconciler.ManagementFlags{
		MsgVpns:                 viper.GetBool("SOL_MANAGE_MSG_VPNS"),
		Queues:                  viper.GetBool("SOL_MANAGE_QUEUES"),
		ACLProfiles:             viper.GetBool("SOL_MANAGE_ACL_PROFILES"),
		ClientUsernames:         viper.GetBool("SOL_MANAGE_CLIENT_USERNAMES"),
		ClientProfiles:          viper.GetBool("SOL_MANAGE_CLIENT_PROFILES"),
		Bridges:                 viper.GetBool("SOL_MANAGE_BRIDGES"),
		DMRBridges:              viper.GetBool("SOL_MANAGE_DMR_BRIDGES"),
		BridgeRemoteMsgVpns:     viper.GetBool("SOL_MANAGE_BRIDGE_REMOTE_MSG_VPNS"),
		JndiConnectionFactories: viper.GetBool("SOL_MANAGE_JNDI_CONNECTION_FACTORIES"),
		JndiQueues:              viper.GetBool("SOL_MANAGE_JNDI_QUEUES"),
		JndiTopics:              viper.GetBool("SOL_MANAGE_JNDI_TOPICS"),
		Proxies:                 viper.GetBool("SOL_MANAGE_PROXIES"),
		RestDeliveryPoints:      viper.GetBool("SOL_MANAGE_REST_DELIVERY_POINTS"),
		QueueTemplates:          viper.GetBool("SOL_MANAGE_QUEUE_TEMPLATES"),
		TopicEndpoints:          viper.GetBool("SOL_MANAGE_TOPIC_ENDPOINTS"),
		TopicEndpointTemplates:  viper.GetBool("SOL_MANAGE_TOPIC_ENDPOINT_TEMPLATES"),
	}

	engine := reconciler.NewEngine(sempClient, msgVPN, dryRun, flags)

	// Central dispatch function
	applyState := func(source string, target *state.TargetState) {
		if target == nil {
			return
		}

		brokerVersion, err := reconciler.GetAppliedVersionFromBroker(ctx, sempClient, msgVPN)
		if err != nil {
			logrus.WithField("category", "VersionCheck").Errorf("Failed to check broker version: %v", err)
		}
		if versionCheckEnabled && brokerVersion != "" && !reconciler.IsVersionNewer(target.Version, brokerVersion) {
			logrus.WithFields(logrus.Fields{
				"category": "VersionCheck",
				"source":   source,
				"state":    target.Version,
				"broker":   brokerVersion,
			}).Infof("Incoming state version is not newer than currently applied state on broker. Skipping.")
			return
		}

		success, errors := engine.Reconcile(ctx, target)
		if success && !dryRun {
			if err := reconciler.SaveAppliedVersionToBroker(ctx, sempClient, msgVPN, target.Version); err != nil {
				logrus.WithField("category", "VersionCheck").Errorf("Failed to save applied version to broker: %v", err)
			}
		}

		if statusReporterCfg.WebhookEnabled || statusReporterCfg.MessageEnabled {
			status.SendReport(statusReporterCfg, success, errors, target, certPool)
		}
	}

	// Solace Messaging Configuration
	consumerConfig := messaging.ConsumerConfig{
		URL:            viper.GetString("SOL_STATE_BROKER_URL"),
		VPN:            viper.GetString("SOL_STATE_MSG_VPN"),
		Username:       viper.GetString("SOL_STATE_USERNAME"),
		Password:       viper.GetString("SOL_STATE_PASSWORD"),
		Topic:          viper.GetString("SOL_STATE_TOPIC"),
		TrustStorePath: trustStorePath,
		ValidateCert:   validateCert,
		AESKey:         aesKey,
	}

	viper.SetDefault("SOL_STATE_CONSUMER_ENABLED", true)
	stateConsumerEnabled := viper.GetBool("SOL_STATE_CONSUMER_ENABLED")

	viper.SetDefault("SOL_INITIAL_STATE_ENABLED", true)
	initialStateEnabled := viper.GetBool("SOL_INITIAL_STATE_ENABLED")
	initialStateFile := viper.GetString("SOL_INITIAL_STATE_FILE")

	hasMessaging := consumerConfig.URL != "" && consumerConfig.Username != "" && consumerConfig.Password != "" && consumerConfig.Topic != ""
	messagingConfigured := stateConsumerEnabled && hasMessaging
	fileConfigured := initialStateEnabled && initialStateFile != ""

	viper.SetDefault("SOL_HTTP_LISTENER_ENABLED", false)
	httpListenerEnabled := viper.GetBool("SOL_HTTP_LISTENER_ENABLED")

	if !messagingConfigured && !fileConfigured && !httpListenerEnabled {
		logrus.WithField("category", "Config").Fatal("No active configuration source was provided! Configure messaging, SOL_INITIAL_STATE_FILE, or SOL_HTTP_LISTENER_ENABLED.")
	}

	// 1. Initial State File Loader & Watcher
	if fileConfigured {
		logrus.WithField("category", "Bootstrap").Infof("SOL_INITIAL_STATE_FILE configured: %s", initialStateFile)

		// Wait for SEMP API to become reachable
		logrus.WithField("category", "Bootstrap").Info("Waiting for SEMP API to become reachable...")
		for {
			_, err := reconciler.GetAppliedVersionFromBroker(ctx, sempClient, msgVPN)
			if err == nil || semputil.IsNotFoundError(err) {
				break
			}
			logrus.WithField("category", "Bootstrap").Warnf("SEMP API not reachable yet, retrying in 5s... (%v)", err)
			select {
			case <-ctx.Done():
				return
			case <-time.After(5 * time.Second):
			}
		}

		fw := watcher.NewFileWatcher(initialStateFile, aesKey, 10*time.Second, func(st *state.TargetState) {
			applyState("FileWatcher", st)
		})

		// Synchronous initial apply
		if st, err := fw.ReadState(); err != nil {
			logrus.WithField("category", "Bootstrap").Errorf("Failed to read initial state file %s: %v", initialStateFile, err)
		} else {
			applyState("Bootstrap", st)
		}

		// Background watcher
		go fw.Watch(ctx)
	}

	// 2. Solace Messaging State Consumer
	if messagingConfigured {
		reconnectInterval := 5 * time.Second
		if val := viper.GetString("SOL_RECONNECT_INTERVAL"); val != "" {
			if d, err := time.ParseDuration(val); err == nil {
				reconnectInterval = d
			}
		}

		consumer := messaging.NewConsumer(consumerConfig)
		consumer.OnStateReceived(func(st *state.TargetState) {
			logrus.WithField("category", "State").Infof("Received target state message via topic. Version: %s", st.Version)
			applyState("Messaging", st)
		})

		for {
			logrus.WithField("category", "Solace state consumer").Info("Attempting to connect to Solace broker...")
			if err := consumer.Connect(); err != nil {
				logrus.WithField("category", "Solace state consumer").Errorf("Failed to connect: %v, retrying in %v...", err, reconnectInterval)
				select {
				case <-ctx.Done():
					return
				case <-time.After(reconnectInterval):
					continue
				}
			}
			logrus.WithField("category", "Solace state consumer").Info("Successfully connected to Solace broker")
			break
		}
		defer consumer.Disconnect()
	}

	logrus.WithField("category", "Lifecycle").Info("Solace DSEMP Agent is active and running.")

	// 3. Optional HTTP listener
	if httpListenerEnabled {
		viper.SetDefault("SOL_HTTP_LISTENER_ADDR", ":9001")
		httpSrv := httplistener.NewServer(httplistener.Config{
			Enabled:     true,
			Address:     viper.GetString("SOL_HTTP_LISTENER_ADDR"),
			TLSEnabled:  viper.GetBool("SOL_HTTP_TLS_ENABLED"),
			TLSCertFile: viper.GetString("SOL_HTTP_TLS_CERT_FILE"),
			TLSKeyFile:  viper.GetString("SOL_HTTP_TLS_KEY_FILE"),
			Username:    viper.GetString("SOL_HTTP_LISTENER_USER"),
			Password:    viper.GetString("SOL_HTTP_LISTENER_PASS"),
			AESKey:      aesKey,
		}, func(st *state.TargetState) {
			applyState("HTTP", st)
		})
		go httpSrv.Start(ctx)
	}

	<-ctx.Done()
	logrus.WithField("category", "Lifecycle").Info("Solace DSEMP Agent exited cleanly.")
}
