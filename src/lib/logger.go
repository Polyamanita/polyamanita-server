package lib

import (
	"fmt"
	"io"
	"log"
	"runtime"
)

type Logger struct {
	debugLog *log.Logger
	errLog   *log.Logger
}

func NewLogger(out io.Writer) *Logger {
	debugLog := log.New(out, "[DEBUG] ", log.Ldate|log.Ltime)
	errLog := log.New(out, "[ERROR] ", log.Ldate|log.Ltime)

	return &Logger{
		debugLog: debugLog,
		errLog:   errLog,
	}
}

func (l *Logger) Debug(v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	p := []interface{}{fmt.Sprintf("%v:%v:", file, line)}
	p = append(p, v)
	l.debugLog.Println(p...)
}

func (l *Logger) Error(v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	p := []interface{}{fmt.Sprintf("%v:%v:", file, line)}
	p = append(p, v)
	l.errLog.Println(p...)
}
