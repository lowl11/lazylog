package file_logger

const (
	debugLevel = "[DEBUG] "
	infoLevel  = "[INFO] "
	warnLevel  = "[WARN] "
	errorLevel = "[ERROR] "
	fatalLevel = "[FATAL] "
)

func (logger *Logger) Debug(message string) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()
	logger.updateFile()
	logger.writer.SetPrefix(debugLevel)
	logger.writer.Println(message)
}

func (logger *Logger) Info(message string) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()
	logger.updateFile()
	logger.writer.SetPrefix(infoLevel)
	logger.writer.Println(message)
}

func (logger *Logger) Warn(message string) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()
	logger.updateFile()
	logger.writer.SetPrefix(warnLevel)
	logger.writer.Println(message)
}

func (logger *Logger) Error(err error, message string) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()
	logger.updateFile()
	logger.writer.SetPrefix(errorLevel)
	logger.writer.Println(err.Error() + " | " + message)
}

func (logger *Logger) Fatal(err error, message string) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()
	logger.updateFile()
	logger.writer.SetPrefix(fatalLevel)
	logger.writer.Println(err.Error() + " | " + message)
}
