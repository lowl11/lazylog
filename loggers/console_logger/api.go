package console_logger

import (
	"github.com/lowl11/lazylog/console_tools"
	"github.com/lowl11/lazylog/message_tools"
)

const (
	debugLevel = "[DEBUG] "
	infoLevel  = "[INFO] "
	warnLevel  = "[WARN] "
	errorLevel = "[ERROR] "
	fatalLevel = "[FATAL] "
)

func (logger *Logger) Debug(args ...any) {
	logger.writer.SetPrefix(message_tools.BuildPrefix(debugLevel))
	logger.writer.Println(console_tools.Debug(message_tools.Build(args...)))
}

func (logger *Logger) Info(args ...any) {
	logger.writer.SetPrefix(message_tools.BuildPrefix(infoLevel))
	logger.writer.Println(console_tools.Info(message_tools.Build(args...)))
}

func (logger *Logger) Warn(args ...any) {
	logger.writer.SetPrefix(message_tools.BuildPrefix(warnLevel))
	logger.writer.Println(console_tools.Warn(message_tools.Build(args...)))
}

func (logger *Logger) Error(err error, args ...any) {
	logger.writer.SetPrefix(message_tools.BuildPrefix(errorLevel))
	logger.writer.Println(console_tools.Error(err.Error() + message_tools.BuildError(args...)))
}

func (logger *Logger) Fatal(err error, args ...any) {
	logger.writer.SetPrefix(message_tools.BuildPrefix(fatalLevel))
	logger.writer.Println(console_tools.Fatal(err.Error() + message_tools.BuildError(args...)))
}
