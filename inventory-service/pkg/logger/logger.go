package logger

import "sync"

var (
	Log  *LogWrapper
	once sync.Once
)

type LogWrapper struct {
	logger   ILogger
	logLevel LogLevel
}

func InitLogger(logger ILogger, logLevel LogLevel) {
	once.Do(func() {
		Log = &LogWrapper{
			logger:   logger,
			logLevel: logLevel,
		}
	})
}

func (lw *LogWrapper) Debug(message string, keyValues ...interface{}) {
	if lw.logLevel <= LogLevelDebug {
		lw.logger.Debug(message, keyValues...)
	}
}

func (lw *LogWrapper) Info(message string, keyValues ...interface{}) {
	if lw.logLevel <= LogLevelInfo {
		lw.logger.Info(message, keyValues...)
	}
}

func (lw *LogWrapper) Warn(message string, keyValues ...interface{}) {
	if lw.logLevel <= LogLevelWarn {
		lw.logger.Warn(message, keyValues...)
	}
}

func (lw *LogWrapper) Error(message string, keyValues ...interface{}) {
	if lw.logLevel <= LogLevelError {
		lw.logger.Error(message, keyValues...)
	}
}

func (lw *LogWrapper) Fatal(message string, keyValues ...interface{}) {
	if lw.logLevel <= LogLevelDebug {
		lw.logger.Fatal(message, keyValues...)
	}
}
