//go:build windows

package wails

import (
	"embed"
	"fmt"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"qso-log/backend/app"
)

//go:embed all:assets
var assets embed.FS

func Run(application app.App) error {
	opts := &options.App{
		Title:             application.Config().Title(),
		Width:             1024, // 16:10
		Height:            640,  // 16:10
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		MinWidth:          1024, // 16:10
		MinHeight:         640,  // 16:10
		MaxWidth:          -1,
		MaxHeight:         -1,
		StartHidden:       false,
		HideWindowOnClose: false,
		AlwaysOnTop:       false,
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		// RGBA:              &options.RGBA{},
		Menu: nil,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},

		//		Logger:             app.App().Log().Wails(),
		//		LogLevel:           logger.INFO,
		//		LogLevelProduction: logger.ERROR,
		OnStartup: startup,
		//		OnDomReady:    domReady,
		OnShutdown: shutdown,
		//		OnBeforeClose: beforeClose,
		Bind: []interface{}{
			application,
			//NewBinding(),
			//tray.Tray(),
			//local.Service(),
		},
		WindowStartState: options.Normal,
		Windows: &windows.Options{
			WebviewIsTransparent:              true,
			WindowIsTranslucent:               false,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: false,
			WebviewUserDataPath:               "",
			WebviewBrowserPath:                "",
			Theme:                             windows.SystemDefault,
			CustomTheme:                       nil, /*&windows.ThemeSettings{
			  // Theme to use when window is active
			  DarkModeTitleBar:   windows.RGB(255, 0, 0), // Red
			  DarkModeTitleText:  windows.RGB(0, 255, 0), // Green
			  DarkModeBorder:     windows.RGB(0, 0, 255), // Blue
			  LightModeTitleBar:  windows.RGB(200, 200, 200),
			  LightModeTitleText: windows.RGB(20, 20, 20),
			  LightModeBorder:    windows.RGB(200, 200, 200),
			  // Theme to use when window is inactive
			  DarkModeTitleBarInactive:   windows.RGB(128, 0, 0),
			  DarkModeTitleTextInactive:  windows.RGB(0, 128, 0),
			  DarkModeBorderInactive:     windows.RGB(0, 0, 128),
			  LightModeTitleBarInactive:  windows.RGB(100, 100, 100),
			  LightModeTitleTextInactive: windows.RGB(10, 10, 10),
			  LightModeBorderInactive:    windows.RGB(100, 100, 100),
			},*/
			Messages:         nil,
			ResizeDebounceMS: 0,
			//			OnSuspend:        suspend,
			//			OnResume:         resume,
		},
	}

	if err := wails.Run(opts); err != nil {
		return fmt.Errorf("error while starting Wails: %w", err)
	}

	return nil
}
