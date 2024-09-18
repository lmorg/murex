//go:build !windows
// +build !windows

package home

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

// TestMyHome tests your home directory can be derived
func TestMyHome(t *testing.T) {
	count.Tests(t, 1)
	if MyDir == "" {
		t.Error("MyDir not set (murex will still function)")
	}
}
