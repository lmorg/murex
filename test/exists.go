package test

import (
	"os"
	"testing"
)

// Exists tests if a file exists
func Exists(t *testing.T, path string) {
	if os.Getenv("MUREX_TEST_SKIP_EXISTS") != "" {
		t.Skip("Environmental variable `MUREX_TEST_SKIP_EXISTS` set")
		return
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Error("Missing file", path)
	}
}
