package logger

import (
	"os"
	"time"
	"fmt"
	"strings"
)

const (
	LDebug = iota
	LInfo
	LWarn
	LError
	LFatal
)

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
		logger.File, err = os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
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
	// 2017/12/28 16:53:12
	tfmt := "2006/01/2 15:04:05 MST"
	prefix := time.Now().Format(tfmt)
	msg := "" + tag + "|" + "" + prefix + "| " + fmtLogMsg(args...)
	fmt.Println(msg)

	if l.File != nil {
		l.File.WriteString(msg + "\n")
		l.File.Sync()
	}
}

func fmtLogMsg(args ...interface{}) string {
	var argsStr = []string{}

	for _, v := range args {
		argsStr = append(argsStr, fmt.Sprintf("%#v", v))
	}

	return strings.Join(argsStr, ", ")
}

func check(e error){
	if e != nil {
		panic(e)
	}
}
