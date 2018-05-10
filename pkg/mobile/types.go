package mobile

import (
	"encoding/json"
	"fmt"
)

// swagger:ignore
type AppConfig struct {
	DBConnectionString string
}

// ClientMetric struct is what the client payload should be parsed into
// Need to figure out how to structure this
//
// swagger:model
type Metric struct {
	// required: true
	// example: 1520853523661
	ClientTimestamp json.Number `json:"timestamp,omitempty"`
	// required: true
	// example: some-unique-client-id
	ClientId string `json:"clientId"`
	// required: true
	Data *MetricData `json:"data,omitempty"`
	// required: true
	// example: init
	EventType string `json:"type"`
}

// swagger:model
type MetricData struct {
	// required: true
	App *AppMetric `json:"app,omitempty"`
	// required: true
	Device *DeviceMetric `json:"device,omitempty"`
	// required: true
	Security *SecurityMetrics `json:"security,omitempty"`
}

// swagger:model
type AppMetric struct {
	// required: true
	// example: com.example.myapp
	ID string `json:"appId"`
	// required: true
	// example: 1.0.0
	SDKVersion string `json:"sdkVersion"`
	// required: true
	// example: 2.1.0
	AppVersion string `json:"appVersion"`
	// required : true
	// example: cordova
	Framework string `json:"framework"`
}

// swagger:model
type DeviceMetric struct {
	// required: true
	// example: android
	Platform string `json:"platform"`
	// required: true
	// example: 27
	PlatformVersion string `json:"platformVersion"`
}

type SecurityMetrics []SecurityMetric

// swagger:model
type SecurityMetric struct {
	// required: true
	// example: com.example.DeveloperMode
	Id *string `json:"id,omitempty"`
	// required: true
	// example: Developer Mode
	Name *string `json:"name,omitempty"`
	// required: true
	Passed *bool `json:"passed,omitempty"`
}

const clientIdMaxLength = 128
const eventTypeMaxLength = 128
const securityMetricsMaxLength = 30

const MissingClientIdError = "Missing clientId in payload"
const MissingEventTypeError = "Missing type in payload"
const UnknownTypeError = "payload type unknown"
const MissingAppError = "Missing data.app in init-type payload"
const MissingDeviceError = "Missing data.device in init-type payload"
const MissingSecurityError = "Missing data.security in security-type payload"
const MissingDataError = "Missing metrics data in payload"
const InitMetricsIncompleteError = "data.app and data.security must both be present simultaneously"

const InvalidTimestampError = "timestamp must be a valid number"
const SecurityMetricsEmptyError = "data.security cannot be empty"
const SecurityMetricMissingIdError = "invalid element in data.security at position %v, id must be included"
const SecurityMetricMissingNameError = "invalid element in data.security at position %v, name must be included"
const SecurityMetricMissingPassedError = "invalid element in data.security at position %v, passed must be included"

var ClientIdLengthError = fmt.Sprintf("clientId exceeded maximum length of %v", clientIdMaxLength)

var EventTypeLengthError = fmt.Sprintf("type exceeded maximum length of %v", eventTypeMaxLength)

var SecurityMetricsLengthError = fmt.Sprintf("maximum length of data.security %v", securityMetricsMaxLength)

func (m *Metric) Validate() (valid bool, reason string) {
	if m.ClientId == "" {
		return false, MissingClientIdError
	}

	if len(m.ClientId) > clientIdMaxLength {
		return false, ClientIdLengthError
	}

	if m.EventType == "" {
		return false, MissingEventTypeError
	}

	if len(m.EventType) > eventTypeMaxLength {
		return false, EventTypeLengthError
	}

	if m.ClientTimestamp != "" {
		if _, err := m.ClientTimestamp.Int64(); err != nil {
			return false, InvalidTimestampError
		}
	}

	// check if data field was Missing or empty object
	if m.Data == nil || (MetricData{}) == *m.Data {
		return false, MissingDataError
	}

	switch m.EventType {
	case "init":
		return validateInitMetric(m.Data)
	case "security":
		return validateSecurityMetric(m.Data)
	default:
		return false, UnknownTypeError
	}
}

func validateInitMetric(data *MetricData) (valid bool, reason string) {
	if data.App == nil {
		return false, MissingAppError
	}
	if data.Device == nil {
		return false, MissingDeviceError
	}
	return true, ""
}

func validateSecurityMetric(data *MetricData) (valid bool, reason string) {
	// security type includes data from 'init'
	if ok, reason := validateInitMetric(data); !ok {
		return false, reason
	}
	if data.Security == nil {
		return false, MissingSecurityError
	}

	if len(*data.Security) == 0 {
		return false, SecurityMetricsEmptyError
	}
	if len(*data.Security) > securityMetricsMaxLength {
		return false, SecurityMetricsLengthError
	}
	for i, sm := range *data.Security {
		if sm.Id == nil {
			return false, fmt.Sprintf(SecurityMetricMissingIdError, i)
		}
		if sm.Name == nil {
			return false, fmt.Sprintf(SecurityMetricMissingNameError, i)
		}
		if sm.Passed == nil {
			return false, fmt.Sprintf(SecurityMetricMissingPassedError, i)
		}
	}
	return true, ""
}
