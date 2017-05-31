package main

import (
	"flag"
	"fmt"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/debug"
	"os"
)

var (
	fCommand string
	fStdin   bool

	fSource []string

	fHelp1 bool
	fHelp2 bool
	fHelp3 bool
)

func readFlags() {
	flag.StringVar(&fCommand, "c", "", "Run command block - read from parameters")
	flag.BoolVar(&fStdin, "stdin", false, "Run command block - read from STDIN")

	flag.BoolVar(&fHelp1, "?", false, "Help")
	flag.BoolVar(&fHelp2, "h", false, "Help")
	flag.BoolVar(&fHelp3, "help", false, "Help")

	flag.BoolVar(&debug.Enable, "debug", false, "Debug")
	flag.BoolVar(&debug.EchoOn, "echo", false, "Echo on")

	flag.Parse()

	if fHelp1 || fHelp2 || fHelp3 {
		fmt.Println(config.AppName)
		fmt.Println(config.Version)
		flag.Usage()
		os.Exit(1)
	}

	fSource = flag.Args()
}
