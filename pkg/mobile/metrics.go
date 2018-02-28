package mobile

import (
	"encoding/json"
	"time"
)

type MetricsService struct {
	mdao MetricCreator
}

func NewMetricsService(dao MetricCreator) *MetricsService {
	return &MetricsService{mdao: dao}
}

func (m MetricsService) Create(metric Metric) (Metric, error) {
	metricsData, err := json.Marshal(metric.Data)

	if err != nil {
		return metric, err
	}

	t, err := metric.ClientTimestamp.Int64()

	// happens timestamp is empty
	if err != nil {
		return metric, m.mdao.Create(metric.ClientId, metricsData, nil)
	}

	// convert to time object
	clientTime := time.Unix(t, 0)

	return metric, m.mdao.Create(metric.ClientId, metricsData, &clientTime)
}
