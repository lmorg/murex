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
const goLang = "package docs\n\nfunc init() {\nDefinition[`%s`] = `%s`\n}"

func compile(dest string, gocode string) {
	for name := range define {
		b := compilePage(name)
		writeDefinitions(dest+"/"+name+".md", b)
		writeGoCode(gocode+"/autogen-func-"+name+".go", name, b)
	}

	writeGoDigests(gocode + "/autogen-digests.go")
	writeGoSynonyms(gocode + "/autogen-synonyms.go")
	writeIndex(dest + "/README.md")
}

func compilePage(funcname string) []byte {
	s := heading
	s += subHeading + ": " + funcname + "\n\n"

	if digest[funcname] != "" {
		s += "> " + digest[funcname] + "\n\n"
	}

	s += define[funcname] + "\n\n"

	if len(synonym[funcname]) > 0 {
		s += "### Synonyms\n\n"

		for _, syn := range synonym[funcname] {
			s += "* " + syn + "\n"
		}

		s += "\n"
	}

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

	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	//defer gz.Close()

	i, err := gz.Write(code)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if i != len(code) {
		fmt.Println("Amount written to gzip writer differs from size of markdown doc.")
		os.Exit(1)
	}

	err = gz.Flush()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	b64 := base64.StdEncoding.EncodeToString(buf.Bytes())

	s := fmt.Sprintf(goLang, funcname, b64)
	_, err = f.WriteString(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func writeGoDigests(filename string) {
	if verbose {
		fmt.Println("Writing " + filename)
	}

	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = f.WriteString("package docs\n\n// Digest stores a 1 line summary of each builtins\nvar Digest map[string]string = map[string]string{\n")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for name, dig := range digest {
		dig = strings.Replace(dig, "`", "'", -1)
		dig = strings.Replace(dig, "\r", "", -1)
		dig = strings.Replace(dig, "\n", " ", -1)
		_, err := f.WriteString(fmt.Sprintf("\t`%s`: `%s`,\n", name, dig))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	_, err = f.WriteString("}\n")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func writeGoSynonyms(filename string) {
	if verbose {
		fmt.Println("Writing " + filename)
	}

	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = f.WriteString("package docs\n\n//Synonym is used for builtins that might have more than one internal alias\nvar Synonym map[string]string = map[string]string{\n")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for name, syns := range synonym {
		for i := range syns {
			syn := strings.Replace(syns[i], "`", "'", -1)
			_, err := f.WriteString(fmt.Sprintf("\t`%s`: `%s`,\n", syn, name))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}

	for name := range define {
		name := strings.Replace(name, "`", "'", -1)
		_, err := f.WriteString(fmt.Sprintf("\t`%s`: `%s`,\n", name, name))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	_, err = f.WriteString("}\n")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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
