package models

import "time"

// Metric struct is what the client payload should be parsed into
// Need to figure out how to structure this
type Metric struct {
	Timestamp time.Time   `json:"timestamp"`
	ClientID  string      `json:"clientId"`
	Data      interface{} `json:"data"`
}
