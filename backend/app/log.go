//go:build windows

package app

import (
	wailsLogger "github.com/wailsapp/wails/v2/pkg/logger"
	"qso-log/backend/utils"
)

type Log interface {
	Wails() wailsLogger.Logger
}

type log struct {
	wails wailsLogger.Logger
}

func NewConsoleLogger() Log {
	return &log{
		wails: utils.NewWailsConsoleLogger("WAILS"),
	}
}

func (l *log) Wails() wailsLogger.Logger {
	return l.wails
}
