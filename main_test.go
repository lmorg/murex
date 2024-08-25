//go:build !js
// +build !js

package main

import (
	"testing"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test/count"
)

// TestMurex tests murex runtime environment can be initialized and and simple
// command line can execute
func TestMurex(t *testing.T) {
	count.Tests(t, 1)

	lang.InitEnv()

	block := []rune("a [Mon..Fri]->regexp m/^T/")

	_, err := lang.ShellProcess.Fork(lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_NO_STDERR).Execute(block)
	if err != nil {
		t.Error(err.Error())
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

func TestRunSourceGzMods(t *testing.T) {
	count.Tests(t, 1)

	file := "test/source.mx.gz"
	runSource(file)
}
