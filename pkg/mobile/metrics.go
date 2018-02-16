package mobile

import (
	"errors"

)

type MetricsService struct {
	mdao MetricCreator
}

func NewMetricsService(dao MetricCreator)*MetricsService{
	return &MetricsService{mdao:dao}
}



func (m MetricsService) Create(metric Metric) (Metric, error) {
	return metric, errors.New("This is an error")
}
