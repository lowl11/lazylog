package console_logger

import "github.com/lowl11/lazylog/console_tools"

const (
	debugLevel = "[DEBUG] "
	infoLevel  = "[INFO] "
	warnLevel  = "[WARN] "
	errorLevel = "[ERROR] "
	fatalLevel = "[FATAL] "
)

func (logger *Logger) Debug(message string) {
	logger.writer.SetPrefix(debugLevel)
	logger.writer.Println(console_tools.Debug(message))
}

func (logger *Logger) Info(message string) {
	logger.writer.SetPrefix(infoLevel)
	logger.writer.Println(console_tools.Info(message))
}

func (logger *Logger) Warn(message string) {
	logger.writer.SetPrefix(warnLevel)
	logger.writer.Println(console_tools.Warn(message))
}

func (logger *Logger) Error(err error, message string) {
	logger.writer.SetPrefix(errorLevel)
	logger.writer.Println(console_tools.Error(err.Error() + " | " + message))
}

func (logger *Logger) Fatal(err error, message string) {
	logger.writer.SetPrefix(fatalLevel)
	logger.writer.Println(console_tools.Fatal(err.Error() + " | " + message))
}
