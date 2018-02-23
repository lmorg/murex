package main

import (
	"flag"
	"fmt"
	"github.com/lmorg/murex/debug"
	"os"
)

var (
	Define  map[string]string   = make(map[string]string)
	Digest  map[string]string   = make(map[string]string)
	Related map[string][]string = make(map[string][]string)
	Verbose bool
)

func main() {
	src, dest := readFlags()

	scanSrc(src)
	debug.Json("Definitions", Define)
	debug.Json("Digests", Digest)
	debug.Json("Related", Related)

	compile(dest)
}

func readFlags() (src, dest string) {
	flag.StringVar(&src, "src", "", "Location of definition files")
	flag.StringVar(&dest, "dest", "", "Destination to write docs")
	flag.BoolVar(&debug.Enable, "debug", false, "Debug messages")
	flag.BoolVar(&Verbose, "v", false, "Verbose")

	flag.Parse()

	if src == "" || dest == "" {
		fmt.Println("Missing required flag. Need both -src and -dest")
		flag.Usage()
		os.Exit(1)
	}

	return
}
