//go:build windows

package wails

import (
	"context"
	"fmt"
	"qso-log/backend/app"
)

func startup(ctx context.Context) {
	app.Get().SetCtx(ctx)
	fmt.Println("starting wails app")
}

func shutdown(ctx context.Context) {
	fmt.Println("shutting down wails app")
}
