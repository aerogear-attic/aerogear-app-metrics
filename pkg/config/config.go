package config

import "os"

type config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	SSLMode    string
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}

func GetConfig() *config {
	return &config{
		DBHost:     getEnv("PGHOST", "localhost"),
		DBUser:     getEnv("PGPORT", "postgres"),
		DBPassword: getEnv("PGPASSWORD", "postgres"),
		DBName:     getEnv("PGDATABASE", "aerogear_mobile_metrics"),
		SSLMode:    getEnv("PGSSLMODE", "disable"),
	}
}
