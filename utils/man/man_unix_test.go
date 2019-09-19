// +build !windows

package man

import (
	"os"
	"testing"

	"github.com/lmorg/murex/test/count"
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
		t.Error("No flags returned for `cat`")
	}

	s := ParseSummary(files)
	if s == "" {
		t.Error("No description returned for `cat`")
	}
}
