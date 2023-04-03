package message_tools

import "strings"

func Build(value string, args ...any) string {
	stringArgs := make([]string, 0, len(args))
	for _, arg := range args {
		stringArgs = append(stringArgs, toString(arg))
	}
	return value + " " + strings.Join(stringArgs, " ")
}
