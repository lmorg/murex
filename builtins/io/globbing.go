package io

import (
	"errors"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"os"
	"path/filepath"
	"regexp"
)

func init() {
	proc.GoFunctions["g"] = proc.GoFunction{Func: cmdLsG, TypeIn: types.Null, TypeOut: types.Json}
	proc.GoFunctions["rx"] = proc.GoFunction{Func: cmdLsRx, TypeIn: types.Null, TypeOut: types.Json}
	proc.GoFunctions["f"] = proc.GoFunction{Func: cmdLsF, TypeIn: types.Generic, TypeOut: types.Json}
}

func cmdLsG(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Json)
	glob := p.Parameters.StringAll()

	files, err := filepath.Glob(glob)
	if err != nil {
		return
	}

	j, err := utils.JsonMarshal(files, p.Stdout.IsTTY())
	if err != nil {
		return
	}

	_, err = p.Stdout.Writeln(j)
	return
}

func cmdLsRx(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Json)
	rx, err := regexp.Compile(p.Parameters.StringAll())
	if err != nil {
		return
	}

	files, err := filepath.Glob("*")
	if err != nil {
		return
	}

	var matched []string
	for i := range files {
		if rx.MatchString(files[i]) {
			matched = append(matched, files[i])
		}
	}

	j, err := utils.JsonMarshal(matched, p.Stdout.IsTTY())
	if err != nil {
		return
	}

	_, err = p.Stdout.Writeln(j)
	return
}

func cmdLsF(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Json)
	var (
		file      bool
		directory bool
		symlink   bool = true
	)

	for _, flag := range p.Parameters.StringArray() {
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
			return errors.New("Invalid flag. `f -h` for usage.")
		}
	}

	var files, matched []string

	if p.IsMethod {
		p.Stdin.ReadArray(func(b []byte) {
			files = append(files, string(b))
		})

	} else {
		files, err = filepath.Glob("*")
		if err != nil {
			return
		}
	}

	for _, f := range files {
		debug.Log("f->", f)
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
	b, err = utils.JsonMarshal(matched, p.Stdout.IsTTY())
	if err == nil {
		_, err = p.Stdout.Writeln(b)
	}

	return
}
