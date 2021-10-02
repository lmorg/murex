package arraytools

import (
	"regexp"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.DefineMethod("jsplit", cmdJsplit, types.Text, types.Json)
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
	// remove trailing \n
	for i := range split {
	trimString:
		if len(split[i]) > 0 &&
			(split[i][len(split[i])-1] == '\r' || split[i][len(split[i])-1] == '\n') {
			split[i] = utils.CrLfTrimString(split[i])
			goto trimString
		}
	}

trimArray:
	if len(split) > 0 && split[len(split)-1] == "" {
		split = split[:len(split)-1]
		goto trimArray
	}

	json, err := json.Marshal(split, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(json)
	return err
}
