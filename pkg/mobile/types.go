package mobile

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
	ID         string `json:"id"`
	SDKVersion string `json:"sdkVersion"`
	AppVersion string `json:"appVersion"`
}

type DeviceMetric struct {
	Platform        string `json:"platform"`
	PlatformVersion string `json:"platformVersion"`
}

func (m *Metric) Validate() (valid bool, reason string) {
	if m.ClientId == "" {
		return false, "missing clientId in payload"
	}

	// check if data field was missing or empty object
	if m.Data == nil || (MetricData{}) == *m.Data {
		return false, "missing metrics data in payload"
	}
	return true, ""
}
