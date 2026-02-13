package semplegacy

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/xml"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func GetBrokerVersion(brokerURL, sempUser, sempPass string, validateCert bool, trustStorePath string) string {
	const category = "Statusreport"

	sempURL := brokerURL + "/SEMP"
	xmlPayload := `<rpc><show><version/></show></rpc>`

	req, err := http.NewRequest("POST", sempURL, bytes.NewBuffer([]byte(xmlPayload)))
	if err != nil {
		logrus.WithField("category", category).Errorf("Failed to create SEMP request for version: %v", err)
		return ""
	}

	req.SetBasicAuth(sempUser, sempPass)
	req.Header.Set("Content-Type", "application/xml")

	tlsConfig := &tls.Config{
		InsecureSkipVerify: !validateCert,
	}

	if validateCert && trustStorePath != "" {
		certData, err := os.ReadFile(trustStorePath)
		if err != nil {
			logrus.WithField("category", category).Errorf("Failed to read trust store: %v", err)
			return ""
		}

		certPool := x509.NewCertPool()
		if !certPool.AppendCertsFromPEM(certData) {
			logrus.WithField("category", category).Errorf("Failed to parse trust store certificates")
			return ""
		}

		tlsConfig.RootCAs = certPool
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	sempClient := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}
	resp, err := sempClient.Do(req)
	if err != nil {
		logrus.WithField("category", category).Errorf("Failed to execute SEMP request for version: %v", err)
		return ""
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.WithField("category", category).Errorf("Failed to read SEMP response body for version: %v", err)
		return ""
	}

	var rpcReply struct {
		Rpc struct {
			Show struct {
				Version struct {
					Description string `xml:"description"`
				} `xml:"version"`
			} `xml:"show"`
		} `xml:"rpc"`
	}

	if err := xml.Unmarshal(bodyBytes, &rpcReply); err != nil {
		logrus.WithField("category", category).Errorf("Failed to unmarshal SEMP version response: %v", err)
		return ""
	}

	return rpcReply.Rpc.Show.Version.Description
}
