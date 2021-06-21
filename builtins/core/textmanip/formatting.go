package textmanip

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.GoFunctions["pretty"] = cmdPretty

	defaults.AppendProfile(`
		autocomplete: set pretty { [{
			"Flags": [ "--strict" ]
		}] }
	`)
}

func cmdPretty(p *lang.Process) error {
	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	flags, _, err := p.Parameters.ParseFlags(&parameters.Arguments{
		Flags: map[string]string{
			"--strict": "bool",
		},
	})
	if err != nil {
		return err
	}

	switch {
	case flags["--strict"] == types.TrueString:
		return prettyStrict(p)
	default:
		return prettyDefault(p)
	}
}

func prettyStrict(p *lang.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if dt == types.Json {
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
	}

	_, err := io.Copy(p.Stdout, p.Stdin)
	return err
}

func prettyDefault(p *lang.Process) error {
	p.Stdout.SetDataType(types.Json)

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
}
