package console

import (
	"fmt"
	"sync"
)

var (
	consoleLog *ConsoleLog
	once       sync.Once
)

type ConsoleLog struct{}

func GetConsoleLog() *ConsoleLog {
	once.Do(func() {
		initConsoleLog()
	})
	return consoleLog
}

func initConsoleLog() {
	consoleLog = &ConsoleLog{}
}

func logMessage(message string, keyValues ...interface{}) {
	if len(keyValues)%2 != 0 {
		fmt.Println("Error: Key-value pairs must be provided in pairs.")
		return
	}

	logMessage := message + " "
	for i := 0; i < len(keyValues); i += 2 {
		key := fmt.Sprintf("%v", keyValues[i])
		value := fmt.Sprintf("%v", keyValues[i+1])
		logMessage += fmt.Sprintf("%s=%s ", key, value)
	}

	fmt.Println(logMessage)
}

func (c *ConsoleLog) Debug(message string, keyValues ...interface{}) {
	logMessage(message, keyValues...)
}

func (c *ConsoleLog) Info(message string, keyValues ...interface{}) {
	logMessage(message, keyValues...)
}

func (c *ConsoleLog) Warn(message string, keyValues ...interface{}) {
	logMessage(message, keyValues...)
}

func (c *ConsoleLog) Error(message string, keyValues ...interface{}) {
	logMessage(message, keyValues...)
}

func (c *ConsoleLog) Fatal(message string, keyValues ...interface{}) {
	logMessage(message, keyValues...)
}
