package test

import (
	"os"
	"runtime"
	"testing"
)

func TestInstalledDepsTest(t *testing.T) {
	if os.Getenv("MUREX_TEST_NO_EXEC_DEPS") != "" {
		return
	}

	var exec string

	switch runtime.GOOS {
	case "windows":
		exec = "cmd.exe"

	case "plan9":
		exec = "rc"

	default:
		exec = "sh"
	}

	if !InstalledDepsTest(t, exec) {
		t.Errorf("TestInstalledDepsTest failed")
	}
}
