package message_tools

import "strings"

func Build(args ...string) string {
	return strings.Join(args, " ")
}
