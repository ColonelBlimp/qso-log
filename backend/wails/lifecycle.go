//go:build windows

package wails

import (
	"context"
	"fmt"
)

func startup(ctx context.Context) {
	fmt.Println("starting wails app")
}

func shutdown(ctx context.Context) {
	fmt.Println("shutting down wails app")
}
