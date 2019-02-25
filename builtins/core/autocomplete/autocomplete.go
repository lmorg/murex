package cmdautocomplete

import (
	"errors"
	"sort"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.GoFunctions["autocomplete"] = cmdAutocomplete
}

func cmdAutocomplete(p *lang.Process) error {
	mode, err := p.Parameters.String(0)
	if err != nil {
		p.Stdout.SetDataType(types.Null)
		return err
	}

	switch mode {
	case "get":
		return get(p)

	case "set":
		return set(p)

	default:
		p.Stdout.SetDataType(types.Null)
		return errors.New("Not a valid mode. Please use `get` or `set`")
	}
}

func get(p *lang.Process) error {
	p.Stdout.SetDataType(types.Json)

	b, err := json.Marshal(autocomplete.ExesFlags, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Writeln(b)
	return err
}

func set(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	exe, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	var jf []byte

	if p.IsMethod {

		jf, err = p.Stdin.ReadAll()
		if err != nil {
			return err
		}

	} else {

		jfr, err := p.Parameters.Block(2)
		if err == nil {
			jf = []byte(string(jfr))
		} else {
			jf, err = p.Parameters.Byte(2)
			if err != nil {
				return err
			}
		}
	}

	var flags []autocomplete.Flags
	err = json.UnmarshalMurex(jf, &flags)
	if err != nil {
		return err
	}

	for i := range flags {
		// So we don't have nil values in JSON
		if len(flags[i].Flags) == 0 {
			flags[i].Flags = make([]string, 0)
		}

		sort.Strings(flags[i].Flags)
	}

	autocomplete.ExesFlags[exe] = flags
	autocomplete.ExesFlagsMod[exe] = p.Module
	return nil
}
