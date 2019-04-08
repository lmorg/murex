package consts

import (
	"testing"
)

// TestConsts tests the projects constants package
func TestTempDir(t *testing.T) {
	if TempDir == "" {
		t.Error("No temp directory specified")
	}

	if TempDir != tempDir {
		t.Log("ioutil.TempDir() failed so using fallback path:", tempDir)
	}
}
