package handlers

import (
	"net/http"
	"time"

	"github.com/aerogear/aerogear-metrics-api/pkg/models"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	status := models.Healthz{
		Timestamp: time.Now().UTC(),
		Status:    "ok",
	}

	respondWithJSON(w, 200, status)
}
