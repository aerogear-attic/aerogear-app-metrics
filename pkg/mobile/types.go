package mobile

import (
	"encoding/json"
	"fmt"
)

type AppConfig struct {
	DBConnectionString string
}

// ClientMetric struct is what the client payload should be parsed into
// Need to figure out how to structure this
type Metric struct {
	ClientTimestamp json.Number `json:"timestamp,omitempty"`
	ClientId        string      `json:"clientId"`
	Data            *MetricData `json:"data,omitempty"`
}

type MetricData struct {
	App      *AppMetric       `json:"app,omitempty"`
	Device   *DeviceMetric    `json:"device,omitempty"`
	Security *SecurityMetrics `json:"security,omitempty"`
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

type SecurityMetrics struct {
	EmulatorCheck      *SecurityMetric `json:"org.aerogear.mobile.security.checks.EmulatorCheck,omitempty,string"`
	DeveloperModeCheck *SecurityMetric `json:"org.aerogear.mobile.security.checks.DeveloperModeCheck,omitempty,string"`
	DebuggerCheck      *SecurityMetric `json:"org.aerogear.mobile.security.checks.DebuggerCheck,omitempty,string"`
	RootedCheck        *SecurityMetric `json:"org.aerogear.mobile.security.checks.RootedCheck,omitempty,string"`
	ScreenLockCheck    *SecurityMetric `json:"org.aerogear.mobile.security.checks.ScreenLockCheck,omitempty,string"`
}

type SecurityMetric struct {
	Passed bool `json:"passed"`
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

	if m.ClientTimestamp != "" {
		if _, err := m.ClientTimestamp.Int64(); err != nil {
			return false, "timestamp must be a valid number"
		}
	}

	// check if data field was missing or empty object
	if m.Data == nil || (MetricData{}) == *m.Data {
		return false, "missing metrics data in payload"
	}
	return true, ""
}
