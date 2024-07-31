//go:build windows

package db

import (
	"os"
	"path/filepath"
	"qso-log/backend/config"
	"testing"
)

func init() {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	workingDir = filepath.Join(workingDir, "..", "..", "build", "bin")

	err = os.Chdir(workingDir)
	if err != nil {
		panic(err)
	}
}

func TestInit(t *testing.T) {
	cfg := config.Get()

	err := Init(cfg.DatabaseDir(), cfg.DatabaseDriverName(), cfg.DataSourceName())
	if err != nil {
		t.Error(err)
	}
}
