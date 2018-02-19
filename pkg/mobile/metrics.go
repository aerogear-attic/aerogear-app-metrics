package mobile

type MetricsService struct {
	mdao MetricCreator
}

func NewMetricsService(dao MetricCreator) *MetricsService {
	return &MetricsService{mdao: dao}
}

func (m MetricsService) Create(metric Metric) (Metric, error) {
	return metric, nil
}
