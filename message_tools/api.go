package message_tools

import "strings"

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
