//go:build windows

package utils

import (
	"log"
	"os"
)

type Logger struct {
	*log.Logger
}

func NewConsoleLogger(prefix string) *Logger {
	l := &Logger{}
	l.Logger = log.New(os.Stdout, l.prefix(prefix), log.Ldate|log.Ltime|log.Llongfile)
	return l
}

func NewFileLogger(prefix string, f *os.File) *Logger {
	l := &Logger{}
	l.Logger = log.New(f, l.prefix(prefix), log.Ldate|log.Ltime|log.Llongfile)
	return l
}

func (l *Logger) prefix(prefix string) string {
	if prefix == "" {
		return prefix
	}
	return "[" + prefix + "] "
}
