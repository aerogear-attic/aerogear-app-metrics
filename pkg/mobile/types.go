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

type SecurityMetrics []SecurityMetric

type SecurityMetric struct {
	Id     *string `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Passed *bool   `json:"passed,omitempty"`
}

const clientIdMaxLength = 128
const securityMetricsMaxLength = 30

const missingClientIdError = "missing clientId in payload"
const invalidTimestampError = "timestamp must be a valid number"
const missingDataError = "missing metrics data in payload"
const securityMetricsEmptyError = "data.security cannot be empty"
const securityMetricMissingIdError = "invalid element in data.security at position %v, id must be included"
const securityMetricMissingNameError = "invalid element in data.security at position %v, name must be included"
const securityMetricMissingPassedError = "invalid element in data.security at position %v, passed must be included"

var clientIdLengthError = fmt.Sprintf("clientId exceeded maximum length of %v", clientIdMaxLength)
var securityMetricsLengthError = fmt.Sprintf("maximum length of data.security %v", securityMetricsMaxLength)

func (m *Metric) Validate() (valid bool, reason string) {
	if m.ClientId == "" {
		return false, missingClientIdError
	}

	if len(m.ClientId) > clientIdMaxLength {
		return false, clientIdLengthError
	}

	if m.ClientTimestamp != "" {
		if _, err := m.ClientTimestamp.Int64(); err != nil {
			return false, invalidTimestampError
		}
	}

	// check if data field was missing or empty object
	if m.Data == nil || (MetricData{}) == *m.Data {
		return false, missingDataError
	}

	if m.Data.Security != nil {
		if len(*m.Data.Security) == 0 {
			return false, securityMetricsEmptyError
		}
		if len(*m.Data.Security) > securityMetricsMaxLength {
			return false, securityMetricsLengthError
		}
		for i, sm := range *m.Data.Security {
			if sm.Id == nil {
				return false, fmt.Sprintf(securityMetricMissingIdError, i)
			}
			if sm.Name == nil {
				return false, fmt.Sprintf(securityMetricMissingNameError, i)
			}
			if sm.Passed == nil {
				return false, fmt.Sprintf(securityMetricMissingPassedError, i)
			}
		}
	}
	return true, ""
}
