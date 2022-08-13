package console_logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	fileNamePattern = "info_%s.log"
)

func createFile() *os.File {
	fileName := fmt.Sprintf(fileNamePattern, time.Now().Format("02-01-2006"))

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil
	}

	_ = file.Close()

	return file
}

func (logger *Logger) updateFile() {
	newFile := createFile()
	if newFile != nil {
		logger.writer = log.New(newFile, "", log.Ltime)
	}
}
