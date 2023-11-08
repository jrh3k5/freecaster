package logging

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jrh3k5/freecaster/environment"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	zapLoggerOnce    sync.Once
	zapLoggerInitErr error
	zapLogger        *zap.Logger
)

// GetZapLogger gets a zap logger.
func GetZapLogger(_ context.Context) (*zap.Logger, error) {
	zapLoggerOnce.Do(func() {
		var zapConfig zap.Config
		if environment.GetEnviroment() == "production" {
			zapConfig = zap.NewProductionConfig()
		} else {
			zapConfig = zap.NewDevelopmentConfig()
		}

		logLevel := zap.WarnLevel
		if logLevelStr := os.Getenv("LOG_LEVEL"); logLevelStr != "" {
			parsedLogLevel, logLevelParseErr := zapcore.ParseLevel(logLevelStr)
			if logLevelParseErr != nil {
				log.Print("Failed to parse log level; default will be used")
			} else {
				logLevel = parsedLogLevel
			}
		}

		zapConfig.Level = zap.NewAtomicLevelAt(logLevel)

		zapLogger, zapLoggerInitErr = zapConfig.Build()
	})

	if zapLoggerInitErr != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", zapLoggerInitErr)
	}

	return zapLogger, nil
}
