package main

import (
	"flag"
	"fmt"
	golog "log"
	"os"
)

// Version is the release ID of docgen
const Version = "1.1.10"

// flags
var (
	fConfigFile string
	fVerbose    bool
	fDebug      bool
)

func main() {
	defer func() {
		// Write a stack trace on error
		if !fDebug {
			if r := recover(); r != nil {
				golog.Fatalln("[ERROR]", r)
			}
		}
	}()

	readFlags()
	readConfig(fConfigFile)
	walkSourcePath(Config.SourcePath)
	renderAll(Documents)
}

func readFlags() error {
	flag.StringVar(&fConfigFile, "config", "", "Location of the base docgen config file")
	flag.BoolVar(&fDebug, "debug", false, "Write a stack trace on error")
	flag.BoolVar(&fVerbose, "verbose", false, "Verbose")
	version := flag.Bool("version", false, "Output docgen version number and exit")

	flag.Parse()

	if *version {
		fmt.Printf("docgen version %s\nLicence GPL v2, (C) 2018 Laurence Morgan", Version)
		os.Exit(0)
	}

	if fConfigFile == "" {
		panic("missing required flag: -config")
	}

	return nil
}

func log(v ...interface{}) {
	if fVerbose {
		golog.Println(append([]interface{}{"[LOG]"}, v...)...)
	}
}

func warning(v ...interface{}) {
	golog.Println(append([]interface{}{"[WARNING]"}, v...)...)
}
