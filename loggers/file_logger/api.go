package file_logger

import "github.com/lowl11/lazylog/message_tools"

const (
	debugLevel = "[DEBUG] "
	infoLevel  = "[INFO] "
	warnLevel  = "[WARN] "
	errorLevel = "[ERROR] "
	fatalLevel = "[FATAL] "
)

func (logger *Logger) Debug(value string, args ...any) {
	logger.updateFile()
	logger.writer.SetPrefix(debugLevel)
	logger.writer.Println(message_tools.Build(value, args...))
}

func (logger *Logger) Info(value string, args ...any) {
	logger.updateFile()
	logger.writer.SetPrefix(infoLevel)
	logger.writer.Println(message_tools.Build(value, args...))
}

func (logger *Logger) Warn(value string, args ...any) {
	logger.updateFile()
	logger.writer.SetPrefix(warnLevel)
	logger.writer.Println(message_tools.Build(value, args...))
}

func (logger *Logger) Error(err error, value string, args ...any) {
	logger.updateFile()
	logger.writer.SetPrefix(errorLevel)
	logger.writer.Println(err.Error() + " | " + message_tools.Build(value, args...))
}

func (logger *Logger) Fatal(err error, value string, args ...any) {
	logger.updateFile()
	logger.writer.SetPrefix(fatalLevel)
	logger.writer.Println(err.Error() + " | " + message_tools.Build(value, args...))
}
