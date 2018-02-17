package management

import (
	"errors"

	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

func init() {
	proc.GoFunctions["config"] = cmdConfig
}

func cmdConfig(p *proc.Process) error {
	if p.Parameters.Len() == 0 {
		p.Stdout.SetDataType(types.Json)

		b, err := utils.JsonMarshal(p.Config.Dump(), p.Stdout.IsTTY())
		if err != nil {
			return err
		}

		_, err = p.Stdout.Writeln(b)
		return err
	}

	option, _ := p.Parameters.String(0)
	switch option {
	case "get":
		app, _ := p.Parameters.String(1)
		key, _ := p.Parameters.String(2)
		val, err := p.Config.Get(app, key, types.String)
		if err != nil {
			return err
		}
		p.Stdout.SetDataType(p.Config.DataType(app, key))
		p.Stdout.Writeln([]byte(val.(string)))

	case "set":
		p.Stdout.SetDataType(types.Null)
		app, _ := p.Parameters.String(1)
		key, _ := p.Parameters.String(2)
		var val string

		if p.IsMethod == true {
			b, err := p.Stdin.ReadAll()
			if err != nil {
				return err
			}
			val = string(b)

		} else {
			val, _ = p.Parameters.String(3)
		}

		err := p.Config.Set(app, key, val)
		return err

	case "define":
	default:
		p.Stdout.SetDataType(types.Null)
		return errors.New("Unknown option. Please get, set or define.")
	}

	return nil
}
