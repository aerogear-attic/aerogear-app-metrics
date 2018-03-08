package web

import (
	"encoding/json"
	"net/http"

	"github.com/aerogear/aerogear-app-metrics/pkg/mobile"
	"github.com/darahayes/go-boom"
	log "github.com/sirupsen/logrus"
)

type metricsHandler struct {
	metricService MetricsServiceInterface
}

func NewMetricsHandler(ms MetricsServiceInterface) *metricsHandler {
	return &metricsHandler{metricService: ms}
}

func (mh *metricsHandler) CreateMetric(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var metric mobile.Metric

	// decode the client payload into the metric var
	if err := json.NewDecoder(r.Body).Decode(&metric); err != nil {
		log.WithFields(log.Fields{"error": err.Error()}).Error("error parsing client payload")
		boom.BadRequest(w, "invalid JSON payload")
		return
	}

	if valid, reason := metric.Validate(); !valid {
		log.WithFields(log.Fields{"reason": reason}).Info("invalid client payload")
		boom.BadRequest(w, reason)
		return
	}

	// create the record in the db
	_, err := mh.metricService.Create(metric)

	// handle errors
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error()}).Error("error creating metric")
		boom.BadImplementation(w)
		return
	}

	w.WriteHeader(204)
}
