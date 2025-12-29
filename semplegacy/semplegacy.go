package semplegacy

import (
	"bytes"
	"encoding/xml"
	"io"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func GetBrokerVersion(brokerURL, sempUser, sempPass string) string {
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

	sempClient := &http.Client{Timeout: 10 * time.Second}
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
