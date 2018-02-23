package web

import (
	"net/http"
	"time"

	"github.com/darahayes/go-boom"
)

// not reusable type so make it package local

type healthResponse struct {
	Timestamp time.Time `json:"time"`
	Status    string    `json:"status"`
}

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

// Healthz is the implementation for a liveness endpoint.
// It signals the process is up and running but doesn't guarantee connectivity
func (hh *healthHandler) Healthz(w http.ResponseWriter, r *http.Request) {
	status := healthResponse{
		Timestamp: time.Now().UTC(),
		Status:    "ok",
	}

	if err := withJSON(w, 200, status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Ping is the implementation for a readiness endpoint.
// It signals this process is ready to accept connections
// and respond to API requests
func (hh *healthHandler) Ping(w http.ResponseWriter, r *http.Request) {
	healthy, _ := hh.healthCheckTarget.IsHealthy()
	if !healthy {
		boom.ServerUnavailable(w)
		return
	}

	status := healthResponse{
		Timestamp: time.Now().UTC(),
		Status:    "ok",
	}
	if err := withJSON(w, 200, status); err != nil {
		boom.Internal(w)
		return
	}
}
