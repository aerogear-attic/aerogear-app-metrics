package mobile

import "time"

// MetricCreator defines how a metric can be created
type MetricCreator interface {
	Create(clientId string, metric Metric, clientTime *time.Time) error
}
