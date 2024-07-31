//go:build windows

package app

import (
	"context"
	"fmt"
	"qso-log/backend/config"
	"qso-log/backend/db"
)

const (
	errorFormatGet = "app.Get: %w"
)

// App is the external interface to the application. This holds some of the global resources and states.
type App interface {
	Config() config.Config
	Log() *applog
	Ctx() context.Context
	SetCtx(ctx context.Context)
}

// app the actual internal data.
type app struct {
	cfg config.Config
	log *applog         // Logger for the whole application
	ctx context.Context // wails context
}

// _app is our internal variable.
var _app App

// Get returns a pointer to the application. The first time this is called the application is initialized.
// Subsequent calls returns the pointer to the previously initialized application.
//
// The initialization process is: initialize configuration, initialize logging, initialize database and initialize i18n.
// Errors are handle by panic().
func Get() App {
	if _app != nil {
		return _app
	}
	_app = &app{
		cfg: config.Get(),
	}

	err := db.Init(_app.Config().DatabaseDir(), _app.Config().DatabaseDriverName(), _app.Config().DataSourceName())
	if err != nil {
		panic(fmt.Errorf(errorFormatGet, err))
	}

	return _app
}

// Config returns a pointer to the configuration object.
func (a *app) Config() config.Config {
	return a.cfg
}

// Log returns a pointer to the logging object.
func (a *app) Log() *applog {
	return a.log
}

// Ctx return a context.Context from wails.
func (a *app) Ctx() context.Context {
	return a.ctx
}

// SetCtx sets the context.Context from wails.
func (a *app) SetCtx(ctx context.Context) {
	a.ctx = ctx
}
