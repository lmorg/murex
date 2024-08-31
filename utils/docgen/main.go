package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	docgen "github.com/lmorg/murex/utils/docgen/api"
)

const (
	// Version is the release ID of docgen
	Version = "4.0.0"

	// Copyright is the copyright owner string
	Copyright = "(c) 2018-2024 Laurence Morgan"

	// License is the projects software license
	License = "License GPL v2"
)

// flags
var (
	fConfigFile string
)

func main() {
	readFlags()

	err := docgen.ReadConfig(fConfigFile)
	if err != nil {
		log.Fatalln("[ERROR]", err)
	}

	err = docgen.Render()
	if err != nil {
		log.Fatalln("[ERROR]", err)
	}

	os.Exit(docgen.ExitStatus)
}

func readFlags() error {
	flag.BoolVar(&docgen.Panic, "panic", false, "Write a stack trace on error")
	flag.BoolVar(&docgen.Verbose, "verbose", false, "Verbose output (all log messages inc warnings)")
	flag.BoolVar(&docgen.Warning, "warning", false, "Display warning messages (recommended)")
	flag.BoolVar(&docgen.ReadOnly, "readonly", false, "Don't write output to disk. Use this to test the config")

	flag.StringVar(&fConfigFile, "config", "", "Location of the base docgen config file")
	version := flag.Bool("version", false, "Output docgen version number and exit")

	flag.Parse()

	if *version {
		fmt.Printf("docgen version %s\n%s, %s", Version, License, Copyright)
		os.Exit(0)
	}

	if fConfigFile == "" {
		log.Fatalln("missing required flag: -config")
	}

	return nil
}
