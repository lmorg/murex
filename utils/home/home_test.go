// +build !windows

package home

import (
	"os/user"
	"testing"
)

// TestHome tests the home directories can be derived
func TestHome(t *testing.T) {
	if MyDir == "" {
		t.Error("MyDir not set (murex will still function)")
	}

	u, err := user.Current()
	if err != nil {
		t.Error("Error querying user.Current: " + err.Error() + " (murex will still function)")
	}

	if UserDir(u.Username) == "" {
		t.Error("UserDir returning empty string (murex will still function)")
	}
}
