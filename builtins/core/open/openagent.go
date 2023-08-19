package open

import (
	"errors"
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineFunction("openagent", shell, types.Generic)
	lang.DefineFunction("!openagent", shell, types.Null)
}

const usage = "Please use `get` or `set` (or use the bang (!) to unset) followed by a murex data-type"

func shell(p *lang.Process) error {
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
		return errors.New(err.Error() + ". " + usage)
	}

	dataType, err := p.Parameters.String(1)
	if err != nil {
		return errors.New(err.Error() + ". " + usage)
	}

	switch flag {
	case "get":
		agent, err := OpenAgents.Get(dataType)
		if err != nil {
			return err
		}
		_, err = p.Stdout.Write([]byte(string(agent.Block)))
		return err

	case "set":
		block, err := p.Parameters.Block(2)
		if err != nil {
			return err
		}
		OpenAgents.Set(dataType, block, p.FileRef)
		return nil

	default:
		return fmt.Errorf("invalid option\n%s", usage)
	}
}
