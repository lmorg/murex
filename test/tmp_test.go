package test_test

import (
	"os"
	"testing"

	"github.com/lmorg/murex/test"
	"github.com/lmorg/murex/test/count"
)

func TestTempDir(t *testing.T) {
	count.Tests(t, 1)

	path, err := test.TempDir()
	if err != nil {
		t.Fatalf("Error getting the testing temporary directory: %s", err.Error())
		return
	}

	fi, err := os.Stat(path)
	if err != nil {
		t.Fatalf("Error stating the testing temporary directory: %s", err.Error())
		return
	}

	if !fi.IsDir() {
		t.Fatalf("The testing temporary directory is not a directory: %s", path)
		return
	}
}
