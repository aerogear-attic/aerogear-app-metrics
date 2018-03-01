package config

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type config struct {
	DBConnectionString string
	DBMaxConnections   int
	ListenAddress      string
	LogLevel           string
	LogFormat          string
}

func GetConfig() config {
	return config{
		DBConnectionString: getDBConnectionString(),
		DBMaxConnections:   getEnvInt("DBMAX_CONNECTIONS", 100),
		ListenAddress:      fmt.Sprintf(":%v", getEnvInt("PORT", 3000)),
		LogLevel:           strings.ToLower(getEnv("LOG_LEVEL", "info")),
		LogFormat:          strings.ToLower(getEnv("LOG_FORMAT", "text")), //can be text or json
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}

// builds a libpq compatible connection string e.g. "user=postgresql host=localhost password=postgres"
func getDBConnectionString() string {

	// These are all of the options supported by pq
	dbConnectionConfig := map[string]string{
		"dbname":                    getEnv("PGDATABASE", "aerogear_mobile_metrics"),
		"user":                      getEnv("PGUSER", "postgresql"),
		"password":                  getEnv("PGPASSWORD", "postgres"),
		"host":                      getEnv("PGHOST", "localhost"),
		"port":                      getEnv("PGPORT", "5432"),
		"sslmode":                   getEnv("PGSSLMODE", "disable"),
		"connect_timeout":           getEnv("PGCONNECT_TIMEOUT", "5"),
		"fallback_application_name": getEnv("PGAPPNAME", ""),
		"sslcert":                   getEnv("PGSSLCERT", ""),
		"sslkey":                    getEnv("PGSSLKEY", ""),
		"sslrootcert":               getEnv("PGSSLROOTCERT", ""),
	}

	var options []string

	// build the list of options such as "user=postgresql"
	// omit any unset options
	for k, v := range dbConnectionConfig {
		if dbConnectionConfig[k] != "" {
			options = append(options, fmt.Sprintf("%v=%v", k, v))
		}
	}

	// sort the list (because ordering in maps is random)
	// this ensures connection string is the same every time. Makes testing much easier
	sort.Strings(options)

	return strings.Join(options, " ")
}
