package logapi

import (
	"github.com/lowl11/lazylog/internal/line_event"
	"github.com/lowl11/lazylog/internal/loggers/console_logger"
	"github.com/lowl11/lazylog/internal/loggers/file_logger"
	"github.com/lowl11/lazylog/internal/message_tools"
	"github.com/lowl11/lazylog/logapi/log_levels"
	"os"
	"sync"
	"time"
)

type Logger struct {
	loggers       []ILogger
	customLoggers []ILogger
	mutex         sync.Mutex
	line          *line_event.Event

	exitDuration     time.Duration
	customDuration   time.Duration
	isCustomDuration bool

	level uint
}

func New() *Logger {
	return &Logger{
		loggers: []ILogger{
			console_logger.Create(),
		},
		mutex:        sync.Mutex{},
		exitDuration: time.Millisecond * defaultExitDuration,
	}
}

func (logger *Logger) Level(level uint) *Logger {
	if level > log_levels.FATAL {
		return logger
	}

	logger.level = level
	return logger
}

func (logger *Logger) File(fileBase string, filePath ...string) *Logger {
	var singleFilePath string
	if len(filePath) > 0 {
		singleFilePath = filePath[0]
	}

	logger.loggers = append(logger.loggers, file_logger.Create(fileBase, singleFilePath))
	return logger
}

func (logger *Logger) Custom(customLogger ILogger) *Logger {
	logger.customLoggers = append(logger.customLoggers, customLogger)
	if !logger.isCustomDuration {
		logger.exitDuration = logger.exitDuration + time.Millisecond*defaultExitDuration
	}

	if logger.line == nil && len(logger.customLoggers) > 0 {
		logger.line = line_event.New()
	}
	return logger
}

func (logger *Logger) CustomExitDuration(duration time.Duration) *Logger {
	if duration < defaultExitDuration {
		return logger
	}

	logger.isCustomDuration = true
	logger.customDuration = duration

	return logger
}

func (logger *Logger) JSON() *Logger {
	message_tools.JsonMode = true
	return logger
}

func (logger *Logger) NoTime() *Logger {
	message_tools.NoTimeMode = true
	return logger
}

func (logger *Logger) NoPrefix() *Logger {
	message_tools.NoPrefixMode = true
	return logger
}

func (logger *Logger) Debug(args ...any) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	// skip log by level
	if log_levels.Check(logger.level, log_levels.DEBUG) {
		return
	}

	for _, loggerItem := range logger.loggers {
		loggerItem.Debug(args...)
	}

	for _, customLogger := range logger.customLoggers {
		logger.line.AddInfo(customLogger.Debug, args...)
	}
}

func (logger *Logger) Info(args ...any) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	// skip log by level
	if log_levels.Check(logger.level, log_levels.INFO) {
		return
	}

	for _, loggerItem := range logger.loggers {
		loggerItem.Info(args...)
	}

	for _, customLogger := range logger.customLoggers {
		logger.line.AddInfo(customLogger.Info, args...)
	}
}

func (logger *Logger) Warn(args ...any) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	// skip log by level
	if log_levels.Check(logger.level, log_levels.WARN) {
		return
	}

	for _, loggerItem := range logger.loggers {
		loggerItem.Warn(args...)
	}

	for _, customLogger := range logger.customLoggers {
		logger.line.AddInfoCustom(customLogger.Warn, args...)
	}
}

func (logger *Logger) Error(err error, args ...any) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	// skip log by level
	if log_levels.Check(logger.level, log_levels.ERROR) {
		return
	}

	for _, loggerItem := range logger.loggers {
		loggerItem.Error(err, args...)
	}

	for _, customLogger := range logger.customLoggers {
		logger.line.AddErrorCustom(customLogger.Error, err, args...)
	}
}

func (logger *Logger) Fatal(err error, args ...any) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	for _, loggerItem := range logger.loggers {
		loggerItem.Fatal(err, args...)
	}

	for _, customLogger := range logger.customLoggers {
		logger.line.AddErrorCustom(customLogger.Fatal, err, args...)
	}

	if logger.isCustomDuration {
		time.Sleep(logger.customDuration)
	} else {
		time.Sleep(logger.exitDuration)
	}
	os.Exit(1)
}
