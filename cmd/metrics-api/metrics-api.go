// API for the Metrics for the AeroGear Metrics Service
//     Schemes: http
//     Title: AeroGear Metrics Service API
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Host: localhost:3000
//
//
// swagger:meta
package main

//go:generate swagger generate spec -m -o ../../swagger.json

import (
	"net/http"

	"github.com/aerogear/aerogear-app-metrics/internal/setup"
	"github.com/aerogear/aerogear-app-metrics/pkg/config"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

func main() {
	config := config.GetConfig()
	metricsDao := setup.InitDao(config.DBConnectionString, config.DBMaxConnections)
	defer metricsDao.Close()
	router := setup.InitRouter(metricsDao)

	initLogger(config.LogLevel, config.LogFormat)

	log.WithFields(log.Fields{"listenAddress": config.ListenAddress}).Info("Starting application")

	// allow CORS for localhost
	handler := cors.New(cors.Options{
		AllowCredentials: true,
	}).Handler(router)

	//start
	if err := http.ListenAndServe(config.ListenAddress, handler); err != nil {
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
