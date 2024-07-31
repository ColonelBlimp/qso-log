//go:build windows

package app

import (
	"context"
	"qso-log/backend/config"
)

type App interface {
	Config() config.Config
	Log() *applog
	Ctx() context.Context
}

type app struct {
	config config.Config
	log    *applog         // Logger for the whole application
	ctx    context.Context // wails context
}

var _app *app

// Init initializes the application.
func Init() (*app, error) {
	if _app != nil {
		return _app, nil
	}
	_app = &app{
		config: config.Get(),
	}

	return _app, nil
}

func (a *app) Config() config.Config {
	return a.config
}

func (a *app) Log() *applog {
	return a.log
}

func (a *app) Ctx() context.Context {
	return a.ctx
}

func (a *app) SetCtx(ctx context.Context) {
	a.ctx = ctx
}
