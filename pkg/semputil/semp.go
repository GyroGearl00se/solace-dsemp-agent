package semputil

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
	"github.com/sirupsen/logrus"
)

// SempErrorResponse models the error response structure from SEMP v2.
type SempErrorResponse struct {
	Meta struct {
		Error struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
			Status      string `json:"status"`
		} `json:"error"`
		ResponseCode int `json:"responseCode"`
	} `json:"meta"`
}

// IsNotFoundError checks if an error returned by SEMP is a 404/NOT_FOUND.
func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}

	if swgErr, ok := err.(swagger.GenericSwaggerError); ok {
		var sempHdr SempErrorResponse
		if json.Unmarshal(swgErr.Body(), &sempHdr) == nil {
			if sempHdr.Meta.Error.Status == "NOT_FOUND" || sempHdr.Meta.Error.Code == 6 {
				return true
			}
		}
	}

	errStr := err.Error()
	return strings.Contains(errStr, "NOT_FOUND") || strings.Contains(errStr, "Could not find match")
}

// BoolPtrPtr returns a pointer to a pointer of a boolean, matching Swagger's **bool.
func BoolPtrPtr(b bool) **bool {
	p := &b
	return &p
}

// NewSwaggerConfig constructs a configured swagger.Configuration with TLS and timeouts.
func NewSwaggerConfig(brokerURL string, validateCert bool, certPool *x509.CertPool) (*swagger.Configuration, error) {
	swaggerConfHost, err := url.Parse(brokerURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse broker URL %q: %w", brokerURL, err)
	}

	swaggerConf := swagger.NewConfiguration()
	swaggerConf.Host = swaggerConfHost.Hostname()
	swaggerConf.BasePath = strings.TrimRight(brokerURL, "/") + "/SEMP/v2/config"
	swaggerConf.Scheme = swaggerConfHost.Scheme

	swaggerConf.HTTPClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: !validateCert,
				RootCAs:            certPool,
			},
		},
		Timeout: 10 * time.Second,
	}

	return swaggerConf, nil
}

// LoadCertificates loads certificates from a file or directory into a cert pool.
func LoadCertificates(path string, certPool *x509.CertPool) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("certificate path does not exist: %w", err)
	}

	if fileInfo.IsDir() {
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

