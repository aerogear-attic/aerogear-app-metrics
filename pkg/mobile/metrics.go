package mobile

import "encoding/json"

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

	return metric, m.mdao.Create(metric.ClientId, metricsData)
}
