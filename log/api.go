package log

import "github.com/lowl11/lazylog/log/log_internal"

func Debug(args ...any) {
	log_internal.Debug(args...)
}

func Info(args ...any) {
	log_internal.Info(args...)
}

func Warn(args ...any) {
	log_internal.Warn(args...)
}

func Error(err error, args ...any) {
	log_internal.Error(err, args...)
}

func Fatal(err error, args ...any) {
	log_internal.Fatal(err, args...)
}
