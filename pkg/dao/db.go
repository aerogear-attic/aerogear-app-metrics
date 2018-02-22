package dao

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

// var db *sql.DB

type DatabaseHandler struct {
	DB *sql.DB
}

func (handler *DatabaseHandler) Connect(dbHost, dbUser, dbPassword, dbName, sslMode string) error {
	if handler.DB != nil {
		return nil
	}
	//connection logic
	connStr := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=%v", dbHost, dbUser, dbPassword, dbName, sslMode)

	// sql.Open doesn't initialize the connection immediately
	dbInstance, err := sql.Open("postgres", connStr)

	// an error can happen here if the connection string is invalid
	if err != nil {
		return err
	}

	// an error happens here if we cannot connect
	if err = dbInstance.Ping(); err != nil {
		return err
	}

	// assign db variable declared above
	handler.DB = dbInstance
	return nil
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
	if _, err := handler.DB.Exec("CREATE TABLE IF NOT EXISTS mobileappmetrics(clientId varchar(30) NOT NULL CHECK (clientId <> ''), event_time timestamptz NOT NULL DEFAULT now() Not NULL, data jsonb)"); err != nil {
		return err
	}
	return nil
}
