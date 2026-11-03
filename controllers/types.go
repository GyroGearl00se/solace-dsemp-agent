package controllers

import (
	"context"
	"path"
	"strings"

	"github.com/GyroGearl00se/solace-dsemp-agent/config"
	swagger "github.com/GyroGearl00se/solace-dsemp-agent/semp_swagger/config"
	"github.com/sirupsen/logrus"
)

// isSystemResource checks if a resource name is a system resource (starts with #)
func isSystemResource(name string) bool {
	return strings.HasPrefix(name, "#")
}

// isWhitelisted checks if a resource name matches any of the given patterns
func isWhitelisted(name string, patterns []string) bool {
	for _, pattern := range patterns {
		// Handle recursive matching with />
		if strings.HasSuffix(pattern, "/>") {
			prefix := strings.TrimSuffix(pattern, "/>")
			if strings.HasPrefix(name, prefix) {
				return true
			}
			continue
		}

		// Standard pattern matching
		matched, err := path.Match(pattern, name)
		if err == nil && matched {
			return true
		}
	}
	return false
}

// Controller defines the interface that resource controllers must implement
type ResourceController interface {
	Create(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error
	Get(ctx context.Context, client *swagger.APIClient, msgVpn string) ([]interface{}, error)
	Update(ctx context.Context, client *swagger.APIClient, msgVpn string, obj interface{}) error
	Delete(ctx context.Context, client *swagger.APIClient, msgVpn string, identifier string) error
	GetIdentifier(obj interface{}) string
	Equal(obj1, obj2 interface{}) bool
	ShouldManage(obj interface{}) bool
}

// GenericCRUDHandler provides a reusable way to handle CRUD operations
type GenericCRUDHandler struct {
	ResourceType string
	Controller   ResourceController
	GetState     func() []interface{}
}

// ProcessState handles the CRUD operations for a resource type
func (h *GenericCRUDHandler) ProcessState(ctx context.Context, client *swagger.APIClient, msgVpn string, dryRun bool, method string) []config.Error {

	var errors []config.Error

	// Get current state from broker
	brokerObjs, err := h.Controller.Get(ctx, client, msgVpn)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"category": h.ResourceType,
			"error":    err,
		}).Error("Failed to get resources from broker")
		errors = append(errors, config.Error{
			Category: h.ResourceType,
			Action:   "GET",
			Message:  err.Error(),
		})
	}

	// Create maps for easier lookup, filtering out system resources
	brokerMap := make(map[string]interface{})
	for _, obj := range brokerObjs {
		if h.Controller.ShouldManage(obj) {
			id := h.Controller.GetIdentifier(obj)
			brokerMap[id] = obj
		} else {
			logrus.WithFields(logrus.Fields{
				"category": h.ResourceType,
				"id":       h.Controller.GetIdentifier(obj),
			}).Debug("Skipping system resource or default object from broker")
		}
	}

	// Get desired state
	stateObjs := h.GetState()
	stateMap := make(map[string]interface{})
	for _, obj := range stateObjs {
		if h.Controller.ShouldManage(obj) {
			id := h.Controller.GetIdentifier(obj)
			stateMap[id] = obj
		} else {
			logrus.WithFields(logrus.Fields{
				"category": h.ResourceType,
				"id":       h.Controller.GetIdentifier(obj),
			}).Debug("Skipping system resource or default object from state")
		}
	}

	if method == "upsert" || method == "crud" {
		// Process CREATE and UPDATE operations
		for id, stateObj := range stateMap {
			if brokerObj, exists := brokerMap[id]; !exists {
				// CREATE
				if dryRun {
					logrus.WithFields(logrus.Fields{
						"category":  h.ResourceType,
						"operation": "CREATE",
						"id":        id,
					}).Info("[DRY RUN] Would create resource")
				} else {
					logrus.WithFields(logrus.Fields{
						"category":  h.ResourceType,
						"operation": "CREATE",
						"id":        id,
					}).Info("Creating new resource")
					if err := h.Controller.Create(ctx, client, msgVpn, stateObj); err != nil {
						logrus.WithFields(logrus.Fields{
							"category": h.ResourceType,
							"error":    err,
							"id":       id,
						}).Error("Failed to create resource")
						errors = append(errors, config.Error{
							Category:   h.ResourceType,
							ResourceID: id,
							Action:     "CREATE",
							Message:    err.Error(),
						})
						continue
					}
				}

			} else {
				// Check if UPDATE is needed
				if NeedsUpdate(brokerObj, stateObj) {
					if dryRun {
						logrus.WithFields(logrus.Fields{
							"category":  h.ResourceType,
							"operation": "UPDATE",
							"id":        id,
						}).Info("[DRY RUN] Would update resource")
					} else {
						logrus.WithFields(logrus.Fields{
							"category":  h.ResourceType,
							"operation": "UPDATE",
							"id":        id,
						}).Info("Updating existing resource")
						if err := h.Controller.Update(ctx, client, msgVpn, stateObj); err != nil {
							logrus.WithFields(logrus.Fields{
								"category": h.ResourceType,
								"error":    err,
								"id":       id,
							}).Error("Failed to update resource")
							errors = append(errors, config.Error{
								Category:   h.ResourceType,
								ResourceID: id,
								Action:     "UPDATE",
								Message:    err.Error(),
							})
							continue
						}
					}
				}
			}
		}
	}
	if method == "delete" || method == "crud" {
		// Process DELETE operations
		for id := range brokerMap {
			if _, exists := stateMap[id]; !exists {
				if dryRun {
					logrus.WithFields(logrus.Fields{
						"category":  h.ResourceType,
						"operation": "DELETE",
						"id":        id,
					}).Info("[DRY RUN] Would delete resource")
				} else {
					logrus.WithFields(logrus.Fields{
						"category":  h.ResourceType,
						"operation": "DELETE",
						"id":        id,
					}).Info("Deleting resource not in state")
					if err := h.Controller.Delete(ctx, client, msgVpn, id); err != nil {
						logrus.WithFields(logrus.Fields{
							"category": h.ResourceType,
							"error":    err,
							"id":       id,
						}).Error("Failed to delete resource")
						errors = append(errors, config.Error{
							Category:   h.ResourceType,
							ResourceID: id,
							Action:     "DELETE",
							Message:    err.Error(),
						})
						continue
					}
				}
			}
		}
	}

	return errors

}

// Helper function to check if an update is needed using reflection (delegates to NeedsUpdate)
func needsUpdate(broker, state interface{}) bool {
	return NeedsUpdate(broker, state)
}
