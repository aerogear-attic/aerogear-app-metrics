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

type HealthHandler struct {
	pingTarget Pingable
}

// NewHealthHandler creates a handler for the Health endpoints
// pingTarget must contain an implementation that will be utilized to signal that
// the service is ready to accept requests.
func NewHealthHandler(pingTarget Pingable) *HealthHandler {
	return &HealthHandler{
		pingTarget: pingTarget,
	}
}

func (hh *HealthHandler) Healthz(w http.ResponseWriter, r *http.Request) {
	status := healthResponse{
		Timestamp: time.Now().UTC(),
		Status:    "ok",
	}

	if err := withJSON(w, 200, status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (hh *HealthHandler) Ping(w http.ResponseWriter, r *http.Request) {
	err := hh.pingTarget.Ping()
	if err != nil {
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
