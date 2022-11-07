package file_logger

import (
	"log"
	"sync"
)

type Logger struct {
	writer   *log.Logger
	fileName string
	fileBase string

	mutex sync.Mutex
}

func Create(fileBase string) *Logger {
	logger := &Logger{
		fileBase: fileBase,

		mutex: sync.Mutex{},
	}
	logger.writer = log.New(logger.createFile(), "", 0)

	return logger
}
