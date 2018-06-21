# GoLangLog
Minimalistic Golang logger that logs to stdout and files.

## Installation
    go get github.com/munirehmad/golanglog
    
## Usage
    package main
    
    import (
        "github.com/munirehmad/golanglog"
    )
      
    func main(){
        log := golanglog.NewLogger("./logs/main.log", true)

        /* Available Logging Levels are
        LDebug
        LInfo
        LWarn
        LError
        LFatal*/
        
        log.Level = golanglog.LInfo

        // This will be ignored because log level is set at Info`
        log.Debug("This is debug", 1, true, log)

        log.Info("And this is info", 1)
        
        /* Available Logging functions are
        log.Debug
        log.Info
        log.Warn
        log.Error
        log.Fatal // This kills the execution with panic()*/
    }
