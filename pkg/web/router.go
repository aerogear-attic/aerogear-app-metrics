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
	router.Use(boom.RecoverHandler)

	router.NotFoundHandler = http.HandlerFunc(boom.NotFoundHandler)

	return router
}

func MetricsRoute(r *mux.Router, handler *metricsHandler) {
	// swagger:operation POST /metrics metrics
	//
	// Creates a metric
	//
	// ---
	//
	// consumes:
	// - application/json
	// produces:
	// - application/json
	// parameters:
	// - in: body
	//   name: body
	//   description: "Metrics object to create"
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Metric"
	// responses:
	//   '204':
	//     description: Metric created
	//   '400':
	//     description: Bad Request
	//   '500':
	//     description: Server Error
	r.HandleFunc("/metrics", handler.CreateMetric).Methods("POST")
}

func HealthzRoute(r *mux.Router, handler *healthHandler) {
	// swagger:operation GET /healthz healthz
	//
	// ---
	// produces:
	// - text/html
	// responses:
	//   '200':
	//     description: Health OK
	r.HandleFunc("/healthz", handler.Healthz)
	// swagger:operation GET /ping ping
	//
	// ---
	// produces:
	// - text/html
	// responses:
	//   '200':
	//     description: Ping OK
	r.HandleFunc("/ping", handler.Ping)
}
