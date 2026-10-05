package server

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Config struct {
	Address         string
	ShutdownTimeout time.Duration
	LogLevel        string
}

func LoadConfig() (Config, error) {
	cfg := Config{
		Address:         envOrDefault("HC_SERVER_ADDR", ":8080"),
		LogLevel:        strings.ToUpper(envOrDefault("HC_SERVER_LOG_LEVEL", "INFO")),
	}

	timeoutText := envOrDefault("HC_SERVER_SHUTDOWN_TIMEOUT", "10s")
	timeout, err := time.ParseDuration(timeoutText)
	if err != nil || timeout <= 0 {
		return Config{}, fmt.Errorf("HC_SERVER_SHUTDOWN_TIMEOUT must be a positive duration: %q", timeoutText)
	}
	cfg.ShutdownTimeout = timeout

	switch cfg.LogLevel {
	case "DEBUG", "INFO", "WARN", "ERROR":
	default:
		return Config{}, fmt.Errorf("HC_SERVER_LOG_LEVEL must be DEBUG, INFO, WARN, or ERROR: %q", cfg.LogLevel)
	}

	return cfg, nil
}

func envOrDefault(key, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(key)); value != "" {
		return value
	}
	return fallback
}
