package logapi

import (
	"lazylog/loggers/console_logger"
	"os"
)

type Logger struct {
	loggers []ILogger
}

func New(loggerTypes ...string) ILogger {
	loggers := make([]ILogger, 0, 2)

	var isConsole, isFile bool
	for _, loggerType := range loggerTypes {
		switch loggerType {
		case Console:
			isConsole = true
			break
		case File:
			isFile = true
			break
		}
	}

	if isConsole {
		loggers = append(loggers, console_logger.Create(isFile))
	}

	if len(loggerTypes) == 0 && len(loggers) == 0 {
		loggers = append(loggers, console_logger.Create(isFile))
	}
	return &Logger{
		loggers: loggers,
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
