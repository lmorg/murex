package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

var (
	fCommand  string
	fSource   []string
	fLoadMods bool
	fEcho     bool
	fHelp1    bool
	fHelp2    bool
	fHelp3    bool
	fSh       bool
	fRunTests bool
)

func readFlags() {
	flag.StringVar(&fCommand, "c", "", "Run code block - read from parameters")
	flag.BoolVar(&fLoadMods, "load-modules", false, "Load modules and profile when in non-interactive mode ")

	flag.BoolVar(&fHelp1, "?", false, "Help")
	flag.BoolVar(&fHelp2, "h", false, "Help")
	flag.BoolVar(&fHelp3, "help", false, "Help")

	flag.BoolVar(&debug.Enabled, "debug", false, "Debug mode (for debugging murex code. This can also be enabled from inside the shell.")
	flag.BoolVar(&fRunTests, "run-tests", false, "Run all tests and exit")
	flag.BoolVar(&fEcho, "echo", false, "Echo on")
	flag.BoolVar(&fSh, "murex", false, "")

	flag.BoolVar(&lang.FlagTry, "try", false, "Enable a global `try` block")
	flag.BoolVar(&lang.FlagTryPipe, "trypipe", false, "Enable a global `trypipe` block")

	flag.Parse()

	if fHelp1 || fHelp2 || fHelp3 {
		fmt.Println(config.AppName)
		fmt.Println(config.Version)
		flag.Usage()
		os.Exit(1)
	}

	config.InitConf.Define("shell", "echo", config.Properties{
		Description: "Echo shell functions",
		Default:     fEcho,
		DataType:    types.Boolean,
	})

	fSource = flag.Args()
}
