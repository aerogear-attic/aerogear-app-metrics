package dao

import (
	"database/sql"
	"time"
)

type MetricsDAO struct {
	db *sql.DB
}

// Create a metrics record
func (m *MetricsDAO) Create(clientId string, metricsData []byte, clientTime *time.Time) error {
	_, err := m.db.Exec("INSERT INTO mobileappmetrics(clientId, data, client_time) VALUES($1, $2, $3)", clientId, metricsData, clientTime)
	return err
}

// IsHealthy checks that we are connected to the database
// This will be used by the healthcheck
func (m *MetricsDAO) IsHealthy() error {
	// bug in m.db.Ping() doesn't return error if db goes down
	_, err := m.db.Exec("SELECT 1;")
	return err
}

func NewMetricsDAO(db *sql.DB) *MetricsDAO {
	return &MetricsDAO{
		db: db,
	}
}
