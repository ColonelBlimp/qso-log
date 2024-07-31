//go:build windows

package config

import "testing"

func TestConfigLoad(t *testing.T) {
	cfg := Get()
	if cfg == nil {
		t.Error("Load Config fail")
	}

	if cfg.Title() != "QSO Logger" {
		t.FailNow()
	}
}
