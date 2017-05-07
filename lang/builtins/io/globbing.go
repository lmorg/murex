package io

import (
	"encoding/json"
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"os"
	"path/filepath"
	"regexp"
)

func init() {
	proc.GoFunctions["g"] = proc.GoFunction{Func: cmdLsG, TypeIn: types.Null, TypeOut: types.Json}
	proc.GoFunctions["rx"] = proc.GoFunction{Func: cmdLsRx, TypeIn: types.Null, TypeOut: types.Json}
	proc.GoFunctions["ff"] = proc.GoFunction{Func: cmdLsFf, TypeIn: types.Json, TypeOut: types.Json}
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

func cmdLsFf(p *proc.Process) error {
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
	if err := json.Unmarshal(p.Stdin.ReadAll(), &files); err != nil {
		return err
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
	return nil
}
