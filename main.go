//go:build !js
// +build !js

package main

import (
	"os"
	"time"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/builtins/pipes/term"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/config/profile"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/shell"
	signalhandler "github.com/lmorg/murex/shell/signal_handler"
	"github.com/lmorg/murex/shell/signal_handler/sigfns"
	"github.com/lmorg/murex/utils/cache"
	"github.com/lmorg/murex/utils/readline"
)

const (
	interactive    bool = true
	nonInteractive bool = false
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

	defaults.Config(lang.ShellProcess.Config, nonInteractive)
	signalhandler.EventLoop(nonInteractive)

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
	signalhandler.EventLoop(nonInteractive)

	// compiled profile
	defaultProfile()

	// load modules and profile
	if fLoadMods {
		profile.Execute(profile.F_PRELOAD | profile.F_MODULES | profile.F_PROFILE)
	}

	// read block from command line parameters
	term.OutSetDataTypeIPC()
	sourceRef := ref.Source{
		DateTime: time.Now(),
		Filename: "",
		Module:   "murex/-c",
	}
	execSource([]rune(commandLine), &sourceRef, true)

	if fInteractive {
		shell.Start()
	}
}

func runSource(filename string) {
	lang.InitEnv()

	// default config
	defaults.Config(lang.ShellProcess.Config, nonInteractive)
	signalhandler.EventLoop(nonInteractive)

	// compiled profile
	defaultProfile()

	// load modules a profile
	if fLoadMods {
		profile.Execute(profile.F_PRELOAD | profile.F_MODULES | profile.F_PROFILE)
	}

	// read block from disk
	term.OutSetDataTypeIPC()
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

	sourceRef := ref.Source{
		DateTime: time.Now(),
		Filename: filename,
		Module:   "murex/#!",
	}
	execSource([]rune(string(disk)), &sourceRef, true)
}

func startMurex() {
	lang.InitEnv()

	// default config
	defaults.Config(lang.ShellProcess.Config, interactive)

	cache.SetPath(profile.ModulePath() + "cache.db")
	cache.InitCache()

	// compiled profile
	defaultProfile()

	// load modules and profile
	profile.Execute(profile.F_PRELOAD | profile.F_MODULES | profile.F_PROFILE)

	// start interactive shell
	shell.Start()
}

func registerSignalHandlers() {
	signalhandler.Handlers = &signalhandler.SignalFunctionsT{
		Sigint:  sigfns.Sigint,
		Sigterm: sigfns.Sigterm,
		Sigquit: sigfns.Sigquit,
		Sigtstp: sigfns.Sigtstp,
		Sigchld: sigfns.Sigchld,
	}
	signalhandler.EventLoop(nonInteractive)
}
