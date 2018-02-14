package main

import (
	"net/http"

	"github.com/aerogear/aerogear-metrics-api/pkg/handlers"
	"github.com/aerogear/aerogear-metrics-api/pkg/metrics"
	"github.com/aerogear/aerogear-metrics-api/pkg/web"
)

func main() {

	app := metrics.NewApp(metrics.AppConfig{
		DBConnectionString: "localhost", // refactor this into env var
	})

	h := handlers.BuildHandlers(app)

	router := web.NewRouter(web.Config{
		Routes: h.Routes,
	})

	port := ":3001"
	http.ListenAndServe(port, router)
}
