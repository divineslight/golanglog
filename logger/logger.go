package logger

import (
	"os"
	"time"
	"fmt"
)

const LDebug = 0
const LInfo = 1
const LWarn = 2
const LError = 3
const LFatal = 4

type Logger struct {
	FileName string
	File *os.File
	Level uint
	PrintToStdout bool
}

func NewLogger(file string, printToStdout bool) *Logger {
	logger := Logger{FileName: file, PrintToStdout: printToStdout}
	var err error

	if file != "" {
		logger.File, err = os.Create(file)
		check(err)
	}

	return &logger
}

func (l Logger) Debug(args ...interface{}) {
	if l.Level <= LDebug {
		l.Log("DEBUG", args...)
	}
}

func (l Logger) Info(args ...interface{}) {
	if l.Level <= LInfo {
		l.Log("INFO", args...)
	}
}

func (l Logger) Warn(args ...interface{}) {
	if l.Level <= LInfo {
		l.Log("WARN", args...)
	}
}

func (l Logger) Error(args ...interface{}) {
	if l.Level <= LInfo {
		l.Log("ERROR", args...)
	}
}

func (l Logger) Fatal(args ...interface{}) {
	if l.Level <= LInfo {
		l.Log("FATAL", args...)
	}
}

func (l Logger) Log(tag string, args ...interface{}){
	prefix := time.Now().String() + ": "
	msg := "[" + tag + "] " + prefix + fmtLogMsg(args...)
	fmt.Println(msg)

	if l.File != nil {
		l.File.WriteString(msg + "\n")
		l.File.Sync()
	}
}

func fmtLogMsg(args ...interface{}) string {
	out := ""

	for _, v := range args {
		out += fmt.Sprintf("%T(%+v), ", v, v)
	}

	return out
}

func check(e error){
	if e != nil {
		panic(e)
	}
}
