package dao

import (
	"database/sql"
	"errors"

	"github.com/aerogear/aerogear-metrics-api/pkg/mobile"
)

type MetricsDAO struct {
	db *sql.DB
}

// Create a metrics record
func (m *MetricsDAO) Create(metric mobile.Metric) (mobile.Metric, error) {
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

// Ping checks that we are connected to the database
// This will be used by the healthcheck
func (m *MetricsDAO) Ping() error {
	return m.db.Ping()
}

func NewMetricsDAO(db *sql.DB) *MetricsDAO {
	return &MetricsDAO{
		db: db,
	}
}
