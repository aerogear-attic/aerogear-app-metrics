package main

import (
	"net/http"

	"github.com/aerogear/aerogear-app-metrics/pkg/config"
	"github.com/aerogear/aerogear-app-metrics/pkg/dao"
	"github.com/aerogear/aerogear-app-metrics/pkg/mobile"
	"github.com/aerogear/aerogear-app-metrics/pkg/web"
	log "github.com/sirupsen/logrus"
)

func main() {

	config := config.GetConfig()

	initLogger(config.LogLevel, config.LogFormat)

	dbHandler := dao.DatabaseHandler{}

	err := dbHandler.Connect(config.DBConnectionString, config.DBMaxConnections)

	if err != nil {
		panic("failed to connect to sql database : " + err.Error())
	}

	if err := dbHandler.DoInitialSetup(); err != nil {
		panic("failed to perform database setup : " + err.Error())
	}

	metricsDao := dao.NewMetricsDAO()
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

	log.WithFields(log.Fields{"listenAddress": config.ListenAddress}).Info("Starting application")

	//start
	if err := http.ListenAndServe(config.ListenAddress, router); err != nil {
		panic("failed to start " + err.Error())
	}
}

func initLogger(level, format string) {
	logLevel, err := log.ParseLevel(level)

	if err != nil {
		log.Fatalf("log level %v is not allowed. Must be one of [debug, info, warning, error, fatal, panic]", level)
		logLevel = log.InfoLevel
	}

	log.SetLevel(logLevel)

	switch format {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	case "text":
		log.SetFormatter(&log.TextFormatter{DisableColors: true})
	default:
		log.Fatalf("log format %v is not allowed. Must be one of [text, json]", format)
	}
}
