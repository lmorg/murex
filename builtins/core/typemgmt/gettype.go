package typemgmt

import (
	"errors"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineFunction("get-type", cmdGetType, types.String)
}

func cmdGetType(p *lang.Process) error {
	if p.IsMethod {
		return errors.New("this shouldn't be run as a method. Run `murex-docs get-type` for usage")
	}

	v, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	if len(v) == 0 {
		_, err = p.Stdout.Write([]byte(types.Null))
		return err
	}

	var dt string

	switch {
	case v[0] == '$':
		if len(v) == 1 {
			return errors.New("variable data-type requested but with no variable name")
		}
		_, err := p.Variables.GetValue(v[1:])
		if err != nil {
			return err
		}
		dt = p.Variables.GetDataType(v[1:])

	case v == "stdin":
		if (p.Scope.Stdin == nil) {
			return errors.New("stdin must be used within a function or code block. Run `murex-docs stdin` for more information.")
		}
		dt = p.Scope.Stdin.GetDataType()

	default:
		pipe, err := lang.GlobalPipes.Get(v)
		if err != nil {
			return err
		}
		dt = pipe.GetDataType()
	}

	_, err = p.Stdout.Write([]byte(dt))
	return err
}
