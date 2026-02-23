package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	AppPort                string
	MySQLDSN               string
	JWTSecret              string
	LoginMaxAttempts       int
	LoginAttemptWindowMin  int
	LoginBlockMinutes      int
}

func Load() Config {
	return Config{
		AppPort:               getEnv("APP_PORT", "8080"),
		MySQLDSN:              getEnv("MYSQL_DSN", "root:root@tcp(127.0.0.1:3306)/admin?parseTime=true&charset=utf8mb4&loc=Local"),
		JWTSecret:             getEnv("JWT_SECRET", "change-me-in-prod"),
		LoginMaxAttempts:      getEnvInt("LOGIN_MAX_ATTEMPTS", 5),
		LoginAttemptWindowMin: getEnvInt("LOGIN_ATTEMPT_WINDOW_MIN", 10),
		LoginBlockMinutes:     getEnvInt("LOGIN_BLOCK_MIN", 15),
	}
}

func LoginWindow(cfg Config) time.Duration {
	return time.Duration(cfg.LoginAttemptWindowMin) * time.Minute
}

func LoginBlock(cfg Config) time.Duration {
	return time.Duration(cfg.LoginBlockMinutes) * time.Minute
}

func getEnv(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}

func getEnvInt(key string, fallback int) int {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return fallback
	}
	return n
}
