//go:build windows

package main

import (
	"errors"
	"fmt"
	"qso-log/backend/app"
	"qso-log/backend/wails"
)

func main() {
	if err := wails.Run(app.Get()); err != nil { //nolint:errcheck
		fmt.Println("Exited with an error: %w", errors.Unwrap(err))
	}
}
