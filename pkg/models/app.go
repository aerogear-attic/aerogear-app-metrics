package models

type MetricsInterface interface {
	Create(Metric) (Metric, error)
}

type HealthInterface interface {
	IsHealthy() bool
}

type App struct {
	Metrics MetricsInterface
	Health  HealthInterface
}
