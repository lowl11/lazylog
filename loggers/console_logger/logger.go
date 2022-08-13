package console_logger

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	writer *log.Logger
}

func Create(withFile bool) *Logger {
	writers := []io.Writer{os.Stdout}
	if withFile {
		writers = append(writers, createFile())
	}
	multipleWriter := io.MultiWriter(writers...)
	writer := log.New(multipleWriter, "", 0)

	return &Logger{
		writer: writer,
	}
}
