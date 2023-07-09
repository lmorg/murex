//go:build !windows
// +build !windows

package man

import (
	"os"
	"strings"
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
	if len(files) == 0 || strings.Contains(files[0], "'unminimize'") {
		t.Log("Could not find any man pages so reverting to local copy")

		gopath := os.Getenv("GOPATH")
		if gopath == "" {
			t.Errorf("env var GOPATH is not set")
		}
		files = []string{gopath + "/src/github.com/lmorg/murex/test/cat.1.gz"}
	}

	/*lang.InitEnv()
	flags, _ := ParseByPaths("cat", files)
	if len(flags) == 0 {
		t.Fatalf("No flags returned for `cat` in: %s", json.LazyLogging(files))
	}

	if strings.HasPrefix(flags[0], errPrefix) {
		t.Fatalf(flags[0])
	}*/

	s := ParseSummary(files)
	if s == "" {
		t.Errorf("No summary returned for `cat` in: %s", json.LazyLogging(files))
	}
}
