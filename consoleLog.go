package main

import (
	"github.com/rs/zerolog"
	"os"
)

func getConsoleLogWriter() *zerolog.ConsoleWriter {
	return &zerolog.ConsoleWriter{Out: os.Stdout}
}
