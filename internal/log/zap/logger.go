package zaplogger

import (
	"github.com/Mindslave/skade/internal/engine"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type engineLogger struct {
	logger *zap.SugaredLogger
}

func NewEngineZapLogger() engine.Logger {
	logger := initLogger()
	return &engineLogger{
		logger,
	}

}

func initLogger() *zap.SugaredLogger {
	baselogger, err := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "msg",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalColorLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}.Build()
	if err != nil {
		panic("failed to initalize logger")
	}
	logger := baselogger.Sugar()
	return logger
}

func (el *engineLogger) Debug(message string) {
	el.logger.Debug(message)
}

func (el *engineLogger) Info(message string) {
	el.logger.Info(message)
}

func (el *engineLogger) Error(message string) {
	el.logger.Error(message)
}
