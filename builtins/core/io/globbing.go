package io

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.GoFunctions["g"] = cmdLsG
	lang.GoFunctions["rx"] = cmdLsRx
	lang.GoFunctions["f"] = cmdLsF
}

func cmdLsG(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.Json)
	glob := p.Parameters.StringAll()

	files, err := filepath.Glob(glob)
	if err != nil {
		return
	}

	j, err := json.Marshal(files, p.Stdout.IsTTY())
	if err != nil {
		return
	}

	_, err = p.Stdout.Writeln(j)
	return
}

func cmdLsRx(p *lang.Process) (err error) {
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

	j, err := json.Marshal(matched, p.Stdout.IsTTY())
	if err != nil {
		return
	}

	_, err = p.Stdout.Writeln(j)
	return
}

func cmdLsF(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.Json)
	var (
		file      bool
		directory bool
		symlink   = true
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
			return errors.New("Invalid flag. `f -h` for usage")
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

	if debug.Enabled {
		for _, f := range files {
			debug.Log("f->", f)
		}
	}

	for i := range files {
		if p.HasCancelled() {
			break
		}

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
	b, err = json.Marshal(matched, p.Stdout.IsTTY())
	if err == nil {
		_, err = p.Stdout.Writeln(b)
	}

	return
}
