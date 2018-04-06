package mobile

import "time"

// MetricCreator defines how a metric can be created
type MetricCreator interface {
	Create(clientId string, eventType string, metricsData []byte, clientTime *time.Time) error
}
