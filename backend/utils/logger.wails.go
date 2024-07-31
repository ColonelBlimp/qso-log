//go:build windows

package utils

import (
	"github.com/wailsapp/wails/v2/pkg/logger"
	"log"
)

type WailsLogger struct {
	log *log.Logger
}

func (w *WailsLogger) Print(message string) {
	w.log.Println(message)
}

func (w *WailsLogger) Trace(message string) {
	w.log.Println("TRACE: " + message)
}

func (w *WailsLogger) Debug(message string) {
	w.log.Println("DEBUG: " + message)
}

func (w *WailsLogger) Info(message string) {
	w.log.Println("INFO: " + message)
}

func (w *WailsLogger) Warning(message string) {
	w.log.Println("WARN: " + message)
}

func (w *WailsLogger) Error(message string) {
	w.log.Println("ERROR: " + message)
}

func (w *WailsLogger) Fatal(message string) {
	w.log.Println("FATAL: " + message)
}

func NewWailsConsoleLogger(prefix string) logger.Logger {
	return &WailsLogger{
		log: NewConsoleLogger(prefix).Logger,
	}
}
