package web

import (
	"net/http"

	"github.com/darahayes/go-boom"
)

type healthHandler struct {
	healthCheckTarget HealthCheckable
}

// NewHealthHandler creates a handler for the Health endpoints
// healthCheckTarget must contain an implementation that will be utilized to signal that
// the service is ready to accept requests.
func NewHealthHandler(healthCheckTarget HealthCheckable) *healthHandler {
	return &healthHandler{
		healthCheckTarget: healthCheckTarget,
	}
}

// Ping is the implementation for a liveness endpoint.
// It signals the process is up and running but doesn't guarantee connectivity
func (hh *healthHandler) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

// Healthz is the implementation for a readiness endpoint.
// It signals this process is ready to accept connections
// and respond to API requests
func (hh *healthHandler) Healthz(w http.ResponseWriter, r *http.Request) {
	if err := hh.healthCheckTarget.IsHealthy(); err != nil {
		boom.ServerUnavailable(w)
		return
	}

	w.WriteHeader(200)
}
