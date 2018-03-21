package logger

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	LDebug = iota
	LInfo
	LWarn
	LError
	LFatal
)

type logger struct {
	FileName      string
	Level         uint
	PrintToStdout bool

	file *os.File
}

func NewLogger(file string, printToStdout bool) *logger {
	logger := logger{FileName: file, PrintToStdout: printToStdout}
	var err error

	if file != "" {
		logger.file, err = os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		check(err)
	}

	return &logger
}

func (l *logger) Debug(args ...interface{}) {
	if l.Level <= LDebug {
		l.Log("DEBUG", args...)
	}
}

func (l *logger) Info(args ...interface{}) {
	if l.Level <= LInfo {
		l.Log("INFO", args...)
	}
}

func (l *logger) Warn(args ...interface{}) {
	if l.Level <= LInfo {
		l.Log("WARN", args...)
	}
}

func (l *logger) Error(args ...interface{}) {
	if l.Level <= LInfo {
		l.Log("ERROR", args...)
	}
}

func (l *logger) Fatal(args ...interface{}) {
	if l.Level <= LInfo {
		panic(l.Log("FATAL", args...))
	}
}

func (l *logger) Log(tag string, args ...interface{}) (msg string) {
	// 2017/12/28 16:53:12
	timeFmt := "2006/01/2 15:04:05 MST"
	timeStamp := time.Now().Format(timeFmt)
	msg = "" + tag + "|" + "" + timeStamp + "| " + fmtLogMsg(args...) + "\n"

	// log to stdout
	fmt.Print(msg)

	// log to file if set
	if l.file != nil {
		l.file.WriteString(msg)
		l.file.Sync()
	}

	return
}

func fmtLogMsg(args ...interface{}) string {
	var argsStr = []string{}

	for _, v := range args {
		// error needs special treatment
		if _, ok := v.(error); ok {
			argsStr = append(argsStr, fmt.Sprintf("%v", v))
			continue
		}
		argsStr = append(argsStr, fmt.Sprintf("%+v", v))
	}

	return strings.Join(argsStr, ", ")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
