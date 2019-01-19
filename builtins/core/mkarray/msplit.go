package mkarray

import (
	"regexp"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.GoFunctions["jsplit"] = cmdJsplit
}

func cmdJsplit(p *lang.Process) error {
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
	json, err := json.Marshal(split, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(json)
	return err
}
