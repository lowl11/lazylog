package lazylog

import (
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"os"
	"strings"
)

func getFileForLog(filePath string) (*os.File, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func getFileLogWriter(file *os.File) io.Writer {
	fileWriter := &zerolog.ConsoleWriter{Out: file}

	fileWriter.FormatTimestamp = func(i interface{}) string {
		// 2021-08-18T22:40:32+06:00
		strValue := fmt.Sprintf("%s", i)
		return strings.Replace(strValue[0:19], "T", " ", -1)
	}

	fileWriter.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %s |", i))
	}
	fileWriter.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	fileWriter.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("| %s: ", i)
	}
	fileWriter.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}

	return fileWriter
}
