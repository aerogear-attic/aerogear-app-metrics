// +build integration

package dao

import (
	"testing"

	"github.com/aerogear/aerogear-metrics-api/pkg/config"
)

func TestConnect(t *testing.T) {
	config := config.GetConfig()
	dbHandler := DatabaseHandler{}

	err := dbHandler.Connect(config["DBHost"], config["DBUser"], config["DBPassword"], config["DBName"], config["SSLMode"])

	if err != nil {
		t.Errorf("Connect() returned an error: %s", err.Error())
	}

	err = dbHandler.DB.Ping()

	if err != nil {
		t.Errorf("Failed to Ping the database after Connect(): %s", err.Error())
	}
}

func TestConnectAlreadyConnected(t *testing.T) {
	config := config.GetConfig()
	dbHandler := DatabaseHandler{}

	err := dbHandler.Connect(config["DBHost"], config["DBUser"], config["DBPassword"], config["DBName"], config["SSLMode"])

	if err != nil {
		t.Errorf("Connect() returned an error: %s", err.Error())
	}

	err = dbHandler.Connect(config["DBHost"], config["DBUser"], config["DBPassword"], config["DBName"], config["SSLMode"])

	if err != nil {
		t.Errorf("Connect() returned an error: %s", err.Error())
	}
}

func TestDisconnect(t *testing.T) {
	config := config.GetConfig()
	dbHandler := DatabaseHandler{}

	err := dbHandler.Connect(config["DBHost"], config["DBUser"], config["DBPassword"], config["DBName"], config["SSLMode"])

	if err != nil {
		t.Errorf("Connect() returned an error: %s", err.Error())
	}

	err = dbHandler.Disconnect()

	if err != nil {
		t.Errorf("Disconnect() returned an error: %s", err.Error())
	}

	err = dbHandler.DB.Ping()

	if err == nil {
		t.Errorf("Ping did not return an error after calling Disconnect()")
	}
}

func TestDisconnectNotConnected(t *testing.T) {
	dbHandler := DatabaseHandler{}

	err := dbHandler.Disconnect()

	if err != nil {
		t.Errorf("Disconnect() returned an error: %s", err.Error())
	}
}

func TestDoInitialSetup(t *testing.T) {
	config := config.GetConfig()
	dbHandler := DatabaseHandler{}

	err := dbHandler.Connect(config["DBHost"], config["DBUser"], config["DBPassword"], config["DBName"], config["SSLMode"])

	if err != nil {
		t.Errorf("Connect() returned an error: %s", err.Error())
	}

	err = dbHandler.DoInitialSetup()

	if err != nil {
		t.Errorf("DoInitialSetup() returned an error: %s", err.Error())
	}

	var exists bool

	err = dbHandler.DB.QueryRow("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'mobileappmetrics');").Scan(&exists)

	if err != nil {
		t.Errorf("Database returned an error while checking if table exists: %s", err.Error())
	}

	if !exists {
		t.Errorf("Expected table mobileappmetrics does not exist")
	}
}

func TestDoInitialSetupNotConnected(t *testing.T) {
	dbHandler := DatabaseHandler{}

	err := dbHandler.DoInitialSetup()

	if err == nil {
		t.Errorf("DoInitialSetup did not return an error")
	}
}
