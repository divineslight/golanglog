# GoLangLog
Minimalistic Golang logger that logs to stdout and files.

## Installation
    go get github.com/munirehmad/golanglog
    
## Usage
    package main
    
    import (
        "github.com/munirehmad/golanglog/logger"
    )
      
    func main(){
        log := logger.NewLogger("./logs/main.log", true)
        log.Level = logger.LInfo

        // This will be ignored because log level is set at Info`
        log.Debug("This is debug", 1, true, log)

        log.Info("And this is info", 1)
    }
