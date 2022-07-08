//go:build !windows
// +build !windows

package man

import (
	"os"
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
)

// TestMan tests the builtins package
func TestMan(t *testing.T) {
	if os.Getenv("MUREX_TEST_SKIP_MAN") != "" {
		t.Skip("Environmental variable `MUREX_TEST_SKIP_MAN` set")
		return
	}

	count.Tests(t, 3)

	files := GetManPages("cat")
	if len(files) == 0 {
		t.Error("Could not find any man pages")
	}

	flags := ParseFlags(files)
	if len(flags) == 0 {
		t.Errorf("No flags returned for `cat` in: %s", json.LazyLogging(files))
	}

	s := ParseSummary(files)
	if s == "" {
		t.Errorf("No description returned for `cat` in: %s", json.LazyLogging(files))
	}
}
