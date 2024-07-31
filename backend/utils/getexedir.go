//go:build windows

package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

const errorFormatGetExeDir = "utils.GetExeDir: error determining the executable's directory: %w"

// GetExeDir returns the absolute path of the executable.
func GetExeDir() (string, error) {
	path, err := os.Executable()
	if err != nil {
		return path, fmt.Errorf(errorFormatGetExeDir, err)
	}
	path, err = filepath.Abs(filepath.Dir(path))
	if err != nil {
		return path, fmt.Errorf(errorFormatGetExeDir, err)
	}
	return path, nil
}
