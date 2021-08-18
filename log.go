package main

import (
	"github.com/rs/zerolog"
	"io"
)

type LoggerSettings struct {
	Debug      bool
	FileLogger bool
	FilePath   string
}

func CreateLogger(settings *LoggerSettings) *zerolog.Logger {
	writers := make([]io.Writer, 0, 3)

	defaultConsoleWriter := getConsoleLogWriter()
	writers = append(writers, defaultConsoleWriter)

	if settings.FileLogger {
		file, err := getFileForLog(settings.FilePath)
		if err != nil {
			defaultLogger := zerolog.New(defaultConsoleWriter).With().Timestamp().Logger()
			defaultLogger.Fatal().Err(err).Msgf("[Lazy Log] Opening log file '%s' error", settings.FilePath)
		}

		writers = append(writers, getFileLogWriter(file))
	}

	multiWriter := zerolog.MultiLevelWriter(writers...)
	logger := zerolog.New(multiWriter).With().Timestamp().Logger()

	if settings.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	return &logger
}
