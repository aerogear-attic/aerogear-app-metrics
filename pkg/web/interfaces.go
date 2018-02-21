package web

import "github.com/aerogear/aerogear-metrics-api/pkg/mobile"

type MetricsServiceInterface interface {
	Create(m mobile.Metric) (mobile.Metric, error)
}