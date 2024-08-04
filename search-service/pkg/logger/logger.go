package logger

import (
	"sync"

	"github.com/shashankbiet/go-common/logger"
)

var (
	once sync.Once
)

func InitLogger() {
	once.Do(func() {
		logger.InitDefaultLogger(logger.LogTypeZap, logger.LogLevelDebug)
	})
}

func Debug(message string, keyValues ...interface{}) {
	logger.Debug(message, keyValues...)
}

func Info(message string, keyValues ...interface{}) {
	logger.Info(message, keyValues...)
}

func Warn(message string, keyValues ...interface{}) {
	logger.Warn(message, keyValues...)
}

func Error(message string, keyValues ...interface{}) {
	logger.Error(message, keyValues...)
}
