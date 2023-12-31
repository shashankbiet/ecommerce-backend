package zap

import (
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	zapLog *ZapLog
	once   sync.Once
)

type ZapLog struct {
	sugaredLogger *zap.SugaredLogger
}

func GetZapLog() *ZapLog {
	once.Do(func() {
		initZapLog()
	})
	return zapLog
}

func initZapLog() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	logger, err := config.Build()
	if err != nil {
		fmt.Printf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()
	zapLog = &ZapLog{
		sugaredLogger: sugar,
	}
}

func (z *ZapLog) Debug(message string, keyValues ...interface{}) {
	z.sugaredLogger.Debug(message, keyValues)
}

func (z *ZapLog) Info(message string, keyValues ...interface{}) {
	z.sugaredLogger.Info(message, keyValues)
}

func (z *ZapLog) Warn(message string, keyValues ...interface{}) {
	z.sugaredLogger.Warn(message, keyValues)
}

func (z *ZapLog) Error(message string, keyValues ...interface{}) {
	z.sugaredLogger.Error(message, keyValues)
}

func (z *ZapLog) Fatal(message string, keyValues ...interface{}) {
	z.sugaredLogger.Fatal(message, keyValues)
}
