package io

import (
	"path/filepath"
	"regexp"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.DefineMethod("rx", cmdLsRx, types.ReadArray, types.Json)
	lang.DefineMethod("!rx", cmdLsRx, types.ReadArray, types.Json)
}

func cmdLsRx(p *lang.Process) (err error) {
	if p.IsMethod {
		return cmdLsRxMethod(p)
	}

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
		if rx.MatchString(files[i]) != p.IsNot {
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

func cmdLsRxMethod(p *lang.Process) (err error) {
	dt := types.Json
	p.Stdout.SetDataType(dt)

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
		if rx.MatchString(files[i]) != !p.IsNot {
			matched = append(matched, files[i])
		}
	}

	aw, err := p.Stdout.WriteArray(dt)
	if err != nil {
		return err
	}

	err = p.Stdin.ReadArray(p.Context, func(b []byte) {
		s := string(b)
		for i := range matched {
			if matched[i] == s {
				return
			}
		}
		err = aw.WriteString(s)
		if err != nil {
			p.Done()
		}
	})
	if err != nil {
		return err
	}

	return aw.Close()
}
