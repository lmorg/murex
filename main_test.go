package main

import (
	"os"
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestMainRunTests(t *testing.T) {
	count.Tests(t, 1)

	if err := os.Setenv(envRunTests, "1"); err != nil {
		t.Error(err)
		return
	}

	if err := runTests(); err != nil {
		t.Error(err)
	}
}

func TestRunCommandLine(t *testing.T) {
	count.Tests(t, 1)

	runCommandLine(`out: "testing" -> null`)
}

func TestRunSource(t *testing.T) {
	count.Tests(t, 1)

	file := "test/source.mx"
	runSource(file)
}

func TestRunSourceGz(t *testing.T) {
	count.Tests(t, 1)

	file := "test/source.mx.gz"
	runSource(file)
}
