package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lmorg/murex/debug"
)

var (
	define  map[string]string   = make(map[string]string)
	digest  map[string]string   = make(map[string]string)
	related map[string][]string = make(map[string][]string)
	synonym map[string][]string = make(map[string][]string)
	verbose bool

	murexDocs docs
)

type doc struct {
	Name       string
	IsFunction bool
	Category   string
	Definition string
	Digest     string
	Related    []string
	Synonyms   []string
}

type docs []doc

// Get returns a doc object based on the documentPath (period separated, eg
// "Category.Name"). If no object found doc returns an error
func (d docs) Get(documentPath string) (document doc, err error) {
	return
}

func main() {
	src, dest, gocode := readFlags()

	scanSrc(src)
	debug.Json("Definitions", define)
	debug.Json("Digests", digest)
	debug.Json("Related", related)
	debug.Json("Synonyms", synonym)

	compile(dest, gocode)
}

func readFlags() (src, dest, gocode string) {
	flag.StringVar(&src, "src", "", "Location of definition files")
	flag.StringVar(&dest, "dest", "", "Destination to write docs")
	flag.StringVar(&gocode, "gocode", "", "Destination to Go code")
	flag.BoolVar(&debug.Enable, "debug", false, "Debug messages")
	flag.BoolVar(&verbose, "v", false, "Verbose")

	flag.Parse()

	if src == "" || dest == "" || gocode == "" {
		fmt.Println("Missing required flags. Need -src, -dest and -gocode")
		flag.Usage()
		os.Exit(1)
	}

	return
}
