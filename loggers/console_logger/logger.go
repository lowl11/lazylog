package console_logger

import (
	"log"
	"os"
	"sync"
)

type Logger struct {
	writer *log.Logger
	mutex  sync.Mutex
}

func Create() *Logger {
	writer := log.New(os.Stdout, "", 0)

	return &Logger{
		writer: writer,
		mutex:  sync.Mutex{},
	}
}
