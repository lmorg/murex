package test

import (
	"os"
	"testing"
)

// Exists tests if a file exists
func Exists(t *testing.T, path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Error("Missing file", path)
	}
}
