package config

type Error struct {
	Category   string `json:"category"`
	ResourceID string `json:"resourceId"`
	Action     string `json:"action"`
	Message    string `json:"message"`
}
