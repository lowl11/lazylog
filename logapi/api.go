package logapi

import (
	"github.com/lowl11/lazylog/loggers/console_logger"
	"github.com/lowl11/lazylog/loggers/file_logger"
	"os"
	"sync"
)

type Logger struct {
	loggers []ILogger
	mutex   sync.Mutex
}

func (logger *Logger) File(fileBase string, filePath ...string) *Logger {
	var singleFilePath string
	if len(filePath) > 0 {
		singleFilePath = filePath[0]
	}

	logger.loggers = append(logger.loggers, file_logger.Create(fileBase, singleFilePath))
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
		mutex: sync.Mutex{},
	}
}

func (logger *Logger) Debug(value string, args ...any) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	for _, logger := range logger.loggers {
		logger.Debug(value, args...)
	}
}

func (logger *Logger) Info(value string, args ...any) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	for _, logger := range logger.loggers {
		logger.Info(value, args...)
	}
}

func (logger *Logger) Warn(value string, args ...any) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	for _, logger := range logger.loggers {
		logger.Warn(value, args...)
	}
}

func (logger *Logger) Error(err error, value string, args ...any) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	for _, logger := range logger.loggers {
		logger.Error(err, value, args...)
	}
}

func (logger *Logger) Fatal(err error, value string, args ...any) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	for _, logger := range logger.loggers {
		logger.Fatal(err, value, args...)
	}
	os.Exit(1)
}
