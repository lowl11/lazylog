package logapi

type ILogger interface {
	Debug(message string, layers ...string)
	Info(message string, layers ...string)
	Warn(message string, layers ...string)
	Error(err error, message string, layers ...string)
	Fatal(err error, message string, layers ...string)
}
