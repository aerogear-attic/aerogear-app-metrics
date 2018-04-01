package setup

import (
	"github.com/aerogear/aerogear-app-metrics/pkg/config"
	"github.com/aerogear/aerogear-app-metrics/pkg/dao"
	"github.com/aerogear/aerogear-app-metrics/pkg/mobile"
	"github.com/aerogear/aerogear-app-metrics/pkg/web"
	"github.com/gorilla/mux"
)

func InitRouter(metricsDao *dao.MetricsDAO) *mux.Router {
	router := web.NewRouter()

	//metrics route
	{
		metricsService := mobile.NewMetricsService(metricsDao)
		metricsHandler := web.NewMetricsHandler(metricsService)
		web.MetricsRoute(router, metricsHandler)
	}
	//health route
	{
		healthHandler := web.NewHealthHandler(metricsDao)
		web.HealthzRoute(router, healthHandler)
	}
	return router
}

func InitDao(config config.Config) *dao.MetricsDAO {

	dbHandler := dao.DatabaseHandler{}

	err := dbHandler.Connect(config.DBConnectionString, config.DBMaxConnections)

	if err != nil {
		panic("failed to connect to sql database : " + err.Error())
	}

	if err := dbHandler.DoInitialSetup(); err != nil {
		panic("failed to perform database setup : " + err.Error())
	}

	return dao.NewMetricsDAO(dbHandler.DB)
}
