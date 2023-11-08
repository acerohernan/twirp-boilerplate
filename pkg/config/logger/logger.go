package logger

import (
	"github.com/acerohernan/twirp-boilerplate/pkg/config"
	"go.uber.org/zap"
)

var defaultLogger *Logger

type Logger struct {
	zap *zap.SugaredLogger
}

func InitLogger(config *config.LoggerConfig) error {
	zapLevel, err := zap.ParseAtomicLevel(config.Level)

	if err != nil {
		return err
	}

	zapConfig := zap.Config{
		Level:            zapLevel,
		Development:      false,
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := zapConfig.Build()

	if err != nil {
		return err
	}

	defaultLogger = &Logger{
		zap: logger.Sugar().WithOptions(zap.AddCallerSkip(1)),
	}

	return nil
}

func Infow(msg string, keysAndValues ...interface{}) {
	if defaultLogger == nil {
		return
	}

	defaultLogger.zap.Infow(msg, keysAndValues...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	if defaultLogger == nil {
		return
	}

	defaultLogger.zap.Debugw(msg, keysAndValues...)
}

func Warnw(msg string, err error, keysAndValues ...interface{}) {
	if defaultLogger == nil {
		return
	}

	defaultLogger.zap.Warnw(msg, keysAndValues...)
}

func Errorw(msg string, err error, keysAndValues ...interface{}) {
	if defaultLogger == nil {
		return
	}

	defaultLogger.zap.Errorw(msg, keysAndValues...)
}
