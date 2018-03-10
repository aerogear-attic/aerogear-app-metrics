package mobile

import (
	"time"
)

type MetricsService struct {
	mdao MetricCreator
}

func NewMetricsService(dao MetricCreator) *MetricsService {
	return &MetricsService{mdao: dao}
}

func (m MetricsService) Create(metric Metric) (Metric, error) {
	t, err := metric.ClientTimestamp.Int64()

	// happens if timestamp is empty
	if err != nil {
		return metric, m.mdao.Create(metric.ClientId, metric, nil)
	}

	// convert to time object
	clientTime := time.Unix(t, 0)

	return metric, m.mdao.Create(metric.ClientId, metric, &clientTime)
}
