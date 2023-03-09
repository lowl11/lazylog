package logapi

type ILogger interface {
	Debug(args ...string)
	Info(args ...string)
	Warn(args ...string)
	Error(err error, args ...string)
	Fatal(err error, args ...string)
}
