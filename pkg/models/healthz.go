package models

import "time"

// Healthz struct is used to return a nice JSON object from the API
// When the healthcheck passes or fails e.g.
// {"time": "123456789", "status": "ok"}
type Healthz struct {
	Timestamp time.Time `json:"time"`
	Status    string    `json:"status"`
}
