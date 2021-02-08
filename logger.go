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

func Success(msg string, params ...interface{}) {
	Log(msg, LogTypeSuccess, params...)
}

func Warning(msg string, params ...interface{}) {
	Log(msg, LogTypeWarning, params...)
}

func Error(msg string, params ...interface{}) {
	Log(msg, LogTypeError, params...)
}

func Info(msg string, params ...interface{}) {
	Log(msg, LogTypeInfo, params...)
}

func Debug(msg string, params ...interface{}) {
	Log(msg, LogTypeDebug, params...)
}

func Dump(i interface{}) {
	b, err := json.Marshal(i)

	if err != nil {
		LogErr(err)
	}

	Log(string(b), LogTypeDebug)
}

func Log(msg string, logType LogType, params ...interface{}) {
	fmt.Printf("%s %s: %s\n", time.Now().Format(time.RFC3339), getLogTypeHeader(logType), fmt.Sprintf(msg, params...))
}

func LogErr(err error) {
	if err == nil {
		return
	}
	_, fileName, lineNumber, _ := runtime.Caller(1)

	fmt.Printf("%s %s: %s\n at: %s:%d \n",
		time.Now().Format(time.RFC3339),
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
