package logging

import (
	"context"
	"log/slog"
	"sync"
)

var (
	loggerOnce sync.Once
	logger     *slog.Logger
)

// GetSlogLogger gets a slog logger.
func GetSlogLogger(ctx context.Context) *slog.Logger {
	loggerOnce.Do(func() {
		logger = slog.Default()
	})

	return logger
}
