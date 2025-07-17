package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(fields ...zap.Field) *zap.Logger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339) // or time.RubyDate or "2006-01-02 15:04:05" or even freaking time.Kitchen
	config.DisableStacktrace = true
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	return logger.With(fields...)
}
