package dao

import (
	"database/sql"
	"time"

	"github.com/aerogear/aerogear-app-metrics/pkg/mobile"
)

type MetricsDAO struct {
	db *sql.DB
}

func (m *MetricsDAO) RunInTransaction(cb func(*sql.Tx) error) (err error) {
	tx, err := m.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if e := recover(); e != nil {
			tx.Rollback()
			panic(e)
		}
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = cb(tx)
	return err
}

// Create a metrics record
func (m *MetricsDAO) Create(clientId string, metric mobile.Metric, clientTime *time.Time) error {
	eventTime := time.Now()
	return m.RunInTransaction(func(tx *sql.Tx) error {
		if metric.Data.App != nil {
			appMetric := metric.Data.App
			stmt, err := tx.Prepare(`INSERT INTO mobilemetrics_app (
				clientId,
				client_time,
				event_time,
				app_id,
				sdk_version,
				app_version
			) VALUES($1, $2, $3, $4, $5, $6)`)
			if err != nil {
				return err
			}

			if _, err := stmt.Exec(
				clientId,
				clientTime,
				eventTime,
				appMetric.ID,
				appMetric.SDKVersion,
				appMetric.AppVersion,
			); err != nil {
				return err
			}
		}
		if metric.Data.Device != nil {
			deviceMetric := metric.Data.Device
			stmt, err := tx.Prepare(`INSERT INTO mobilemetrics_device (
				clientId,
				client_time,
				event_time,
				platform,
				platform_version
			) VALUES($1, $2, $3, $4, $5)`)
			if err != nil {
				return err
			}
			if _, err := stmt.Exec(
				clientId,
				clientTime,
				eventTime,
				deviceMetric.Platform,
				deviceMetric.PlatformVersion,
			); err != nil {
				return err
			}
		}
		if metric.Data.Security != nil {
			stmt, err := tx.Prepare(`INSERT INTO mobilemetrics_security (
				clientId,
				client_time,
				event_time,
				id,
				name,
				passed
			) VALUES($1, $2, $3, $4, $5, $6)`)
			if err != nil {
				return err
			}
			for _, securityMetric := range *metric.Data.Security {
				if _, err := stmt.Exec(
					clientId,
					clientTime,
					eventTime,
					securityMetric.Id,
					securityMetric.Name,
					securityMetric.Passed,
				); err != nil {
					return err
				}
			}
		}
		return nil
	})
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
