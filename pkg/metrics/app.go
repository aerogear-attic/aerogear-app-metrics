package metrics

import (
	"github.com/aerogear/aerogear-metrics-api/pkg/dao"
	"github.com/aerogear/aerogear-metrics-api/pkg/models"
)

func NewApp(config AppConfig) models.App {
	mdao := dao.GetMetricsDAO(config.DBConnectionString)
	mdao.Connect()
	return models.App{
		Metrics: Metrics{
			mdao: mdao,
		},
		Health: Health{
			mdao: mdao,
		},
	}
}
