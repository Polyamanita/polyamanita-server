package lib

import (
	"io"
	"log"
)

type Logger struct {
	debugLog *log.Logger
	errLog   *log.Logger
}

func NewLogger(out io.Writer) *Logger {
	debugLog := log.New(out, "[DEBUG] ", log.Ldate|log.Ltime|log.Llongfile)
	errLog := log.New(out, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)

	return &Logger{
		debugLog: debugLog,
		errLog:   errLog,
	}
}

func (l *Logger) Debug(v ...interface{}) {
	l.debugLog.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.errLog.Println(v...)
}
