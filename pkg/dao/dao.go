package dao

import (
	"database/sql"
	"encoding/json"

	"github.com/aerogear/aerogear-metrics-api/pkg/mobile"
)

type MetricsDAO struct {
	db *sql.DB
}

// Create a metrics record
func (m *MetricsDAO) Create(metric mobile.Metric) error {
	data, err := json.Marshal(metric.Data)

	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO mobileappmetrics(clientId, data) VALUES($1, $2)", metric.ClientId, data)
	return err
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
