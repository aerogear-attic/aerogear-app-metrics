package web

import (
	"net/http"

	"github.com/darahayes/go-boom"
	"github.com/gorilla/mux"
)

// NewRouter returns the router created by the internal newRouter function
// Don't expose it directly in case we need to pass extra config that doesn't
// apply to the router.
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Use(loggerMiddleWare)
	router.Use(jsonContentTypeMiddleWare)
	router.Use(boom.RecoverHandler)

	router.NotFoundHandler = http.HandlerFunc(boom.NotFoundHandler)

	return router
}

func MetricsRoute(r *mux.Router, handler *metricsHandler) {
	r.HandleFunc("/metrics", handler.CreateMetric).Methods("POST")
}

func HealthzRoute(r *mux.Router, handler *healthHandler) {
	r.HandleFunc("/healthz", handler.Healthz)
	r.HandleFunc("/healthz/ping", handler.Ping)
}
