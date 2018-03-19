package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"os"
	"sort"
	"strings"
)

const heading = "# _murex_ Language Guide\n\n"
const subHeading = "## Command reference"
const goLang = "package docs\n\nfunc init() {\n\tdocs[`%s`] = `%s`\n}"

func compile(dest string, gocode string) {
	for name := range define {
		b := compilePage(name)
		writeDefinitions(dest+"/"+name+".md", b)
		writeGoCode(gocode+"/"+name+".go", name, b)
	}

	writeIndex(dest + "/README.md")
}

func compilePage(funcname string) []byte {
	s := heading
	s += subHeading + ": " + funcname + "\n\n"

	if digest[funcname] != "" {
		s += "> " + digest[funcname] + "\n\n"
	}

	s += define[funcname] + "\n\n"

	if len(related[funcname]) > 0 {
		s += "### See also\n\n"

		sort.Strings(related[funcname])

		for _, rel := range related[funcname] {
			var dig string
			cmd := "`" + rel + "`"
			if digest[rel] != "" {
				dig = ": " + digest[rel]
			}
			if define[rel] != "" {
				cmd = "[`" + rel + "`](" + rel + ".md)"
			}

			s += "* " + cmd + dig + "\n"
		}
	}

	return []byte(s)
}

func writeGoCode(filename string, funcname string, code []byte) {
	if verbose {
		fmt.Println("Writing " + filename)
	}

	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var b []byte
	buf := bytes.NewBuffer(b)
	b64 := base64.NewEncoder(base64.StdEncoding, buf)
	gz := gzip.NewWriter(f)

	i, err := gz.Write(code)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if i != len(code) {
		fmt.Println("Amount written to gzip writer differs from size of markdown doc.")
		os.Exit(1)
	}

	gz.Close()
	b64.Close()
	s := fmt.Sprintf(goLang, funcname, string(b))
	f.WriteString(s)
}

func writeDefinitions(filename string, b []byte) {
	if verbose {
		fmt.Println("Writing " + filename)
	}

	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	i, err := f.Write(b)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if i != len(b) {
		fmt.Println("Amount written to file writer differs from size of markdown doc.")
		os.Exit(1)
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

	out("| Command                   | Description |\n")
	out("| ------------------------- | ----------- |\n")

	for _, name := range definitions {
		//var dig string
		cmd := fmt.Sprintf("%25s", "[`"+name+"`]("+name+".md)")
		if digest[name] != "" {
			//dig = ": " + digest[name]
		}

		//out("* " + cmd + dig + "\n")
		out("| " + cmd + " | " + strings.Replace(digest[name], "\n", " ", -1) + " |\n")
	}
}
