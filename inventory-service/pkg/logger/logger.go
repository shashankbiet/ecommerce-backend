package logger

import (
	"sync"

	"github.com/shashankbiet/go-common/logger"
)

var (
	once sync.Once
	Log  *LogWrapper
)

type LogWrapper struct{}

func InitLogger() {
	once.Do(func() {
		logger.InitDefaultLogger(logger.LogTypeZap, logger.LogLevelDebug)
	})
}

func (lw *LogWrapper) Debug(message string, keyValues ...interface{}) {
	logger.Debug(message, keyValues...)
}

func (lw *LogWrapper) Info(message string, keyValues ...interface{}) {
	logger.Info(message, keyValues...)
}

func (lw *LogWrapper) Warn(message string, keyValues ...interface{}) {
	logger.Warn(message, keyValues...)
}

func (lw *LogWrapper) Error(message string, keyValues ...interface{}) {
	logger.Error(message, keyValues...)
}
