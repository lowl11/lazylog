package logapi

type ILogger interface {
	Debug(message string)
	Info(message string)
	Warn(message string)
	Error(err error, message string)
	Fatal(err error, message string)
}
