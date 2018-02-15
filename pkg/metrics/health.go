package metrics

import (
	"github.com/aerogear/aerogear-metrics-api/pkg/dao"
)

type Health struct {
	mdao dao.MetricsDAO
}

func (h Health) IsHealthy() bool {
	return true
}
