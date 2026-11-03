package reconciler

import (
	"context"
	"strconv"
	"strings"

	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
	"github.com/GyroGearl00se/solace-dsemp-agent/pkg/semputil"
)

const MetadataQueueName = "solace-dsemp-agent-metadata"

// GetAppliedVersionFromBroker queries the metadata queue to read the currently applied version.
func GetAppliedVersionFromBroker(ctx context.Context, client *swagger.APIClient, msgVpn string) (string, error) {
	queue, resp, err := client.QueueApi.GetMsgVpnQueue(ctx, msgVpn, MetadataQueueName, nil)
	if err != nil {
		if (resp != nil && resp.StatusCode == 404) || semputil.IsNotFoundError(err) {
			return "", nil // Metadata queue doesn't exist yet
		}
		return "", err
	}
	return queue.Data.Owner, nil
}

// SaveAppliedVersionToBroker records the state version in the metadata queue's Owner field.
func SaveAppliedVersionToBroker(ctx context.Context, client *swagger.APIClient, msgVpn string, version string) error {
	_, resp, err := client.QueueApi.GetMsgVpnQueue(ctx, msgVpn, MetadataQueueName, nil)
	if err != nil {
		if (resp != nil && resp.StatusCode == 404) || semputil.IsNotFoundError(err) {
			queue := swagger.MsgVpnQueue{
				QueueName:      MetadataQueueName,
				Owner:          version,
				IngressEnabled: semputil.BoolPtrPtr(false),
				EgressEnabled:  semputil.BoolPtrPtr(false),
			}
			_, _, err = client.QueueApi.CreateMsgVpnQueue(ctx, queue, msgVpn, nil)
			return err
		}
		return err
	}

	queue := swagger.MsgVpnQueue{
		Owner: version,
	}
	_, _, err = client.QueueApi.UpdateMsgVpnQueue(ctx, queue, msgVpn, MetadataQueueName, nil)
	return err
}

// IsVersionNewer checks if newVer > currentVer using dot-separated semantic comparisons.
func IsVersionNewer(newVer, currentVer string) bool {
	if newVer == "" {
		return false
	}
	if currentVer == "" {
		return true
	}
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
