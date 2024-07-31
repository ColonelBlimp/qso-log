//go:build windows

package main

import (
	"qso-log/backend/app"
	"qso-log/backend/wails"
)

func main() {
	wails.Run(app.Get())
}
