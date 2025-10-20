package config

import (
	"time"
)

type StatusReport struct {
	ConfigVersion string    `json:"configVersion"`
	Timestamp     time.Time `json:"timestamp"`
	Success       bool      `json:"success"`
	Errors        []Error   `json:"errors,omitempty"`
	Hostname      string    `json:"hostname"`
	CountryCode   string    `json:"countrycode"`
	Division      string    `json:"division"`
	Environment   string    `json:"environment"`
	Location      string    `json:"location"`
	LocationID    string    `json:"locationid"`
}

type Error struct {
	Category   string `json:"category"`
	ResourceID string `json:"resourceId"`
	Action     string `json:"action"`
	Message    string `json:"message"`
}
