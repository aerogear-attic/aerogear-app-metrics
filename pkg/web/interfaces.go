package web

import "github.com/aerogear/aerogear-app-metrics/pkg/mobile"

type MetricsServiceInterface interface {
	Create(m mobile.Metric) (mobile.Metric, error)
}

type HealthCheckable interface {
	IsHealthy() error
}
