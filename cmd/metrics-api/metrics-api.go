package main

import (
	"log"
	"net/http"

	"github.com/aerogear/aerogear-metrics-api/pkg/config"
	"github.com/aerogear/aerogear-metrics-api/pkg/dao"
	"github.com/aerogear/aerogear-metrics-api/pkg/mobile"
	"github.com/aerogear/aerogear-metrics-api/pkg/web"
)

func main() {

	config := config.GetConfig()

	dbHandler := dao.DatabaseHandler{}

	err := dbHandler.Connect(config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.SSLMode)

	if err != nil {
		panic("failed to connect to sql database : " + err.Error())
	}
	defer dbHandler.DB.Close()

	if err := dbHandler.DoInitialSetup(); err != nil {
		panic("failed to perform database setup : " + err.Error())
	}

	metricsDao := dao.NewMetricsDAO(dbHandler.DB)
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

	log.Printf("Starting application... going to listen on %v", config.ListenAddress)

	//start
	if err := http.ListenAndServe(config.ListenAddress, router); err != nil {
		panic("failed to start " + err.Error())
	}
}
