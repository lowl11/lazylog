package main

import "github.com/lowl11/lazylog/logapi"

func main() {
	logger := logapi.New().
		File("info")

	logger.Info("123")
}
