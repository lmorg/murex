package io

import (
	"path/filepath"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/lists"
)

func init() {
	lang.DefineFunction("g", cmdLsG, types.Json)
	lang.DefineFunction("!g", cmdLsNotG, types.Json)
	lang.DefineFunction("rx", cmdLsRx, types.Json)
	lang.DefineFunction("!rx", cmdLsRx, types.Json)

}

func cmdLsG(p *lang.Process) (err error) {
	if p.IsMethod {
		return cmdLsGMethod(p)
	}

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

func cmdLsGMethod(p *lang.Process) (err error) {
	dt := types.Json
	p.Stdout.SetDataType(dt)

	glob := p.Parameters.StringAll()
	all, err := filepath.Glob(glob)
	if err != nil {
		return
	}

	aw, err := p.Stdout.WriteArray(dt)
	if err != nil {
		return err
	}

	err = p.Stdin.ReadArray(func(b []byte) {
		s := string(b)
		for i := range all {
			if all[i] == s {
				err = aw.WriteString(s)
				if err != nil {
					p.Done()
				}
			}
		}
	})
	if err != nil {
		return err
	}

	return aw.Close()
}

func cmdLsNotG(p *lang.Process) (err error) {
	if p.IsMethod {
		return cmdLsNotGMethod(p)
	}

	p.Stdout.SetDataType(types.Json)

	glob, err := filepath.Glob(p.Parameters.StringAll())
	if err != nil {
		return
	}

	all, err := filepath.Glob("*")
	if err != nil {
		return
	}

	var files []string
	for i := range all {
		if !lists.Match(glob, all[i]) {
			files = append(files, all[i])
		}
	}

	j, err := json.Marshal(files, p.Stdout.IsTTY())
	if err != nil {
		return
	}

	_, err = p.Stdout.Writeln(j)
	return
}

func cmdLsNotGMethod(p *lang.Process) (err error) {
	dt := types.Json
	p.Stdout.SetDataType(dt)

	glob := p.Parameters.StringAll()
	all, err := filepath.Glob(glob)
	if err != nil {
		return
	}

	aw, err := p.Stdout.WriteArray(dt)
	if err != nil {
		return err
	}

	err = p.Stdin.ReadArray(func(b []byte) {
		s := string(b)
		for i := range all {
			if all[i] == s {
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
