package io

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"

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
	lang.DefineMethod("f", cmdLsF, types.ReadArray, types.Json)
	//lang.DefineMethod("!f", cmdLsF, types.ReadArray, types.Json)
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

func cmdLsNotG(p *lang.Process) (err error) {
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

func cmdLsF(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.Json)
	var (
		add, remove fFlagsT
	)

	params := p.Parameters.RuneArray()
	if len(params) == 0 {
		return errors.New("missing parameters")
	}

	for _, r := range params {
		add, remove, err = fFlagsParser(r, add, remove)
		if err != nil {
			return err
		}
	}

	if remove&fHelp != 0 {
		return cmdLsFHelp(p)
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

	for i := range files {
		if p.HasCancelled() {
			break
		}

		info, err := os.Lstat(files[i])
		if err != nil {
			continue
		}

		if matchFlags(add, remove, info) {
			matched = append(matched, files[i])
		}
	}

	var b []byte
	b, err = json.Marshal(matched, p.Stdout.IsTTY())
	if err == nil {
		_, err = p.Stdout.Writeln(b)
	}

	return
}

func cmdLsFHelp(p *lang.Process) error {
	v := make(map[string]string)
	for r, f := range fFlagLookup {
		v[string([]rune{r})] = f.String()
	}

	b, err := json.Marshal(v, p.Stdout.IsTTY())
	if err != nil {
		return err
	}
	_, err = p.Stdout.Write(b)
	return err
}
