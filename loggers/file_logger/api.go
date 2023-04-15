package file_logger

import "github.com/lowl11/lazylog/message_tools"

const (
	debugLevel = "[DEBUG] "
	infoLevel  = "[INFO] "
	warnLevel  = "[WARN] "
	errorLevel = "[ERROR] "
	fatalLevel = "[FATAL] "
)

func (logger *Logger) Debug(args ...any) {
	logger.updateFile()
	logger.writer.SetPrefix(message_tools.BuildPrefix(debugLevel))
	logger.writer.Println(message_tools.Build(args...))
}

func (logger *Logger) Info(args ...any) {
	logger.updateFile()
	logger.writer.SetPrefix(message_tools.BuildPrefix(infoLevel))
	logger.writer.Println(message_tools.Build(args...))
}

func (logger *Logger) Warn(args ...any) {
	logger.updateFile()
	logger.writer.SetPrefix(message_tools.BuildPrefix(warnLevel))
	logger.writer.Println(message_tools.Build(args...))
}

func (logger *Logger) Error(err error, args ...any) {
	logger.updateFile()
	logger.writer.SetPrefix(message_tools.BuildPrefix(errorLevel))
	logger.writer.Println(err.Error() + message_tools.BuildError(args...))
}

func (logger *Logger) Fatal(err error, args ...any) {
	logger.updateFile()
	logger.writer.SetPrefix(message_tools.BuildPrefix(fatalLevel))
	logger.writer.Println(err.Error() + message_tools.BuildError(args...))
}
