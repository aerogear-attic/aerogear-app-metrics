package mobile

type AppConfig struct {
	DBConnectionString string
}

// ClientMetric struct is what the client payload should be parsed into
// Need to figure out how to structure this
type Metric struct {
	ClientTimestamp int64      `json:"timestamp"`
	ClientId        string     `json:"clientId"`
	Data            MetricData `json:"data"`
}

type MetricData struct {
	App    AppMetric    `json:"app"`
	Device DeviceMetric `json:"device"`
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
