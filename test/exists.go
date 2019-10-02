package test

import (
	"os"
	"strings"
	"testing"

	"github.com/lmorg/murex/test/count"
)

// Exists tests if a file exists
func Exists(t *testing.T, path string) {
	t.Helper()

	if os.Getenv("MUREX_TEST_SKIP_EXISTS") != "" && os.Getenv("MUREX_TEST_SKIP_EXISTS") != "awscodebuild" {
		t.Skip("Environmental variable `MUREX_TEST_SKIP_EXISTS` set")
		return
	}

	count.Tests(t, 1)

	if os.Getenv("MUREX_TEST_SKIP_EXISTS") == "awscodebuild" && strings.HasPrefix(path, "/go:/") {
		path = path[4:]
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Error("Missing file", path)
	}
}
