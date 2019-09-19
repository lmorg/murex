package test

import (
	"os"
	"testing"

	"github.com/lmorg/murex/test/count"
)

// Exists tests if a file exists
func Exists(t *testing.T, path string) {
	t.Helper()

	if os.Getenv("MUREX_TEST_SKIP_EXISTS") != "" {
		t.Skip("Environmental variable `MUREX_TEST_SKIP_EXISTS` set")
		return
	}

	count.Tests(t, 1, "Exists")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Error("Missing file", path)
	}
}
