package consts

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

// TestConsts tests the projects constants package
func TestTempDir(t *testing.T) {
	count.Tests(t, 1, "TestTempDir")

	if TempDir == "" {
		t.Error("No temp directory specified")
	}

	if TempDir == tempDir {
		t.Log("ioutil.TempDir() failed so using fallback path:", tempDir)
	}
}
