//go:build windows

package app

import (
	wailsLogger "github.com/wailsapp/wails/v2/pkg/logger"
)

type Log struct {
	wails wailsLogger.Logger
}
