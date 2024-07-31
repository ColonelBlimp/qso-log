//go:build windows

package config

type AppConfig interface {
	Title() string
}

type appConfig struct {
	title string
}

var _config *appConfig

func Load() *appConfig {
	if _config != nil {
		return _config
	}
	_config = &appConfig{
		title: "QSO Logger",
	}

	return _config
}

func (c *appConfig) Title() string {
	return c.title
}
