package dao

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect(dbHost, dbUser, dbPassword, dbName, sslMode string) (*sql.DB, error) {
	if db != nil {
		return db, nil
	}
	//connection logic
	connStr := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=%v", dbHost, dbUser, dbPassword, dbName, sslMode)

	// sql.Open doesn't initialize the connection immediately
	dbInstance, err := sql.Open("postgres", connStr)

	// an error can happen here if the connection string is invalid
	if err != nil {
		return nil, err
	}

	// an error happens here if we cannot connect
	if err = dbInstance.Ping(); err != nil {
		return nil, err
	}

	// assign db variable declared above
	db = dbInstance
	return dbInstance, nil
}

func Disconnect() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

func DoInitialSetup() error {
	if db == nil {
		return errors.New("cannot setup database, must call Connect() first")
	}
	if _, err := db.Exec("CREATE EXTENSION IF NOT EXISTS timescaledb CASCADE;"); err != nil {
		return err
	}
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS mobileappmetrics(clientId varchar(30) NOT NULL, event_time timestamptz NOT NULL DEFAULT now() Not NULL, data jsonb)"); err != nil {
		return err
	}
	return nil
}
