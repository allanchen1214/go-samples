package main

import (
	"github.com/allanchen1214/bingo/log"
	"go.uber.org/zap"
)

var (
	consoleLogger *zap.Logger
)

func init() {
	consoleLogger = log.ConsoleLogger()
}

func main() {
	consoleLogger.Info("123456")

	log.Info("main start")
	log.Info("main end")

	testRotate()
}

func testRotate() {
	for i := 0; i < 1000000; i++ {
		log.Info("infoMsg")
		log.Error("errorMsg")
	}
}
