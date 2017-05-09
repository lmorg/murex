package io

import (
	"encoding/json"
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var rxWhiteSpace *regexp.Regexp = regexp.MustCompile(`[\n\r\t]+`)

func init() {
	proc.GoFunctions["g"] = proc.GoFunction{Func: cmdLsG, TypeIn: types.Null, TypeOut: types.Json}
	proc.GoFunctions["rx"] = proc.GoFunction{Func: cmdLsRx, TypeIn: types.Null, TypeOut: types.Json}
	proc.GoFunctions["ff"] = proc.GoFunction{Func: cmdLsFf, TypeIn: types.Generic, TypeOut: types.Json}
}

func cmdLsG(p *proc.Process) (err error) {
	glob := p.Parameters.AllString()
	files, err := filepath.Glob(glob)

	j, err := json.MarshalIndent(files, "", "\t")
	if err != nil {
		return
	}

	_, err = p.Stdout.Writeln(j)
	return
}

func cmdLsRx(p *proc.Process) (err error) {
	rx, err := regexp.Compile(p.Parameters.AllString())
	if err != nil {
		return
	}

	files, err := filepath.Glob("*")

	var matched []string
	for i := range files {
		if rx.MatchString(files[i]) {
			matched = append(matched, files[i])
		}
	}

	j, err := json.MarshalIndent(files, "", "\t")
	if err != nil {
		return
	}

	_, err = p.Stdout.Writeln(j)
	return
}

func cmdLsFf(p *proc.Process) (err error) {
	var (
		file      bool
		directory bool
		symlink   bool = true
	)

	for _, flag := range p.Parameters {
		switch flag {
		case "+f":
			file = true
		case "+d":
			directory = true
		case "-s":
			symlink = false
		case "-h":
			p.ExitNum = 2
			usage := []byte("Usage:\n  +f   include files\n  +d   include directories\n  -s   exclude symlinks")
			p.Stderr.Writeln(usage)
			return nil
		default:
			return errors.New("Invalid flag. `ft -h` for usage.")
		}
	}

	var files, matched []string
	var isJson bool
	stdin := p.Stdin.ReadAll()

	// Attempt to auto-detect JSON string or string array
	if stdin[0] == '[' {
		if err := json.Unmarshal(stdin, &files); err != nil {
			return err
		}
		isJson = true
	} else {
		files = rxWhiteSpace.Split(string(stdin), -1)
	}

	for i := range files {
		info, err := os.Stat(files[i])
		if err != nil {
			continue
		}

		if file && !info.IsDir() {
			matched = append(matched, files[i])
		}

		if directory && info.IsDir() {
			matched = append(matched, files[i])
		}

		if symlink {
			// TODO: code me
		}
	}

	var b []byte
	if isJson {
		b, err = json.MarshalIndent(matched, "", "\t")
		if err != nil {
			return err
		}
	} else {
		b = []byte(strings.Join(matched, utils.NewLineString))
	}

	_, err = p.Stdout.Writeln(b)
	return
}
