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

// Config the interface to the application's configuration.
type Config interface {
	IsDevMode() bool
	Title() string
	DatabaseDir() string
	DatabaseDriverName() string
	DataSourceName() string
}

// config the actual internal data.
type config struct {
	isDevMode          bool
	title              string
	databaseDir        string
	databaseDriverName string
	dataSourceName     string
}

// _config is our internal variable.
var _config Config

// Get the application's configuration.
func Get() Config {
	if _config != nil {
		return _config
	}

	workingDir, err := utils.GetExeDir()
	if err != nil {
		panic(fmt.Errorf("config.Get: %w", err))
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

	// See https://github.com/mattn/go-sqlite3?tab=readme-ov-file#connection-string parameters.
	viper.SetDefault(keyDatabaseDSN, fmt.Sprintf("%s?_fk=true", filepath.Join(viper.GetString(keyDatabaseDir), defaultDatabaseFilename)))
}

func populateConfig() Config {
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

// DatabaseDir returns the absolute path to the directory containing the database file.
func (c *config) DatabaseDir() string {
	return c.databaseDir
}

// DatabaseDriverName returns the database driver name. Default is 'sqlite3'.
func (c *config) DatabaseDriverName() string {
	return c.databaseDriverName
}

// DataSourceName returns the database DSN.
func (c *config) DataSourceName() string {
	return c.dataSourceName
}
