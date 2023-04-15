package message_tools

import (
	"strings"
	"time"
)

var (
	NoTimeMode bool
)

func Build(args ...any) string {
	stringArgs := make([]string, 0, len(args))
	for _, arg := range args {
		stringArgs = append(stringArgs, toString(arg))
	}

	return strings.Join(stringArgs, " ")
}

func BuildError(args ...any) string {
	if len(args) == 0 {
		return ""
	}

	stringArgs := make([]string, 0, len(args))
	for _, arg := range args {
		stringArgs = append(stringArgs, toString(arg))
	}

	return " | " + strings.Join(stringArgs, " ")
}

func BuildPrefix(level string) string {
	if NoTimeMode {
		return level
	}

	return time.Now().Format("02-01-2006 15:04:05") + " " + level
}
