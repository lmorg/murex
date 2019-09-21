package main

import (
	"os"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/config/profile"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/shell"
)

const (
	interactive    bool = true
	nonInteractive bool = false

	envRunTests = "MUREX_TEST_MAIN_RUN_TESTS"
)

func main() {
	readFlags()

	switch {
	case fRunTests:
		runTests()

	case fCommand != "":
		runCommandLine(fCommand)

	case len(fSource) > 0:
		runSource(fSource[0])

	default:
		startMurex()
	}

	debug.Log("[FIN]")
}

func runTests() error {
	lang.InitEnv()

	defaults.Defaults(lang.ShellProcess.Config, nonInteractive)
	shell.SignalHandler(nonInteractive)

	// compiled profile
	source := defaults.DefaultMurexProfile()
	srcRef := ref.History.AddSource("(builtin)", "source/builtin", []byte(string(source)))
	execSource(defaults.DefaultMurexProfile(), srcRef)

	// enable tests
	if err := lang.ShellProcess.Config.Set("test", "enabled", true); err != nil {
		return err
	}
	if err := lang.ShellProcess.Config.Set("test", "auto-report", false); err != nil {
		return err
	}
	if err := lang.ShellProcess.Config.Set("test", "verbose", true); err != nil {
		return err
	}

	// exit early if being run under Go test
	if os.Getenv(envRunTests) != "" {
		return nil
	}

	// run unit tests
	passed := lang.GlobalUnitTests.Run(lang.ShellProcess, "*")
	lang.ShellProcess.Tests.WriteResults(lang.ShellProcess.Config, lang.ShellProcess.Stdout)

	if !passed {
		os.Exit(1)
	}

	return nil
}

func runCommandLine(commandLine string) {
	lang.InitEnv()

	// default config
	defaults.Defaults(lang.ShellProcess.Config, nonInteractive)
	shell.SignalHandler(nonInteractive)

	// load modules and profile
	if fLoadMods {
		// compiled profile
		source := defaults.DefaultMurexProfile()
		ref := ref.History.AddSource("(builtin)", "source/builtin", []byte(string(source)))
		execSource(defaults.DefaultMurexProfile(), ref)

		// local profile
		profile.Execute()
	}

	// read block from command line parameters
	execSource([]rune(commandLine), nil)
}

func runSource(filename string) {
	lang.InitEnv()

	// default config
	defaults.Defaults(lang.ShellProcess.Config, nonInteractive)
	shell.SignalHandler(nonInteractive)

	// load modules a profile
	if fLoadMods {
		// compiled profile
		source := defaults.DefaultMurexProfile()
		ref := ref.History.AddSource("(builtin)", "source/builtin", []byte(string(source)))
		execSource(defaults.DefaultMurexProfile(), ref)

		// local profile
		profile.Execute()
	}

	// read block from disk
	disk, err := diskSource(filename)
	if err != nil {
		_, err := os.Stderr.WriteString(err.Error() + "\n")
		if err != nil {
			// wouldn't really make any difference at this point because we
			// cannot write to stderr anyway :(
			panic(err)
		}
		os.Exit(1)
	}
	execSource([]rune(string(disk)), nil)
}

func startMurex() {
	lang.InitEnv()

	// default config
	defaults.Defaults(lang.ShellProcess.Config, interactive)

	// compiled profile
	source := defaults.DefaultMurexProfile()
	ref := ref.History.AddSource("(builtin)", "source/builtin", []byte(string(source)))
	execSource(defaults.DefaultMurexProfile(), ref)

	// load modules and profile
	profile.Execute()

	// start interactive shell
	shell.Start()
}
