package status

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/GyroGearl00se/solace-dsemp-agent/pkg/semplegacy"
	"github.com/GyroGearl00se/solace-dsemp-agent/pkg/state"
)

// ReporterConfig configures status notifications via HTTP webhooks and Solace messages.
type ReporterConfig struct {
	WebhookEnabled   bool
	WebhookURL       string
	WebhookUser      string
	WebhookPass      string
	MessageEnabled   bool
	MessageBrokerURL string
	MessageTopic     string
	MessageUser      string
	MessagePass      string
	MessageMsgVpn    string
	ValidateCert     bool
	TrustStorePath   string
	SempBrokerURL    string
	SempUser         string
	SempPass         string
	ExtraFields      map[string]string
}

// SendReport creates and dispatches the status report over webhook and/or topic.
func SendReport(cfg ReporterConfig, success bool, errors []state.Error, st *state.TargetState, certPool *x509.CertPool) {
	const category = "Statusreport"

	brokerVersion := semplegacy.GetBrokerVersion(cfg.SempBrokerURL, cfg.SempUser, cfg.SempPass, cfg.ValidateCert, certPool)
	logrus.WithField("category", "debug").Infof("brokerVersion=%v", brokerVersion)
	logrus.WithField("category", category).Infof("sending status, success=%v, errors=%d", success, len(errors))

	version := ""
	if st != nil {
		version = st.Version
	}

	msg := map[string]interface{}{
		"timestamp":          time.Now().UTC().Format(time.RFC3339),
		"success":            success,
		"errors":             nil,
		"targetstateversion": version,
	}
	if brokerVersion != "" {
		msg["brokerVersion"] = brokerVersion
	}
	if len(errors) > 0 {
		msg["errors"] = errors
	}
	for k, v := range cfg.ExtraFields {
		msg[k] = v
	}

	body, _ := json.Marshal(msg)

	if cfg.WebhookEnabled && cfg.WebhookURL != "" {
		logrus.WithField("category", category).Infof("Sending status via WEBHOOK")

		tlsConfig := &tls.Config{
			InsecureSkipVerify: !cfg.ValidateCert,
		}
		if cfg.ValidateCert && certPool != nil {
			tlsConfig.RootCAs = certPool
		}

		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
			Timeout: 10 * time.Second,
		}

		req, err := http.NewRequest("POST", cfg.WebhookURL, bytes.NewBuffer(body))
		if err != nil {
			logrus.WithField("category", category).Errorf("Failed to create webhook request: %v", err)
			return
		}

		req.Header.Set("Content-Type", "application/json")
		if cfg.WebhookUser != "" && cfg.WebhookPass != "" {
			req.SetBasicAuth(cfg.WebhookUser, cfg.WebhookPass)
		}

		resp, err := client.Do(req)
		if err != nil {
			logrus.WithField("category", category).Errorf("Failed to send status webhook: %v", err)
		} else {
			resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				logrus.WithField("category", category).Warnf("Status webhook returned HTTP %d", resp.StatusCode)
			} else {
				logrus.WithField("category", category).Infof("WEBHOOK successfully sent to %s", cfg.WebhookURL)
			}
		}
	}

	if cfg.MessageEnabled {
		if cfg.MessageBrokerURL == "" || cfg.MessageTopic == "" || cfg.MessageUser == "" || cfg.MessagePass == "" || cfg.MessageMsgVpn == "" {
			logrus.WithField("category", category).Error("Status message configuration is incomplete")
			return
		}

		if err := SendStatusMessage(success, body, cfg.MessageBrokerURL, cfg.MessageTopic, cfg.MessageUser, cfg.MessagePass, cfg.MessageMsgVpn, cfg.ValidateCert, cfg.TrustStorePath); err != nil {
			logrus.WithField("category", category).Errorf("Failed to send status message: %v", err)
			return
		}

		logrus.WithField("category", category).Infof("Status message sent successfully to %s on topic %s", cfg.MessageBrokerURL, cfg.MessageTopic)
	}
}
