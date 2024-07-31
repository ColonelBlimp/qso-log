//go:build windows

package config

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
	"qso-log/backend/utils"
)

const (
	configFileType = "toml"
	configFileName = "config.toml"
)

const (
	keyIsDevMode        = "isdevmode"
	keyApplicationTitle = "applicationtitle"
	keyDatabaseDir      = "databasedir"
	keyDatabaseDriver   = "databasedriver"
	keyDatabaseDSN      = "databasedsn"
)

const (
	defaultValueApplicationTitle = "QSO Logger"
	defaultValueDatabaseDriver   = "sqlite3"
	defaultDatabaseDir           = "db"
	defaultDatabaseFilename      = "data.db"
)

type Config interface {
	IsDevMode() bool
	Title() string
	DatabaseDir() string
	DatabaseDriverName() string
	DataSourceName() string
}

type config struct {
	isDevMode          bool
	title              string
	databaseDir        string
	databaseDriverName string
	dataSourceName     string
}

// _config is our internal variable.
var _config *config

// Get the application's configuration.
func Get() *config {
	if _config != nil {
		return _config
	}

	workingDir, err := utils.GetExeDir()
	if err != nil {
		panic(fmt.Errorf("config.Load: error determining the executable's directory: %w", err))
	}

	defaultConfig(workingDir)

	viper.SetConfigType(configFileType)
	viper.SetConfigFile(configFileName)
	viper.AddConfigPath(workingDir)

	return populateConfig()
}

func defaultConfig(workingDir string) {
	viper.SetDefault(keyIsDevMode, false)
	viper.SetDefault(keyApplicationTitle, defaultValueApplicationTitle)

	path, err := filepath.Abs(filepath.Join(workingDir, defaultDatabaseDir))
	if err != nil {
		panic(fmt.Errorf("config.defaultConfig: %w", err))
	}
	viper.SetDefault(keyDatabaseDir, path)
	viper.SetDefault(keyDatabaseDriver, defaultValueDatabaseDriver)
	viper.SetDefault(keyDatabaseDSN, filepath.Join(viper.GetString(keyDatabaseDir), defaultDatabaseFilename))
}

func populateConfig() *config {
	return &config{
		isDevMode:          viper.GetBool(keyIsDevMode),
		title:              viper.GetString(keyApplicationTitle),
		databaseDir:        viper.GetString(keyDatabaseDir),
		databaseDriverName: viper.GetString(keyDatabaseDriver),
		dataSourceName:     viper.GetString(keyDatabaseDSN),
	}
}

// IsDevMode returns true or false.
func (c *config) IsDevMode() bool {
	return c.isDevMode
}

// Title returns the application's title a string.
func (c *config) Title() string {
	return c.title
}

func (c *config) DatabaseDir() string {
	return c.databaseDir
}

func (c *config) DatabaseDriverName() string {
	return c.databaseDriverName
}

func (c *config) DataSourceName() string {
	return c.dataSourceName
}
