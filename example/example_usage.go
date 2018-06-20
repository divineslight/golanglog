package main

import (
	"errors"

	"github.com/munirehmad/golanglog"
)

func main() {
	log := golanglog.NewLogger("./log_test.log", true)
	//log.Level = logger.LInfo
	log.Debug("This is debug", 1, true, "nice debugging", log)
	log.Info("And this is info", 1, errors.New("test error"))
}
