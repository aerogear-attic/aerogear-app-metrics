package dao

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/lib/pq"
)

type DatabaseHandler struct {
	DB *sql.DB
}

func (handler *DatabaseHandler) Connect(connStr string, maxConnections int) error {
	if handler.DB != nil {
		return nil
	}

	// sql.Open doesn't initialize the connection immediately
	dbInstance, err := sql.Open("postgres", connStr)
	handler.DB = dbInstance

	dbInstance.SetMaxOpenConns(maxConnections)

	// an error can happen here if the connection string is invalid
	if err != nil {
		return err
	}

	// basic connection retry logic
	// mostly for issues where db server takes a few seconds to be ready
	for retry := 1; retry <= 5; retry++ {
		err = dbInstance.Ping()
		if err == nil {
			return nil
		}
		time.Sleep(1 * time.Second)
	}

	return err
}

func (handler *DatabaseHandler) Disconnect() error {
	if handler.DB != nil {
		return handler.DB.Close()
	}
	return nil
}

func (handler *DatabaseHandler) DoInitialSetup() error {
	if handler.DB == nil {
		return errors.New("cannot setup database, must call Connect() first")
	}
	if _, err := handler.DB.Exec(`CREATE UNLOGGED TABLE IF NOT EXISTS mobilemetrics_app(
		clientId char(80) NOT NULL CHECK (clientId <> ''),
		event_time timestamptz NOT NULL DEFAULT now(),
		client_time timestamptz DEFAULT now(),
		app_id char(40) NOT NULL,
		sdk_version char(20) NOT NULL,
		app_version char(20) NOT NULL,
		PRIMARY KEY (clientId, event_time)
	)`); err != nil {
		return err
	}

	if _, err := handler.DB.Exec(`CREATE UNLOGGED TABLE IF NOT EXISTS mobilemetrics_device(
		clientId char(80) NOT NULL CHECK (clientId <> ''),
		event_time timestamptz NOT NULL DEFAULT now(),
		client_time timestamptz DEFAULT now(),
		platform char(20) NOT NULL,
		platform_version char(20) NOT NULL,
		PRIMARY KEY (clientId, event_time)
	)`); err != nil {
		return err
	}

	if _, err := handler.DB.Exec(`CREATE UNLOGGED TABLE IF NOT EXISTS mobilemetrics_security(
		clientId char(80) NOT NULL CHECK (clientId <> ''),
		event_time timestamptz NOT NULL DEFAULT now(),
		client_time timestamptz DEFAULT now(),
		id char(80) NOT NULL,
		name char(40) NOT NULL,
		passed boolean,
		PRIMARY KEY (clientId, event_time)
	)`); err != nil {
		return err
	}

	return nil

}
