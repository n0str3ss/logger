package logger

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"
)

type LogType int

const (
	LogTypeSuccess LogType = 1
	LogTypeWarning         = 2
	LogTypeError           = 3
	LogTypeInfo            = 4
	LogTypeDebug           = 5
)

func Success(msg string) {
	Log(msg, LogTypeSuccess)
}

func Warning(msg string) {
	Log(msg, LogTypeWarning)
}

func Error(msg string) {
	Log(msg, LogTypeError)
}

func Info(msg string) {
	Log(msg, LogTypeInfo)
}

func Debug(msg string) {
	Log(msg, LogTypeDebug)
}

func Dump(i interface{}) {
	b, err := json.Marshal(i)

	if err != nil {
		LogErr(err)
	}

	Log(string(b), LogTypeDebug)
}

func Log(msg string, logType LogType) {
	fmt.Printf("%s %s: %s\n", time.Now().Format("2006-01-02 15:04:05"), getLogTypeHeader(logType), msg)
}

func LogErr(err error) {
	if err == nil {
		return
	}
	_, fileName, lineNumber, _ := runtime.Caller(1)

	fmt.Printf("%s %s: %s\n at: %s:%d \n",
		time.Now().Format("2006-01-02 15:04:05"),
		getLogTypeHeader(LogTypeError),
		err.Error(),
		fileName,
		lineNumber)
}

func getLogTypeHeader(t LogType) string {
	switch t {
	case LogTypeSuccess:
		return "[SUCCESS]"
	case LogTypeWarning:
		return "[WARNING]"
	case LogTypeError:
		return "[ERROR]"
	case LogTypeInfo:
		return "[INFO]"
	case LogTypeDebug:
		return "[DEBUG]"
	}

	return ""
}