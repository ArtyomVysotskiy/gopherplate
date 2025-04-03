package logger

import (
	"go.uber.org/zap"
)

// Interface - определяет методы логгера.
type Interface interface {
	Debug(message string, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
	Fatal(message string, args ...interface{})
}

// Logger - структура логгера.
type Logger struct {
	logger *zap.SugaredLogger
}

var _ Interface = (*Logger)(nil)

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
	sugar := logger.Sugar()

	return &Logger{logger: sugar}
}

// Debug - логирует отладочное сообщение.
func (l *Logger) Debug(message string, args ...interface{}) {
	l.logger.Debugf(message, args...)
}

// Info - логирует информационное сообщение.
func (l *Logger) Info(message string, args ...interface{}) {
	l.logger.Infof(message, args...)
}

// Warn - логирует предупреждение.
func (l *Logger) Warn(message string, args ...interface{}) {
	l.logger.Warnf(message, args...)
}

// Error - логирует ошибку.
func (l *Logger) Error(message string, args ...interface{}) {
	l.logger.Errorf(message, args...)
}

// Fatal - логирует критическую ошибку и завершает выполнение.
func (l *Logger) Fatal(message string, args ...interface{}) {
	l.logger.Fatalf(message, args...)
}
