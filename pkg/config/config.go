package config

import (
	"fmt"
	"os"
)

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}

func GetConfig() map[string]string {
	return map[string]string{
		"DBHost":        getEnv("PGHOST", "localhost"),
		"DBUser":        getEnv("PGUSER", "postgresql"),
		"DBPassword":    getEnv("PGPASSWORD", "postgres"),
		"DBName":        getEnv("PGDATABASE", "aerogear_mobile_metrics"),
		"SSLMode":       getEnv("PGSSLMODE", "disable"),
		"ListenAddress": fmt.Sprintf(":%s", getEnv("PORT", "3000")),
	}
}
