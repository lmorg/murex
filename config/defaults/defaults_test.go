package defaults

import (
	"strings"
	"testing"

	"github.com/lmorg/murex/test/count"
)

// TestDefaultProfileNotEmpty tests the defaults exist
func TestDefaultProfileNotEmpty(t *testing.T) {
	count.Tests(t, 1, "TestDefaultProfileNotEmpty")

	s := string(DefaultMurexProfile())
	if strings.TrimSpace(s) == "" {
		t.Error("Empty default profile")
	}

}
