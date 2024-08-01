//go:build windows

package app

import (
	wailsLogger "github.com/wailsapp/wails/v2/pkg/logger"
	"qso-log/backend/utils"
)

const (
	LogPrefixWails = "WAILS"
)

// Log the external interface to the application's logger. This holds various specific loggers for various systems.
type Log interface {
	Wails() wailsLogger.Logger
}

// log the actual internal data.
type log struct {
	wails wailsLogger.Logger
}

// NewConsoleLogger initialize a new console logger. Each specific system logger is given its own prefix.
func NewConsoleLogger() Log {
	return &log{
		wails: utils.NewWailsConsoleLogger(LogPrefixWails),
	}
}

// Wails is the wails specific logger.
func (l *log) Wails() wailsLogger.Logger {
	return l.wails
}
