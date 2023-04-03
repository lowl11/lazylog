package logapi

type ILogger interface {
	Debug(value string, args ...any)
	Info(value string, args ...any)
	Warn(value string, args ...any)
	Error(err error, value string, args ...any)
	Fatal(err error, value string, args ...any)
}
