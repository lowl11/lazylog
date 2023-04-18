package logapi

import (
	"github.com/lowl11/lazylog/loggers/console_logger"
	"github.com/lowl11/lazylog/loggers/file_logger"
	"github.com/lowl11/lazylog/message_tools"
	"os"
	"sync"
	"time"
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

func (logger *Logger) JSON() *Logger {
	message_tools.JsonMode = true
	return logger
}

func (logger *Logger) NoTime() *Logger {
	message_tools.NoTimeMode = true
	return logger
}

func (logger *Logger) NoPrefix() *Logger {
	message_tools.NoPrefixMode = true
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

	for _, loggerItem := range logger.loggers {
		go loggerItem.Debug(args...)
	}
}

func (logger *Logger) Info(args ...any) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	for _, loggerItem := range logger.loggers {
		go loggerItem.Info(args...)
	}
}

func (logger *Logger) Warn(args ...any) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	for _, loggerItem := range logger.loggers {
		go loggerItem.Warn(args...)
	}
}

func (logger *Logger) Error(err error, args ...any) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	for _, loggerItem := range logger.loggers {
		go loggerItem.Error(err, args...)
	}
}

func (logger *Logger) Fatal(err error, args ...any) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	for _, loggerItem := range logger.loggers {
		go loggerItem.Fatal(err, args...)
	}

	time.Sleep(time.Second)
	os.Exit(1)
}
