package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type DatabaseHandler struct {
	DB *sql.DB
}

func (handler *DatabaseHandler) Connect(dbHost, dbUser, dbPassword, dbName, sslMode string) error {
	if handler.DB != nil {
		return nil
	}

	connStr := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=%v", dbHost, dbUser, dbPassword, dbName, sslMode)

	// sql.Open doesn't initialize the connection immediately
	dbInstance, err := sql.Open("postgres", connStr)
	handler.DB = dbInstance

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
	if _, err := handler.DB.Exec("CREATE TABLE IF NOT EXISTS mobileappmetrics(clientId varchar(128) NOT NULL CHECK (clientId <> ''), event_time timestamptz NOT NULL DEFAULT now() Not NULL, data jsonb)"); err != nil {
		return err
	}
	return nil
}
