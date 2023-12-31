package logger

import "fmt"

type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
)

var logLevelToString = map[LogLevel]string{
	LogLevelDebug: "DEBUG",
	LogLevelInfo:  "INFO",
	LogLevelWarn:  "WARN",
	LogLevelError: "ERROR",
	LogLevelFatal: "FATAL",
}

var stringToLogLevel = map[string]LogLevel{
	"DEBUG": LogLevelDebug,
	"INFO":  LogLevelInfo,
	"WARN":  LogLevelWarn,
	"ERROR": LogLevelError,
	"FATAL": LogLevelFatal,
}

// GetLogLevelString returns the string representation of the LogLevel.
func (l LogLevel) GetLogLevelString() string {
	return logLevelToString[l]
}

// GetLogLevelFromString returns the LogLevel based on the given string representation.
func GetLogLevelFromString(s string) (LogLevel, error) {
	level, ok := stringToLogLevel[s]
	if !ok {
		return LogLevelDebug, fmt.Errorf("unknown log level: %s", s)
	}
	return level, nil
}
