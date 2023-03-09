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

func (logger *Logger) Debug(args ...string) {
	logger.writer.SetPrefix(debugLevel)
	logger.writer.Println(console_tools.Debug(message_tools.Build(args...)))
}

func (logger *Logger) Info(args ...string) {
	logger.writer.SetPrefix(infoLevel)
	logger.writer.Println(console_tools.Info(message_tools.Build(args...)))
}

func (logger *Logger) Warn(args ...string) {
	logger.writer.SetPrefix(warnLevel)
	logger.writer.Println(console_tools.Warn(message_tools.Build(args...)))
}

func (logger *Logger) Error(err error, args ...string) {
	logger.writer.SetPrefix(errorLevel)
	logger.writer.Println(console_tools.Error(err.Error() + " | " + message_tools.Build(args...)))
}

func (logger *Logger) Fatal(err error, args ...string) {
	logger.writer.SetPrefix(fatalLevel)
	logger.writer.Println(console_tools.Fatal(err.Error() + " | " + message_tools.Build(args...)))
}
