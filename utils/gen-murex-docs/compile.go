package main

import (
	"fmt"
	"os"
)

const heading = "# _murex_ reference documents\n\n"

func compile(dest string) error {
	for name := range define {
		write(dest+"/"+name+".md", name)
	}

	return nil
}

func write(filename, funcname string) {
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
	out("## Builtin function: " + funcname + "\n\n")

	if digest[funcname] != "" {
		out("> " + digest[funcname] + "\n\n")
	}

	out(define[funcname] + "\n\n")

	if len(related[funcname]) > 0 {
		out("### See also\n\n")

		for _, rel := range related[funcname] {
			var dig string
			cmd := rel
			if digest[rel] != "" {
				dig = ": " + digest[rel]
				cmd = "[" + rel + "](" + rel + ".md)"
			}

			out("* " + cmd + dig + "\n")
		}
	}
}
