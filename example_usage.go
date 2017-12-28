package main

import (
	"github.com/munirehmad/golanglog/logger"
)

func main(){
	log := logger.NewLogger("./log_test.log", true)
	//log.Level = logger.LInfo
	log.Debug("This is debug", 1, true, "nice debugging", log)
	log.Info("And this is info", 1)
}
