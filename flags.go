package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lmorg/murex/app"
	"github.com/lmorg/murex/app/whatsnew"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/session"
)

var (
	fInteractive   bool
	fCommand       string
	fExecute       bool
	fCreateSession bool
	fSource        []string
	fLoadMods      bool
	fEcho          bool
	fHelp1         bool
	fHelp2         bool
	fVersion1      bool
	fVersion2      bool
	fSh            bool
	fRunTests      bool
	fQuiet         bool
)

func readFlags() {
	flag.BoolVar(&fInteractive, "i", false, "Start interactive shell after -c execution")
	flag.StringVar(&fCommand, "c", "", "Run code block (str)")
	flag.BoolVar(&fExecute, "execute", false, "Execute a command from tokenized parameters (argv[])")
	flag.BoolVar(&fCreateSession, "setsid", false, "Set session ID: POSIX compatibility for job control (this will break support for some of Murex's job control features)")
	flag.BoolVar(&fLoadMods, "load-modules", false, "Load modules and profile when in non-interactive mode ")

	flag.BoolVar(&fHelp1, "h", false, "Help")
	flag.BoolVar(&fHelp2, "help", false, "Help")

	flag.BoolVar(&fVersion1, "v", false, "Version")
	flag.BoolVar(&fVersion2, "version", false, "Version")

	flag.BoolVar(&debug.Enabled, "debug", false, "Debug mode (for debugging murex code. This can also be enabled from inside the shell.")
	flag.BoolVar(&fRunTests, "run-tests", false, "Run all tests and exit")
	flag.BoolVar(&fEcho, "echo", false, "Echo on")
	flag.BoolVar(&fSh, "murex", false, "")
	flag.BoolVar(&fQuiet, "quiet", false, "Suppress messages about loading profiles and modules.")
	flag.BoolVar(&whatsnew.Ignore, "ignore-whatsnew", false, "Suppress the what's new message which appears once for new versions (not recommended).")

	flag.BoolVar(&lang.FlagTry, "try", false, "Enable a global `try` block")
	flag.BoolVar(&lang.FlagTryPipe, "trypipe", false, "Enable a global `trypipe` block")
	flag.BoolVar(&lang.FlagTryErr, "tryerr", false, "Enable a global `tryerr` block")
	flag.BoolVar(&lang.FlagTryPipeErr, "trypipeerr", false, "Enable a global `trypipeerr` block")

	flag.Parse()

	if fHelp1 || fHelp2 {
		fmt.Fprintf(os.Stdout, "%s v%s\n", app.Name, app.Version())
		flag.Usage()
		lang.Exit(1)
	}

	if fVersion1 || fVersion2 {
		fmt.Fprintf(os.Stdout, "%s v%s\n", app.Name, app.Version())
		fmt.Fprintf(os.Stdout, "%s\n%s\n", app.License, app.Copyright)
		lang.Exit(0)
	}

	config.InitConf.Define("proc", "echo", config.Properties{
		Description: "Echo shell functions",
		Default:     fEcho,
		DataType:    types.Boolean,
	})

	config.InitConf.Define("shell", "quiet", config.Properties{
		Description: "Prevent messages about loading profiles and modules from being printed at startup (this is set at launch via `--quiet`)",
		Default:     fQuiet,
		DataType:    types.Boolean,
		Global:      true,
	})

	if os.Getenv("MUREX_DEBUG") == "true" {
		debug.Enabled = true
	}

	if fCreateSession || os.Getenv("MUREX_CREATE_SESSION") == "true" {
		session.UnixCreateSession()
	}

	fSource = flag.Args()
}
