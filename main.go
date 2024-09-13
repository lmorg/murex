//go:build !js
// +build !js

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
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/utils/cache"
	"github.com/lmorg/murex/utils/escape"
	"github.com/lmorg/murex/utils/readline"
)

func main() {
	readFlags()

	switch {
	case fRunTests:
		runTests()

	case fCommand != "":
		runCommandString(fCommand)

	case fExecute:
		//executeAs(flag.Args())
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
	execSource([]rune(commandString), &sourceRef, true)

	if fInteractive {
		shell.Start()
	}
}

/*func executeAs(argv []string) {
	lang.InitEnv()

	// default config
	defaults.Config(lang.ShellProcess.Config, fInteractive)
	registerSignalHandlers(fInteractive)

	// compiled profile
	defaultProfile()

	// load modules and profile
	if fLoadMods {
		profile.Execute(profile.F_PRELOAD | profile.F_MODULES | profile.F_PROFILE)
	}

	// read block from command line parameters
	term.OutSetDataTypeIPC()

	err := lang.ShellProcess.Config.Set("proc", "force-tty", true, lang.ShellProcess.FileRef)
	if err != nil {
		panic(err)
	}

	lang.ShellProcess.Name.Set(argv[0])
	lang.ShellProcess.Parameters.DefineParsed(argv[1:])

	err = lang.External(lang.ShellProcess)
	if err != nil {
		_, err = os.Stdout.WriteString(err.Error())
		if err != nil {
			panic(err)
		}
	}

	lang.Exit(lang.ShellProcess.ExitNum)
}*/

func runSource(filename string) {
	lang.InitEnv()

	// default config
	defaults.Config(lang.ShellProcess.Config, fInteractive)
	registerSignalHandlers(fInteractive)

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

func startMurexRepl() {
	lang.InitEnv()

	// default config
	defaults.Config(lang.ShellProcess.Config, true)

	cache.SetPath(profile.ModulePath() + "cache.db")
	cache.InitCache()

	// compiled profile
	defaultProfile()

	// load modules and profile
	profile.Execute(profile.F_PRELOAD | profile.F_MODULES | profile.F_PROFILE)

	// start interactive shell
	registerSignalHandlers(true)
	shell.Start()
}
