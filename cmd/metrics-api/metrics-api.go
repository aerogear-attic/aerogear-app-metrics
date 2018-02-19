package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aerogear/aerogear-metrics-api/pkg/dao"
	"github.com/aerogear/aerogear-metrics-api/pkg/mobile"
	"github.com/aerogear/aerogear-metrics-api/pkg/web"
)

func main() {

	// Simple helper function to read an environment or return a default value
	getEnv := func(key string, defaultVal string) string {
		if value := os.Getenv(key); value != "" {
			return value
		}
		return defaultVal
	}

	var (
		dbHost     = getEnv("PGHOST", "localhost")
		dbUser     = getEnv("PGPORT", "postgres")
		dbPassword = getEnv("PGPASSWORD", "postgres")
		dbName     = getEnv("PGDATABASE", "aerogear_mobile_metrics")
		sslMode    = getEnv("PGSSLMODE", "disable")
	)

	db, err := dao.Connect(dbHost, dbUser, dbPassword, dbName, sslMode)

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
		healthHandler := web.NewHealthHandler()
		web.HealthzRoute(router, healthHandler)
	}

	listenAddress := ":3000"

	log.Printf("Starting application... going to listen on %v", listenAddress)

	//start
	if err := http.ListenAndServe(listenAddress, router); err != nil {
		panic("failed to start " + err.Error())
	}
}
