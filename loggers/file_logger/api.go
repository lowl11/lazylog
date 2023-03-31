package file_logger

import "github.com/lowl11/lazylog/message_tools"

const (
	debugLevel = "[DEBUG] "
	infoLevel  = "[INFO] "
	warnLevel  = "[WARN] "
	errorLevel = "[ERROR] "
	fatalLevel = "[FATAL] "
)

func (logger *Logger) Debug(args ...string) {
	logger.updateFile()
	logger.writer.SetPrefix(debugLevel)
	logger.writer.Println(message_tools.Build(args...))
}

func (logger *Logger) Info(args ...string) {
	logger.updateFile()
	logger.writer.SetPrefix(infoLevel)
	logger.writer.Println(message_tools.Build(args...))
}

func (logger *Logger) Warn(args ...string) {
	logger.updateFile()
	logger.writer.SetPrefix(warnLevel)
	logger.writer.Println(message_tools.Build(args...))
}

func (logger *Logger) Error(err error, args ...string) {
	logger.updateFile()
	logger.writer.SetPrefix(errorLevel)
	logger.writer.Println(err.Error() + " | " + message_tools.Build(args...))
}

func (logger *Logger) Fatal(err error, args ...string) {
	logger.updateFile()
	logger.writer.SetPrefix(fatalLevel)
	logger.writer.Println(err.Error() + " | " + message_tools.Build(args...))
}
