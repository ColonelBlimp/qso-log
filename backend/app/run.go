//go:build windows

package app

import (
	"context"
	"qso-log/backend/config"
)

type App interface {
	Config() config.AppConfig
	Ctx() context.Context
}

type app struct {
	config config.AppConfig
	ctx    context.Context // wails context
}

var _app *app

// Init initializes the application.
func Init() (*app, error) {
	if _app != nil {
		return _app, nil
	}
	_app = &app{
		config: config.Load(),
	}

	return _app, nil
}

func (a *app) Config() config.AppConfig {
	return a.config
}

func (a *app) Ctx() context.Context {
	return a.ctx
}
