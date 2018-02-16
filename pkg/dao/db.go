package dao

import "database/sql"

var db *sql.DB

func Connect() (*sql.DB, error) {
	if db != nil {
		return db, nil
	}
	//connection logic
	return nil, nil
}

func Disconnect() error {
	if nil != db {
		return db.Close()
	}
	return nil
}
