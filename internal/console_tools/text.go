package console_tools

func Debug(text string) string {
	return purple + text + reset
}

func Info(text string) string {
	return green + text + reset
}

func Warn(text string) string {
	return yellow + text + reset
}

func Error(text string) string {
	return red + text + reset
}

func Fatal(text string) string {
	return gray + text + reset
}
