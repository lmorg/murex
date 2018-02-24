package main

import (
	"flag"
	"fmt"
	"github.com/lmorg/murex/debug"
	"os"
)

var (
	define  map[string]string   = make(map[string]string)
	digest  map[string]string   = make(map[string]string)
	related map[string][]string = make(map[string][]string)
	verbose bool
)

func main() {
	src, dest := readFlags()

	scanSrc(src)
	debug.Json("Definitions", define)
	debug.Json("Digests", digest)
	debug.Json("Related", related)

	compile(dest)
}

func readFlags() (src, dest string) {
	flag.StringVar(&src, "src", "", "Location of definition files")
	flag.StringVar(&dest, "dest", "", "Destination to write docs")
	flag.BoolVar(&debug.Enable, "debug", false, "Debug messages")
	flag.BoolVar(&verbose, "v", false, "Verbose")

	flag.Parse()

	if src == "" || dest == "" {
		fmt.Println("Missing required flag. Need both -src and -dest")
		flag.Usage()
		os.Exit(1)
	}

	return
}
