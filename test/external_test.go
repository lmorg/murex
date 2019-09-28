package test

import (
	"runtime"
	"testing"
)

func TestInstalledDepsTest(t *testing.T) {
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
