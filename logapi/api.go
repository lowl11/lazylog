package logapi

import (
	"github.com/lowl11/lazylog/loggers/console_logger"
	"github.com/lowl11/lazylog/loggers/file_logger"
	"os"
	"sync"
)

type Logger struct {
	loggers  []ILogger
	jsonMode bool
	mutex    sync.Mutex
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

func (logger *Logger) JSON() *Logger {
	logger.jsonMode = true
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

func (logger *Logger) Debug(args ...any) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	for _, logger := range logger.loggers {
		logger.Debug(args...)
	}
}

func (logger *Logger) Info(args ...any) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	for _, logger := range logger.loggers {
		logger.Info(args...)
	}
}

func (logger *Logger) Warn(args ...any) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	for _, logger := range logger.loggers {
		logger.Warn(args...)
	}
}

func (logger *Logger) Error(err error, args ...any) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	for _, logger := range logger.loggers {
		logger.Error(err, args...)
	}
}

func (logger *Logger) Fatal(err error, args ...any) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	for _, logger := range logger.loggers {
		logger.Fatal(err, args...)
	}
	os.Exit(1)
}
