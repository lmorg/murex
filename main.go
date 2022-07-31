//go:build !js
// +build !js

package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/config/profile"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/utils/readline"
)

const (
	interactive    bool = true
	nonInteractive bool = false
)

func main() {
	readFlags()

	lang.ProfCpuCleanUp = cpuProfile()
	lang.ProfMemCleanUp = memProfile()

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

func cpuProfile() func() {
	if fCpuProfile != "" {
		fmt.Fprintf(os.Stderr, "Writing CPU profile to '%s'\n", fCpuProfile)

		f, err := os.Create(fCpuProfile)
		if err != nil {
			panic(err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			panic(err)
		}

		return func() {
			pprof.StopCPUProfile()
			if err = f.Close(); err != nil {
				panic(err)
			}

			fmt.Fprintf(os.Stderr, "CPU profile written to '%s'\n", fCpuProfile)
		}
	}

	return func() {}
}

func memProfile() func() {
	if fMemProfile != "" {
		fmt.Fprintf(os.Stderr, "Writing memory profile to '%s'\n", fMemProfile)

		f, err := os.Create(fMemProfile)
		if err != nil {
			panic(err)
		}

		return func() {
			runtime.GC() // get up-to-date statistics
			if err := pprof.WriteHeapProfile(f); err != nil {
				panic(err)
			}
			if err = f.Close(); err != nil {
				panic(err)
			}
			fmt.Fprintf(os.Stderr, "Memory profile written to '%s'\n", fMemProfile)
		}
	}

	return func() {}
}

func runTests() error {
	lang.InitEnv()

	defaults.Config(lang.ShellProcess.Config, nonInteractive)
	shell.SignalHandler(nonInteractive)

	// compiled profile
	defaultProfile()

	// enable tests
	if err := lang.ShellProcess.Config.Set("test", "enabled", true, nil); err != nil {
		return err
	}
	if err := lang.ShellProcess.Config.Set("test", "auto-report", false, nil); err != nil {
		return err
	}
	if err := lang.ShellProcess.Config.Set("test", "verbose", false, nil); err != nil {
		return err
	}
	tty := readline.IsTerminal(int(os.Stdout.Fd()))
	if err := lang.ShellProcess.Config.Set("shell", "color", tty, nil); err != nil {
		return err
	}

	// run unit tests
	passed := lang.GlobalUnitTests.Run(lang.ShellProcess, "*")
	lang.ShellProcess.Tests.WriteResults(lang.ShellProcess.Config, lang.ShellProcess.Stdout)

	if !passed {
		lang.Exit(1)
	}

	return nil
}

func runCommandLine(commandLine string) {
	lang.InitEnv()

	// default config
	defaults.Config(lang.ShellProcess.Config, nonInteractive)
	shell.SignalHandler(nonInteractive)

	// compiled profile
	defaultProfile()

	// load modules and profile
	if fLoadMods {
		profile.Execute(profile.F_PRELOAD | profile.F_MODULES | profile.F_PROFILE)
	}

	// read block from command line parameters
	execSource([]rune(commandLine), nil)
}

func runSource(filename string) {
	lang.InitEnv()

	// default config
	defaults.Config(lang.ShellProcess.Config, nonInteractive)
	shell.SignalHandler(nonInteractive)

	// compiled profile
	defaultProfile()

	// load modules a profile
	if fLoadMods {
		profile.Execute(profile.F_PRELOAD | profile.F_MODULES | profile.F_PROFILE)
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
		lang.Exit(1)
	}
	execSource([]rune(string(disk)), nil)
}

func startMurex() {
	lang.InitEnv()

	// default config
	defaults.Config(lang.ShellProcess.Config, interactive)

	// compiled profile
	defaultProfile()

	// load modules and profile
	profile.Execute(profile.F_PRELOAD | profile.F_MODULES | profile.F_PROFILE)

	// start interactive shell
	shell.Start()
}
