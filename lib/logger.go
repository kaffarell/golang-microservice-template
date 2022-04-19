package lib

import (
	"go.uber.org/zap"
)

type Logger = *zap.Logger

var (
	globalLogger *Logger
)

func GetLogger() Logger {
	if globalLogger == nil {
		logger, err := zap.NewProduction()
		if err != nil {
			// Error creating logger
		}
		globalLogger = &logger
		return *globalLogger
	}
	return *globalLogger
}
