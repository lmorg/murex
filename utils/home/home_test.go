//go:build !windows
// +build !windows

package home

import (
	"os/user"
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

// TestUserHome tests a users home directory can be derived
func TestUserHome(t *testing.T) {
	count.Tests(t, 1)

	u, err := user.Current()
	if err != nil {
		t.Error("Error querying user.Current: " + err.Error() + " (murex will still function)")
	}

	if UserDir(u.Username) == "" {
		t.Error("UserDir returning empty string (murex will still function)")
	}
}
