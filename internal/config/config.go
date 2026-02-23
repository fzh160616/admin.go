package config

import "os"

type Config struct {
	AppPort   string
	MySQLDSN  string
	JWTSecret string
}

func Load() Config {
	return Config{
		AppPort:   getEnv("APP_PORT", "8080"),
		MySQLDSN:  getEnv("MYSQL_DSN", "root:root@tcp(127.0.0.1:3306)/admin?parseTime=true&charset=utf8mb4&loc=Local"),
		JWTSecret: getEnv("JWT_SECRET", "change-me-in-prod"),
	}
}

func getEnv(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}
