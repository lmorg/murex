package typemgmt

import (
	"errors"
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.GoFunctions["get-type"] = cmdGetType
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
		if p.Variables.GetValue(v[1:]) == nil {
			return fmt.Errorf("no variable set with the name `%s`", v[1:])
		}
		dt = p.Variables.GetDataType(v[1:])

	case v == "stdin":
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
