package consts

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

// TestConsts tests the projects constants package
func TestTempDir(t *testing.T) {
	count.Tests(t, 1)

	if TmpDir() == "" {
		t.Error("No temp directory specified")
	}

	if TmpDir() == _TMP_DIR {
		t.Log("ioutil.TempDir() failed so using fallback path:", _TMP_DIR)
	}
}
