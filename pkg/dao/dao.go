package dao

import (
	"errors"

	"github.com/aerogear/aerogear-metrics-api/pkg/models"
)

type MetricsDAO struct {
	DBConnectionString string
}

// Connect to database
func (m *MetricsDAO) Connect() {

}

// Create a metrics record
func (m *MetricsDAO) Create(metric models.Metric) (models.Metric, error) {
	return metric, errors.New("Not Implemented yet")
}

// Update an existing job
// Not sure if we need this
func (m *MetricsDAO) Update() {

}

// Delete an existing job
// Not sure if we need this
func (m *MetricsDAO) Delete() {

}

// CheckConnection checks that we are connected to the database
// This will be used by the healthcheck
func (m *MetricsDAO) CheckConnection() {

}

func GetMetricsDAO(connectionString string) MetricsDAO {
	return MetricsDAO{
		DBConnectionString: connectionString,
	}
}
