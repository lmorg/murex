package modules

import (
	"testing"

	"github.com/lmorg/murex/utils/which"
)

func TestGitInstalled(t *testing.T) {
	if which.Which("git") == "" {
		t.Error("`git` isn't installed or not in $PATH")
	}
}

func TestGitUriParser(t *testing.T) {
	URIs := []string{
		"git@github.com:/lmorg/murex-module-murex-dev",
		"git@github.com:/lmorg/murex-module-murex-dev.git",
		"https://github.com/lmorg/murex-module-murex-dev",
		"https://github.com/lmorg/murex-module-murex-dev.git",
	}

	expected := "murex-module-murex-dev"

	for _, test := range URIs {
		actual, err := gitGetPath(test)
		if actual != expected || err != nil {
			t.Errorf("gitGetPath() failed to locate clone destination from URI")
			t.Log("  test:    ", test)
			t.Log("  expected:", expected)
			t.Log("  actual:  ", actual)
			t.Log("  error:   ", err)
		}
	}
}
