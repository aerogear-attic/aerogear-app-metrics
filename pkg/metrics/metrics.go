package metrics

import (
	"errors"

	"github.com/aerogear/aerogear-metrics-api/pkg/dao"
	"github.com/aerogear/aerogear-metrics-api/pkg/models"
)

type Metrics struct {
	mdao dao.MetricsDAO
}

func (m Metrics) Create(metric models.Metric) (models.Metric, error) {
	return metric, errors.New("This is an error")
}
