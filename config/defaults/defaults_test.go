package defaults

import (
	"strings"
	"testing"
)

// TestDefaultProfileNotEmpty tests the defaults exist
func TestDefaultProfileNotEmpty(t *testing.T) {
	s := string(DefaultMurexProfile())
	if strings.TrimSpace(s) == "" {
		t.Error("Empty default profile")
	}

}
