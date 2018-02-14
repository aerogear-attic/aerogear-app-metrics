package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aerogear/aerogear-metrics-api/pkg/models"
	boom "github.com/darahayes/go-boom"
)

func createMetric(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var metric models.Metric

	// decode the client payload into the metric var
	if err := json.NewDecoder(r.Body).Decode(&metric); err != nil {
		boom.BadRequest(w, "Invalid Data")
		return
	}

	// create the record in the db
	result, err := app.Metrics.Create(metric)

	// handle errors
	if err != nil {
		boom.BadImplementation(w)
		return
	}

	respondWithJSON(w, 200, result)
}
