package logging

import (
	"context"
	"log/slog"
	"os"
	"strconv"
	"sync"
)

var (
	loggerOnce sync.Once
	logger     *slog.Logger
)

// GetSlogLogger gets a slog logger.
func GetSlogLogger(ctx context.Context) *slog.Logger {
	loggerOnce.Do(func() {
		logLevel := slog.LevelWarn
		logLevelString := os.Getenv("LOG_LEVEL")
		if parsedLogLevel, parseErr := strconv.ParseInt(logLevelString, 10, 64); parseErr != nil {
			logLevel = slog.Level(parsedLogLevel)
		}

		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: logLevel,
		}))
	})

	return logger
}
