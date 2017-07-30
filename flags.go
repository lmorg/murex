package main

import (
	"flag"
	"fmt"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"os"
)

var (
	fCommand string
	fStdin   bool
	fSource  []string
	fEcho    bool
	fHelp1   bool
	fHelp2   bool
	fHelp3   bool
	fSh      bool
)

func readFlags() {
	flag.StringVar(&fCommand, "c", "", "Run command block - read from parameters")
	flag.BoolVar(&fStdin, "stdin", false, "Run command block - read from STDIN")

	flag.BoolVar(&fHelp1, "?", false, "Help")
	flag.BoolVar(&fHelp2, "h", false, "Help")
	flag.BoolVar(&fHelp3, "help", false, "Help")

	flag.BoolVar(&debug.Enable, "debug", false, "Debug")
	flag.BoolVar(&fEcho, "echo", false, "Echo on")
	flag.BoolVar(&fSh, "murex", false, "")

	flag.Parse()

	if fHelp1 || fHelp2 || fHelp3 {
		fmt.Println(config.AppName)
		fmt.Println(config.Version)
		flag.Usage()
		os.Exit(1)
	}

	proc.GlobalConf.Define("shell", "echo", config.Properties{
		Description: "Echo shell functions",
		Default:     fEcho,
		DataType:    types.Boolean,
	})

	fSource = flag.Args()
	proc.ShellProcess.Parameters.Params = os.Args
}
