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

	db, err := dao.Connect(config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.SSLMode)

	if err != nil {
		panic("failed to connect to sql database : " + err.Error())
	}
	defer db.Close()

	if err := dao.DoInitialSetup(); err != nil {
		panic("failed to perform database setup : " + err.Error())
	}

	metricsDao := dao.NewMetricsDAO(db)
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

	listenAddress := ":3000"

	log.Printf("Starting application... going to listen on %v", listenAddress)

	//start
	if err := http.ListenAndServe(listenAddress, router); err != nil {
		panic("failed to start " + err.Error())
	}
}
