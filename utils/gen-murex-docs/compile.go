package main

import (
	"fmt"
	"os"
	"sort"
)

const heading = "# _murex_ Language Guide\n\n"
const subHeading = "## Command reference"

func compile(dest string) {
	for name := range define {
		writeDefinitions(dest+"/"+name+".md", name)
	}

	writeIndex(dest + "/README.md")
}

func writeDefinitions(filename, funcname string) {
	if verbose {
		fmt.Println("Writing " + filename)
	}

	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	out := func(s string) {
		_, err := f.WriteString(s)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	out(heading)
	out(subHeading + ": " + funcname + "\n\n")

	if digest[funcname] != "" {
		out("> " + digest[funcname] + "\n\n")
	}

	out(define[funcname] + "\n\n")

	if len(related[funcname]) > 0 {
		out("### See also\n\n")

		sort.Strings(related[funcname])

		for _, rel := range related[funcname] {
			var dig string
			cmd := rel
			if digest[rel] != "" {
				dig = ": " + digest[rel]
			}
			if define[funcname] != "" {
				cmd = "[" + rel + "](" + rel + ".md)"
			}

			out("* " + cmd + dig + "\n")
		}
	}
}

func writeIndex(filename string) {
	if verbose {
		fmt.Println("Writing " + filename)
	}

	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	out := func(s string) {
		_, err := f.WriteString(s)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	out(heading)
	out(subHeading + "\n\n")

	var definitions []string
	for name := range define {
		definitions = append(definitions, name)
	}
	sort.Strings(definitions)

	for _, name := range definitions {
		var dig string
		cmd := "[" + name + "](" + name + ".md)"
		if digest[name] != "" {
			dig = ": " + digest[name]
		}

		out("* " + cmd + dig + "\n")
	}
}
