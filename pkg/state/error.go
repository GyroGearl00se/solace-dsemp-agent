package state

// Error represents an operational error encountered during reconciliation.
type Error struct {
	Category   string `json:"category"`
	ResourceID string `json:"resourceId"`
	Action     string `json:"action"`
	Message    string `json:"message"`
}

