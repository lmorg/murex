//go:build !js
// +build !js

//go:generate go get golang.org/x/tools/cmd/stringer
//go:generate go build -v golang.org/x/tools/cmd/stringer

package main

import (
	"flag"
	"os"
	"strings"
	"time"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/builtins/pipes/term"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/config/profile"
	profilepaths "github.com/lmorg/murex/config/profile/paths"
	"github.com/lmorg/murex/config/profile/source"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/utils/cache"
	"github.com/lmorg/murex/utils/escape"
	"github.com/lmorg/readline/v4"
)

func main() {
	readFlags()

	switch {
	case fRunTests:
		runTests()

	case fCommand != "":
		runCommandString(fCommand)

	case fExecute:
		runCommandString(argvToCmdLineStr(flag.Args()))

	case len(fSource) > 0:
		runSource(fSource[0])

	default:
		startMurexRepl()
	}

	debug.Log("[FIN]")
}

func runTests() error {
	lang.InitEnv()

	defaults.Config(lang.ShellProcess.Config, fInteractive)
	registerSignalHandlers(fInteractive)

	// compiled profile
	profile.Execute(profile.F_BUILTIN)

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

func argvToCmdLineStr(argv []string) string {
	cmdLine := make([]string, len(argv))
	copy(cmdLine, argv)
	escape.CommandLine(cmdLine)
	return strings.Join(cmdLine, " ")
}

func runCommandString(commandString string) {
	lang.InitEnv()

	// default config
	defaults.Config(lang.ShellProcess.Config, fInteractive)
	registerSignalHandlers(fInteractive)

	// compiled profile
	profiles := profile.F_BUILTIN

	// load modules and profile
	if fLoadMods {
		profiles |= profile.F_PRELOAD | profile.F_MOD_PRELOAD | profile.F_MODULES | profile.F_PROFILE
	}

	profile.Execute(profiles)

	// read block from command line parameters
	term.OutSetDataTypeIPC()
	sourceRef := ref.Source{
		DateTime: time.Now(),
		Filename: "",
		Module:   "murex/-c",
	}
	source.Exec([]rune(commandString), &sourceRef, true)

	if fInteractive {
		shell.Start()
	}
}

func runSource(filename string) {
	lang.InitEnv()

	// default config
	defaults.Config(lang.ShellProcess.Config, fInteractive)
	registerSignalHandlers(fInteractive)

	// compiled profile
	profiles := profile.F_BUILTIN

	// load modules and profile
	if fLoadMods {
		profiles |= profile.F_PRELOAD | profile.F_MOD_PRELOAD | profile.F_MODULES | profile.F_PROFILE
	}

	profile.Execute(profiles)

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
	source.Exec([]rune(string(disk)), &sourceRef, true)
}

func startMurexRepl() {
	lang.InitEnv()

	// default config
	defaults.Config(lang.ShellProcess.Config, true)

	cache.SetPath(profilepaths.ModulePath() + "cache.db")
	cache.InitCache()

	// compiled profile
	profiles := profile.F_BUILTIN | profile.F_PRELOAD | profile.F_MOD_PRELOAD | profile.F_MODULES | profile.F_PROFILE

	profile.Execute(profiles)

	// start interactive shell
	registerSignalHandlers(true)
	shell.Start()
}
