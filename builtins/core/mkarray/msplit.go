package mkarray

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"regexp"
)

func init() {
	proc.GoFunctions["jsplit"] = cmdJsplit
}

func cmdJsplit(p *proc.Process) error {
	p.Stdout.SetDataType(types.Json)

	b, err := p.Stdin.ReadAll()
	if err != nil {
		return err
	}

	pattern := p.Parameters.StringAll()
	rx, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	split := rx.Split(string(b), -1)
	json, err := utils.JsonMarshal(split, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(json)
	return err
}
