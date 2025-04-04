package profilepaths

import (
	"os"
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/consts"
)

/*
    envvar:   profile paths
	endpoint: file
*/

func TestValidateProfilePathAsFilExistingFile(t *testing.T) {
	count.Tests(t, 1)

	profilePath := t.TempDir() + "/TestValidateProfilePathAsFilExistingFile"

	path := _validateProfilePath("", ProfileFileName, profilePath, "", false)

	if path != profilePath {
		t.Error("Unexpected output:")
		t.Logf("  Expected: %s", profilePath)
		t.Logf("  Actual:   %s", path)
	}
}

func TestValidateProfilePathAsFileNewFile(t *testing.T) {
	count.Tests(t, 1)

	profilePath := t.TempDir()

	path := _validateProfilePath("", ProfileFileName, profilePath, "", false)

	if path != profilePath+consts.PathSlash+ProfileFileName {
		t.Error("Unexpected output:")
		t.Logf("  Expected: %s", profilePath)
		t.Logf("  Actual:   %s", path)
	}
}

/*
    envvar:   profile paths
	endpoint: directory
*/

func TestValidateProfilePathAsDirExistingDir(t *testing.T) {
	count.Tests(t, 1)

	profilePath := t.TempDir()

	path := _validateProfilePath("", moduleDirName, profilePath, "", true)

	if path != profilePath+consts.PathSlash {
		t.Error("Unexpected output:")
		t.Logf("  Expected: %s", profilePath)
		t.Logf("  Actual:   %s", path)
	}
}

func TestValidateProfilePathAsDirNewDir(t *testing.T) {
	count.Tests(t, 1)

	profilePath := t.TempDir() + "/TestValidateProfilePathAsDirNewDir"

	path := _validateProfilePath("", moduleDirName, profilePath, "", true)

	f, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if !f.IsDir() {
		t.Error("path is not a directory:", path)
	}
}
