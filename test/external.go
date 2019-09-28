package test

import (
	"os"
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/which"
)

// InstalledDepsTest checks any external dependencies are installed
func InstalledDepsTest(t *testing.T, exec string) bool {
	t.Helper()

	if os.Getenv("MUREX_TEST_NO_EXEC_DEPS") != "" {
		return false
	}

	count.Tests(t, 1)

	if which.Which(exec) == "" {
		t.Errorf("`%s` isn't installed or not in $PATH", exec)
		return false
	}

	return true
}
