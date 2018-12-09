package open

import (
	"errors"

	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	proc.GoFunctions["openagent"] = shell
	proc.GoFunctions["!openagent"] = shell
}

func shell(p *proc.Process) error {
	p.Stdout.SetDataType(types.Generic)

	if p.IsNot {
		dataType, err := p.Parameters.String(0)
		if err != nil {
			return err
		}
		return OpenAgents.Unset(dataType)
	}

	flag, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	dataType, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	switch flag {
	case "get":
		r, err := OpenAgents.Get(dataType)
		if err != nil {
			return err
		}
		_, err = p.Stdout.Write([]byte(string(r)))
		return err

	case "set":
		block, err := p.Parameters.Block(2)
		if err != nil {
			return err
		}
		OpenAgents.Set(dataType, block)
		return nil

	default:
		return errors.New("Invalid option. Please use `get` or `set`; or use the bang (!) to unset")
	}
}
