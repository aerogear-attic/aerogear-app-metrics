package dao

import (
	"database/sql"
)

type MetricsDAO struct {
	db *sql.DB
}

// Create a metrics record
func (m *MetricsDAO) Create(clientId string, metricsData []byte) error {
	_, err := m.db.Exec("INSERT INTO mobileappmetrics(clientId, data) VALUES($1, $2)", clientId, metricsData)
	return err
}

// IsHealthy checks that we are connected to the database
// This will be used by the healthcheck
func (m *MetricsDAO) IsHealthy() (bool, error) {
	err := m.db.Ping()
	return err == nil, err
}

func NewMetricsDAO(db *sql.DB) *MetricsDAO {
	return &MetricsDAO{
		db: db,
	}
}
