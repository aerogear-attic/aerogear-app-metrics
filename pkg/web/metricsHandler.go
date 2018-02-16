package web

import (
	"encoding/json"
	"net/http"


	"github.com/darahayes/go-boom"
	"github.com/aerogear/aerogear-metrics-api/pkg/mobile"
)

type metricsHandler struct{
	metricService *mobile.MetricsService
}

func NewMetricsHandler(ms *mobile.MetricsService)*metricsHandler  {
	return &metricsHandler{metricService:ms}
}

func(mh * metricsHandler) CreateMetric(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var metric mobile.Metric

	// decode the client payload into the metric var
	if err := json.NewDecoder(r.Body).Decode(&metric); err != nil {
		boom.BadRequest(w, "Invalid Data")
		return
	}

	// create the record in the db
	result, err := mh.metricService.Create(metric)

	// handle errors
	if err != nil {
		boom.BadImplementation(w)
		return
	}

	if err := withJSON(w, 200, result); err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
}
