package web

import (
	"net/http"
	"time"
)

// not reusable type so make it package local

type healthResponse struct {
	Timestamp time.Time `json:"time"`
	Status    string    `json:"status"`
}

type healthHandler struct {
	pingTarget Pingable
}

func NewHealthHandler(pingTarget Pingable) *healthHandler {
	return &healthHandler{
		pingTarget: pingTarget,
	}
}

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

func (hh *healthHandler) Ping(w http.ResponseWriter, r *http.Request) {
	err := hh.pingTarget.Ping()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	status := healthResponse{
		Timestamp: time.Now().UTC(),
		Status:    "ok",
	}
	if err := withJSON(w, 200, status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
