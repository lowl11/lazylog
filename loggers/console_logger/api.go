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

func (logger *Logger) Debug(message string, layers ...string) {
	logger.writer.SetPrefix(debugLevel)
	logger.writer.Println(message_tools.Build(console_tools.Debug(message), layers...))
}

func (logger *Logger) Info(message string, layers ...string) {
	logger.writer.SetPrefix(infoLevel)
	logger.writer.Println(message_tools.Build(console_tools.Info(message), layers...))
}

func (logger *Logger) Warn(message string, layers ...string) {
	logger.writer.SetPrefix(warnLevel)
	logger.writer.Println(message_tools.Build(console_tools.Warn(message), layers...))
}

func (logger *Logger) Error(err error, message string, layers ...string) {
	logger.writer.SetPrefix(errorLevel)
	logger.writer.Println(message_tools.Build(console_tools.Error(err.Error()+" | "+message), layers...))
}

func (logger *Logger) Fatal(err error, message string, layers ...string) {
	logger.writer.SetPrefix(fatalLevel)
	logger.writer.Println(message_tools.Build(console_tools.Fatal(err.Error()+" | "+message), layers...))
}
