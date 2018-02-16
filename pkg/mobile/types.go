package mobile

import "time"

type AppConfig struct {
	DBConnectionString string
}

// ClientMetric struct is what the client payload should be parsed into
// Need to figure out how to structure this
type Metric struct {
	Timestamp time.Time   `json:"timestamp"`
	ClientID  string      `json:"clientId"`
	Data      interface{} `json:"data"`
}