package file_logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	fileNamePattern = "%s_%s.log"
)

func (logger *Logger) createFile() *os.File {
	fileName := fmt.Sprintf(fileNamePattern, logger.fileBase, time.Now().Format("02-01-2006"))

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil
	}

	logger.fileName = fileName
	return file
}

func (logger *Logger) updateFile() {
	fileName := fmt.Sprintf(fileNamePattern, logger.fileBase, time.Now().Format("02-01-2006"))

	if logger.fileName != fileName {
		logger.writer = log.New(logger.createFile(), "", 0)
	}
}
