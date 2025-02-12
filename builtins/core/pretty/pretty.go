package pretty

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/lmorg/murex/builtins/types/xml"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	//lang.GoFunctions["pretty"] = cmdPretty
	lang.DefineMethod("pretty", cmdPretty, types.Json, types.Json)

	defaults.AppendProfile(fmt.Sprintf(`
		autocomplete set pretty %%[{
			Flags: [ %[1]s %[2]s ]
			FlagValues: {
				%[2]s: [{ Flags: [ json xml ] }]
			}
		}]
	`, fStrict, fDataType))
}

const (
	fStrict   = "--strict"
	fDataType = "--type"
)

func cmdPretty(p *lang.Process) error {
	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	flags, _, err := p.Parameters.ParseFlags(&parameters.Arguments{
		Flags: map[string]string{
			fStrict:   types.Boolean,
			fDataType: types.String,
		},
	})
	if err != nil {
		return err
	}

	switch {
	case flags[fDataType] != "":
		return prettyType(p, flags[fDataType], flags[fStrict] == types.TrueString)

	default:
		return prettyType(p, p.Stdin.GetDataType(), flags[fStrict] == types.TrueString)
	}
}

func prettyType(p *lang.Process, dt string, strict bool) error {
	p.Stdout.SetDataType(dt)

	switch dt {
	case types.Json:
		b, err := p.Stdin.ReadAll()
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		err = json.Indent(&prettyJSON, b, "", "    ")
		if err != nil {
			return err
		}

		_, err = p.Stdout.Write(prettyJSON.Bytes())
		return err

	case "xml":
		v, err := xml.UnmarshalFromProcess(p)
		if err != nil {
			return err
		}
		b, err := xml.MarshalTTY(v, true)
		if err != nil {
			return err
		}
		_, err = p.Stdout.Write(b)
		return err

	case types.Generic:
		if !strict {
			var err error
			err = prettyType(p, types.Json, false)
			if err == nil {
				return nil
			}
			err = prettyType(p, "xml", false)
			if err == nil {
				return nil
			}
		}

		fallthrough

	default:
		_, err := io.Copy(p.Stdout, p.Stdin)
		return err
	}
}
