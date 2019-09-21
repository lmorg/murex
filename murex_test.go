package main

import (
	"os"
	"sync"
	"testing"

	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/test/count"
)

// TestMurex tests murex runtime environment can be initialised and and simple
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

var testUnitLock sync.Mutex

func TestUnitTestMurexSourceFiles(t *testing.T) {
	testUnitLock.Lock()
	if os.Getenv("MUREX_TEST_UNIT_TESTS") != "" {
		testUnitLock.Unlock()
		return
	}

	err := os.Setenv("MUREX_TEST_UNIT_TESTS", "1")
	testUnitLock.Unlock()
	if err != nil {
		panic(err)
	}

	lang.InitEnv()

	defaults.Defaults(lang.ShellProcess.Config, nonInteractive)
	shell.SignalHandler(nonInteractive)

	// compiled profile
	source := defaults.DefaultMurexProfile()
	srcRef := ref.History.AddSource("(builtin)", "source/builtin", []byte(string(source)))
	execSource(defaults.DefaultMurexProfile(), srcRef)

	// enable tests
	if err := lang.ShellProcess.Config.Set("test", "enabled", false); err != nil {
		t.Fatal(err)
	}
	if err := lang.ShellProcess.Config.Set("test", "auto-report", false); err != nil {
		t.Fatal(err)
	}
	if err := lang.ShellProcess.Config.Set("test", "verbose", false); err != nil {
		t.Fatal(err)
	}
	if err := lang.ShellProcess.Config.Set("shell", "color", false); err != nil {
		t.Fatal(err)
	}

	// run unit tests
	passed := lang.GlobalUnitTests.Run(lang.ShellProcess, "*")
	count.Tests(t, lang.ShellProcess.Tests.Results.Len())

	if passed {
		return
	}

	t.Error("One or more murex unit tests failed ^")
	err = lang.ShellProcess.Tests.WriteResults(lang.ShellProcess.Config, lang.ShellProcess.Stdout)
	if err != nil {
		t.Error(err)
	}
}
