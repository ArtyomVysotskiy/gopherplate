package logger

import (
	"go.uber.org/zap"
)

// Interface - определяет методы логгера.
type Interface interface {
	Debug(message string, args ...zap.Field)
	Info(message string, args ...zap.Field)
	Warn(message string, args ...zap.Field)
	Error(message string, args ...zap.Field)
	Fatal(message string, args ...zap.Field)
}

// Logger - структура логгера.
type Logger struct {
	logger *zap.Logger
}

// New - создаёт новый экземпляр логгера.
func New(level string) *Logger {
	var cfg zap.Config

	switch level {
	case "debug":
		cfg = zap.NewDevelopmentConfig()
	case "info":
		cfg = zap.NewProductionConfig()
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		cfg = zap.NewProductionConfig()
		cfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		cfg = zap.NewProductionConfig()
		cfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		cfg = zap.NewProductionConfig()
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	logger, _ := cfg.Build()

	return &Logger{logger: logger}
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.logger.Fatal(msg, fields...)
}
