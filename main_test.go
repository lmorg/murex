//go:build !js
// +build !js

package main

import (
	"testing"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
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

	runCommandString(`out: "testing" -> null`)
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

func TestArgvToCmdLineStr(t *testing.T) {
	tests := map[string][]string{
		`foobar`:       []string{`foobar`},
		`\ foo`:        []string{` foo`},
		`foo\ `:        []string{`foo `}, // TODO: this might cause things to break...?
		`foo bar`:      []string{`foo`, `bar`},
		`foo\ bar baz`: []string{`foo bar`, `baz`},
		`foo\  bar`:    []string{`foo `, `bar`},
	}

	count.Tests(t, len(tests))

	for expected, argv := range tests {
		actual := argvToCmdLineStr(argv)

		if expected != actual {
			t.Error("Expected does not match actual:")
			t.Logf("  argv:     %s", json.LazyLogging(argv))
			t.Logf("  expected: '%s'", expected)
			t.Logf("  actual:   '%s'", actual)
		}
	}
}
