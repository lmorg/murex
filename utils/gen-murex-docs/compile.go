package main

import (
	"fmt"
	"os"
)

const heading = "# _murex_ command reference\n\n"

func compile(dest string) error {
	for name := range Define {
		write(dest+"/"+name+".md", name)
	}

	return nil
}

func write(filename, funcname string) {
	if Verbose {
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
	out("## " + funcname + "\n\n")

	if Digest[funcname] != "" {
		out("> " + Digest[funcname] + "\n\n")
	}

	out(Define[funcname] + "\n\n")

	if len(Related[funcname]) > 0 {
		out("### See also\n\n")

		for _, rel := range Related[funcname] {
			var digest string
			cmd := rel
			if Digest[rel] != "" {
				digest = ": " + Digest[rel]
				cmd = "[" + rel + "](" + rel + ".md)"
			}

			out("* " + cmd + digest + "\n")
		}
	}
}
