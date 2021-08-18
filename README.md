# Lazy Log

Logger package wrapper on ZeroLog for Golang

```go
logger := lazylog.CreateLogger(&lazylog.LoggerSettings{
	Debug: true,
	FileLogger: true,
	FilePath: "info.log",
})

logger.Info().Interface("some_var", "some_value").Msg("Hello World")
```
