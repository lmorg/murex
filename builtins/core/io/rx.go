package io

import (
	"path/filepath"
	"regexp"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.DefineFunction("rx", cmdLsRx, types.Json)
	lang.DefineFunction("!rx", cmdLsRx, types.Json)

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
