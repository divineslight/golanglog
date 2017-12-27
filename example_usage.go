package main

import (
	"github.com/munirehmad/golanglog/logger"
)

func main(){
	logger := logger.NewLogger("./log_test.log", true)
	logger.Debug("This is debug", 1, true, "nice debugging", logger)
	logger.Info("And this is info", 1)
}
