//go:build windows

package app

import (
	wailsLogger "github.com/wailsapp/wails/v2/pkg/logger"
)

type applog struct {
	wails wailsLogger.Logger
}
