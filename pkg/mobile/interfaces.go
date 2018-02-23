package mobile

// MetricCreator defines how a metric can be created
type MetricCreator interface {
	Create(clientId string, metricsData []byte) error
}
