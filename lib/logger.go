package lib

import (
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
}

var (
	globalLogger *Logger
)

type LoggerConfig struct {
	logLevel string
}

func GetLogger(loggerConfig LoggerConfig) Logger {
	if globalLogger == nil {
		logger, err := zap.NewProduction()
		if err != nil {
			// Error creating logger
		}
		globalLogger = &Logger{
			logger: logger,
		}
		return *globalLogger
	}
	return *globalLogger
}
