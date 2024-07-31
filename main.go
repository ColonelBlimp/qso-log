//go:build windows

package main

import (
	"fmt"
	"qso-log/backend/app"
	"qso-log/backend/wails"
)

func main() {
	application, err := app.Get()
	if err != nil {
		panic(fmt.Errorf("error initializing application: %w", err))
	}

	wails.Run(application)
}
