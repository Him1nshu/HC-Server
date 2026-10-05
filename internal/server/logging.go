package server

import (
	"log/slog"
	"os"
)

func NewLogger(level string) *slog.Logger {
	var configuredLevel slog.Level
	switch level {
	case "DEBUG":
		configuredLevel = slog.LevelDebug
	case "WARN":
		configuredLevel = slog.LevelWarn
	case "ERROR":
		configuredLevel = slog.LevelError
	default:
		configuredLevel = slog.LevelInfo
	}
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: configuredLevel}))
}
