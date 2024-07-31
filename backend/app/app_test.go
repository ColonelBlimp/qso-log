//go:build windows

package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGet(t *testing.T) {
	assert.NotPanics(t, func() {
		Get()
	}, "The code paniced")
}
