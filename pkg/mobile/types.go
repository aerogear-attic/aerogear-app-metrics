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

type SecurityMetrics []SecurityMetric

const clientIdMaxLength = 128
const eventTypeMaxLength = 128

const missingClientIdError = "missing clientId in payload"
const missingDataError = "missing metrics data in payload"
const securityMetricsEmptyError = "data.security cannot be empty"

const invalidTimestampError = "timestamp must be a valid number"

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

	if m.Data.App != nil {
		if valid, reason := m.Data.App.Validate(); !valid {
			return valid, reason
		}
	}

	if m.Data.Device != nil {
		if valid, reason := m.Data.Device.Validate(); !valid {
			return valid, reason
		}
	}

	if m.Data.Security != nil {
		if len(*m.Data.Security) == 0 {
			return false, securityMetricsEmptyError
		}
		if len(*m.Data.Security) > securityMetricsMaxLength {
			return false, securityMetricsLengthError
		}
		for i, securityCheck := range *m.Data.Security {
			if valid, reason := securityCheck.Validate(i); !valid {
				return valid, reason
			}
		}
	}
	return true, ""
}
