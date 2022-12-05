package main

import (
	"fmt"
	"github.com/lowl11/lazylog/logapi"
)

func main() {
	fmt.Println("") // just

	logger := logapi.New().File("info", "logs")
	logger.Info("hello")
}
