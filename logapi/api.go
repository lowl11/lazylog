package logapi

import (
	"github.com/lowl11/lazylog/loggers/console_logger"
	"github.com/lowl11/lazylog/loggers/file_logger"
	"os"
)

type Logger struct {
	loggers []ILogger
}

func (logger *Logger) File(fileBase string) *Logger {
	logger.loggers = append(logger.loggers, file_logger.Create(fileBase))
	return logger
}

func (logger *Logger) Custom(customLogger ILogger) *Logger {
	logger.loggers = append(logger.loggers, customLogger)
	return logger
}

func New() *Logger {
	return &Logger{
		loggers: []ILogger{
			console_logger.Create(),
		},
	}
}

func (logger *Logger) Debug(message string) {
	for _, logger := range logger.loggers {
		logger.Debug(message)
	}
}

func (logger *Logger) Info(message string) {
	for _, logger := range logger.loggers {
		logger.Info(message)
	}
}

func (logger *Logger) Warn(message string) {
	for _, logger := range logger.loggers {
		logger.Warn(message)
	}
}

func (logger *Logger) Error(err error, message string) {
	for _, logger := range logger.loggers {
		logger.Error(err, message)
	}
}

func (logger *Logger) Fatal(err error, message string) {
	for _, logger := range logger.loggers {
		logger.Fatal(err, message)
	}
	os.Exit(1)
}
