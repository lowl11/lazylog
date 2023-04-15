package message_tools

import (
	"encoding/json"
	"strings"
)

var (
	JsonMode   bool
	NoTimeMode bool
)

func Build(args ...any) string {
	stringArgs := make([]string, 0, len(args))
	for _, arg := range args {
		stringArgs = append(stringArgs, toString(arg))
	}

	return strings.Join(stringArgs, " ")
}

func BuildError(err error, args ...any) string {
	stringArgs := make([]string, 0, len(args))
	for _, arg := range args {
		stringArgs = append(stringArgs, toString(arg))
	}

	var errorMessage string
	if err != nil {
		errorMessage = err.Error() + " | "
	}

	return errorMessage + strings.Join(stringArgs, " ")
}

func BuildPrefix(level string) string {
	if NoTimeMode {
		return level
	}

	return getTime() + " " + level
}

func Json(level string, args ...any) string {
	logMessage := &LogMessage{
		Message: Build(args...),
		Level:   level,
		Time:    getTime(),
	}

	logMessageInBytes, err := json.Marshal(logMessage)
	if err != nil {
		return "|ERROR IN BUILDING MESSAGE|"
	}

	return string(logMessageInBytes)
}

func JsonError(err error, level string, args ...any) string {
	var errorMessage string
	if err != nil {
		errorMessage = err.Error()
	}

	logMessage := &LogMessage{
		Message: Build(args...),
		Level:   level,
		Time:    getTime(),
		Error:   errorMessage,
	}

	logMessageInBytes, err := json.Marshal(logMessage)
	if err != nil {
		return "|ERROR IN BUILDING MESSAGE|"
	}

	return string(logMessageInBytes)
}
