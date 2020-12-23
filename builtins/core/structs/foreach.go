package structs

import (
	"errors"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.GoFunctions["foreach"] = cmdForEach
}

func cmdForEach(p *lang.Process) (err error) {
	dt := p.Stdin.GetDataType()
	if dt == types.Json {
		p.Stdout.SetDataType(types.JsonLines)
	} else {
		p.Stdout.SetDataType(dt)
	}

	var (
		block   []rune
		varName string
	)

	switch p.Parameters.Len() {
	case 1:
		block, err = p.Parameters.Block(0)
		if err != nil {
			return err
		}

	case 2:
		block, err = p.Parameters.Block(1)
		if err != nil {
			return err
		}

		varName, err = p.Parameters.String(0)
		if err != nil {
			return err
		}

	default:
		return errors.New("Invalid number of parameters")
	}

	err = p.Stdin.ReadArray(func(b []byte) {
		if len(b) == 0 || p.HasCancelled() {
			return
		}

		if varName != "" {
			p.Variables.Set(p, varName, string(b), dt)
		}

		fork := p.Fork(lang.F_PARENT_VARTABLE | lang.F_CREATE_STDIN)
		fork.Stdin.SetDataType(dt)
		fork.Stdin.Writeln(b)
		fork.Execute(block)

	})

	return
}
