package config

import (
	"os"
	"reflect"
	"testing"
)

func TestConfig(t *testing.T) {

	cases := []struct {
		Name     string
		Expected Config
		EnvVars  map[string]string
	}{
		{
			Name: "GetConfig() should return sensible defaults when no environemt variables are set",
			Expected: Config{
				ListenAddress:      ":3000",
				DBMaxConnections:   100,
				DBConnectionString: "connect_timeout=5 dbname=aerogear_mobile_metrics host=localhost password=postgres port=5432 sslmode=disable user=postgresql",
				LogFormat:          "text",
				LogLevel:           "info",
			},
			EnvVars: map[string]string{},
		},
		{
			Name: "GetConfig() should return correct config when environment variables are set",
			Expected: Config{
				ListenAddress:      ":3000",
				DBMaxConnections:   100,
				DBConnectionString: "connect_timeout=5 dbname=testing host=testing password=testing port=5432 sslmode=testing user=testing",
				LogFormat:          "testing",
				LogLevel:           "testing",
			},
			EnvVars: map[string]string{
				"PGHOST":     "testing",
				"PGUSER":     "testing",
				"PGPASSWORD": "testing",
				"PGDATABASE": "testing",
				"PGSSLMODE":  "testing",
				"LOG_LEVEL":  "testing",
				"LOG_FORMAT": "testing",
			},
		},
		{
			Name: "GetConfig() should return correct config when empty environment variables are set",
			Expected: Config{
				ListenAddress:      ":3000",
				DBMaxConnections:   100,
				DBConnectionString: "connect_timeout=5 dbname=aerogear_mobile_metrics host=localhost password=postgres port=5432 sslmode=disable user=postgresql",
				LogFormat:          "text",
				LogLevel:           "info",
			},
			EnvVars: map[string]string{
				"PGHOST":     "",
				"PGUSER":     "",
				"PGPASSWORD": "",
				"PGDATABASE": "",
				"PGSSLMODE":  "",
				"PORT":       "",
				"LOG_LEVEL":  "",
				"LOG_FORMAT": "",
			},
		},
		{
			Name: "GetConfig() parse appropriate integer environment variables",
			Expected: Config{
				ListenAddress:      ":4000",
				DBMaxConnections:   5,
				DBConnectionString: "connect_timeout=5 dbname=aerogear_mobile_metrics host=localhost password=postgres port=5432 sslmode=disable user=postgresql",
				LogFormat:          "text",
				LogLevel:           "info",
			},
			EnvVars: map[string]string{
				"DBMAX_CONNECTIONS": "5",
				"PORT":              "4000",
			},
		},
		{
			Name: "GetConfig() should return default values when non-integer environment variables are set",
			Expected: Config{
				ListenAddress:      ":3000",
				DBMaxConnections:   100,
				DBConnectionString: "connect_timeout=5 dbname=aerogear_mobile_metrics host=localhost password=postgres port=5432 sslmode=disable user=postgresql",
				LogFormat:          "text",
				LogLevel:           "info",
			},
			EnvVars: map[string]string{
				"DBMAX_CONNECTIONS": "not an integer",
				"PORT":              "not an integer",
			},
		},
	}

	for _, c := range cases {
		if len(c.EnvVars) != 0 {
			for name, value := range c.EnvVars {
				os.Setenv(name, value)
			}
		}

		config := GetConfig()

		if !reflect.DeepEqual(config, c.Expected) {
			t.Errorf("Failure in test case: %v \nexpected: %v\ngot: %v", c.Name, c.Expected, config)
		}
	}
}
