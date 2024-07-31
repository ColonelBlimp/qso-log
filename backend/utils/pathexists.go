//go:build windows

package utils

import (
	"errors"
	"fmt"
	"os"
)

// PathExists checks the existence of the given path (file or directory). Returns os.ErrExist if the path exists, or
// os.ErrNotExist is the path does not exist. Any other error is returned.
func PathExists(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return os.ErrExist
	}
	if errors.Is(err, os.ErrNotExist) {
		return os.ErrNotExist
	}
	return fmt.Errorf("utils.PathExists: %w", err)
}
