package utils

import (
	"inventory-service/app/config"
	"inventory-service/pkg/logger"
)

func GetLogLevel() logger.LogLevel {
	config := config.GetConfig()
	logLevel, errLogLevel := logger.GetLogLevelFromString(config.LogLevel)
	if errLogLevel != nil {
		defaultLogLevel, errDefaultLogLevel := logger.GetLogLevelFromString(config.DefaultLogLevel)
		if errDefaultLogLevel != nil {
			panic("Unable to get LogLevel")
		}
		return defaultLogLevel
	}
	return logLevel
}
