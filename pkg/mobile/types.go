package mobile

import (
	"fmt"
)

type AppConfig struct {
	DBConnectionString string
}

// ClientMetric struct is what the client payload should be parsed into
// Need to figure out how to structure this
type Metric struct {
	ClientTimestamp int64       `json:"timestamp,omitempty"`
	ClientId        string      `json:"clientId"`
	Data            *MetricData `json:"data,omitempty"`
}

type MetricData struct {
	App    *AppMetric    `json:"app,omitempty"`
	Device *DeviceMetric `json:"device,omitempty"`
}

type AppMetric struct {
	ID         string `json:"appId"`
	SDKVersion string `json:"sdkVersion"`
	AppVersion string `json:"appVersion"`
}

type DeviceMetric struct {
	Platform        string `json:"platform"`
	PlatformVersion string `json:"platformVersion"`
}

const clientIdMaxLength = 128

var clientIdLengthError = fmt.Sprintf("clientId exceeded maximum length of %v", clientIdMaxLength)

func (m *Metric) Validate() (valid bool, reason string) {
	if m.ClientId == "" {
		return false, "missing clientId in payload"
	}

	if len(m.ClientId) > clientIdMaxLength {
		return false, clientIdLengthError
	}

	// check if data field was missing or empty object
	if m.Data == nil || (MetricData{}) == *m.Data {
		return false, "missing metrics data in payload"
	}
	return true, ""
}
