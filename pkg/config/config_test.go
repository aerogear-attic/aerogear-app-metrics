package config

import (
	"os"
	"reflect"
	"testing"
)

func TestGetConfig(t *testing.T) {
	expected := map[string]string{
		"DBHost":        "localhost",
		"DBUser":        "postgres",
		"DBPassword":    "postgres",
		"DBName":        "aerogear_mobile_metrics",
		"SSLMode":       "disable",
		"ListenAddress": ":3000",
	}

	config := GetConfig()

	if !reflect.DeepEqual(config, expected) {
		t.Error("GetConfig() did not return expected result")
	}
}

func TestGetConfigCustomEnvVariables(t *testing.T) {
	expected := map[string]string{
		"DBHost":        "testing",
		"DBUser":        "testing",
		"DBPassword":    "testing",
		"DBName":        "testing",
		"SSLMode":       "testing",
		"ListenAddress": ":testing",
	}

	os.Setenv("PGHOST", "testing")
	os.Setenv("PGUSER", "testing")
	os.Setenv("PGPASSWORD", "testing")
	os.Setenv("PGDATABASE", "testing")
	os.Setenv("PGSSLMODE", "testing")
	os.Setenv("PORT", "testing")

	config := GetConfig()

	if !reflect.DeepEqual(config, expected) {
		t.Error("GetConfig() did not return expected result")
	}
}

func TestGetConfigEmptyEnvVariables(t *testing.T) {
	expected := map[string]string{
		"DBHost":        "localhost",
		"DBUser":        "postgres",
		"DBPassword":    "postgres",
		"DBName":        "aerogear_mobile_metrics",
		"SSLMode":       "disable",
		"ListenAddress": ":3000",
	}

	os.Setenv("PGHOST", "")
	os.Setenv("PGUSER", "")
	os.Setenv("PGPASSWORD", "")
	os.Setenv("PGDATABASE", "")
	os.Setenv("PGSSLMODE", "")
	os.Setenv("PORT", "")

	config := GetConfig()

	if !reflect.DeepEqual(config, expected) {
		t.Error("GetConfig() did not return expected result")
	}
}
