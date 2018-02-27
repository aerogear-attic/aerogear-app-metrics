// +build integration
package dao

import (
	"testing"

	"github.com/aerogear/aerogear-app-metrics/pkg/config"
	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

func TestIsHealthy(t *testing.T) {
	config := config.GetConfig()
	dbHandler := DatabaseHandler{}

	err := dbHandler.Connect(config.DBConnectionString, config.DBMaxConnections)

	if err != nil {
		t.Errorf("Connect() returned an error: %s", err.Error())
	}

	dao := NewMetricsDAO(dbHandler.DB)

	isHealthy, err := dao.IsHealthy()

	if err != nil {
		t.Errorf("isHealthy returned an error %s", err.Error())
	}

	if !isHealthy {
		t.Errorf("isHealthy returned false")
	}
}

func TestIsHealthyWhenDisconnected(t *testing.T) {
	config := config.GetConfig()
	dbHandler := DatabaseHandler{}

	err := dbHandler.Connect(config.DBConnectionString, config.DBMaxConnections)

	if err != nil {
		t.Errorf("Connect() returned an error: %s", err.Error())
	}

	dao := NewMetricsDAO(dbHandler.DB)

	dbHandler.Disconnect()

	isHealthy, err := dao.IsHealthy()

	if err == nil {
		t.Errorf("isHealthy returned no error when disconnected")
	}

	if isHealthy {
		t.Errorf("isHealthy returned true when disconnected")
	}
}

func TestCreate(t *testing.T) {
	config := config.GetConfig()
	dbHandler := DatabaseHandler{}

	err := dbHandler.Connect(config.DBConnectionString, config.DBMaxConnections)

	if err != nil {
		t.Errorf("Connect() returned an error: %s", err.Error())
	}

	dbHandler.DoInitialSetup()

	if err != nil {
		t.Errorf("Connect() returned an error: %s", err.Error())
	}

	dao := NewMetricsDAO(dbHandler.DB)

	clientId := "org.aerogear.metrics.testing"
	metricsData := []byte("{\"app\":{\"id\":\"com.example.someApp\",\"sdkVersion\":\"2.4.6\",\"appVersion\":\"256\"},\"device\":{\"platform\":\"android\",\"platformVersion\":\"27\"}}")

	err = dao.Create(clientId, metricsData)

	if err != nil {
		t.Errorf("Create() returned an error %s", err.Error())
	}
}

func TestCreateBadJSON(t *testing.T) {
	config := config.GetConfig()
	dbHandler := DatabaseHandler{}

	err := dbHandler.Connect(config.DBConnectionString, config.DBMaxConnections)

	if err != nil {
		t.Errorf("Connect() returned an error: %s", err.Error())
	}

	dao := NewMetricsDAO(dbHandler.DB)

	clientId := "org.aerogear.metrics.testing"
	metricsData := []byte("InvalidJSON")

	err = dao.Create(clientId, metricsData)

	if err == nil {
		t.Errorf("Create() with invalid JSON did not return an error")
	}
}

func TestCreateEmptyClientID(t *testing.T) {
	config := config.GetConfig()
	dbHandler := DatabaseHandler{}

	err := dbHandler.Connect(config.DBConnectionString, config.DBMaxConnections)

	if err != nil {
		t.Errorf("Connect() returned an error: %s", err.Error())
	}

	dbHandler.DoInitialSetup()

	if err != nil {
		t.Errorf("Connect() returned an error: %s", err.Error())
	}

	dao := NewMetricsDAO(dbHandler.DB)

	clientId := ""
	metricsData := []byte("{\"app\":{\"id\":\"com.example.someApp\",\"sdkVersion\":\"2.4.6\",\"appVersion\":\"256\"},\"device\":{\"platform\":\"android\",\"platformVersion\":\"27\"}}")

	err = dao.Create(clientId, metricsData)

	if err == nil {
		t.Errorf("Create() with empty clientId did not return an error")
	}
}
