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
	return &Logger{
		Logger: log.New(os.Stdout, buildPrefix(prefix), log.Ldate|log.Ltime|log.Llongfile),
	}
}

func NewFileLogger(prefix string, f *os.File) *Logger {
	return &Logger{
		Logger: log.New(f, buildPrefix(prefix), log.Ldate|log.Ltime|log.Llongfile),
	}
}

func buildPrefix(prefix string) string {
	if prefix == "" {
		return prefix
	}
	return "[" + prefix + "] "
}
