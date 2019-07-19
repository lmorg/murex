package management

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/shell/history"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.GoFunctions["history"] = cmdHistory
	//lang.GoFunctions["history-set-write-pipe"] = cmdHistPipe
}

func cmdHistory(p *lang.Process) (err error) {
	//if !shell.Interactive {
	//	return errors.New("This is only designed to be run when the shell is in interactive mode")
	//}

	list := shell.Prompt.History.Dump().([]history.Item)

	// If outputting to the terminal then lets just do pure JSON for readability
	if p.Stdout.IsTTY() {
		p.Stdout.SetDataType(types.Json)
		b, err := json.Marshal(list, p.Stdout.IsTTY())
		if err != nil {
			return err
		}

		_, err = p.Stdout.Writeln(b)
		return err
	}

	// if not outputting to the terminal, then use jsonlines instead for easier
	// grepping et al

	p.Stdout.SetDataType(types.JsonLines)

	aw, err := p.Stdout.WriteArray(types.JsonLines)
	if err != nil {
		return err
	}

	for i := range list {
		b, err := json.Marshal(list[i], p.Stdout.IsTTY())
		if err != nil {
			return err
		}

		err = aw.Write(b)
		if err != nil {
			return err
		}
	}

	return nil
}

/*func cmdHistPipe(p *lang.Process) error {
	if !shell.Interactive {
		return errors.New("This is only designed to be run when the shell is in interactive mode.")
	}

	p.Stdout.SetDataType(types.Null)

	name, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	if lang.GlobalPipes.Dump()[name] == "" {
		return errors.New("No pipe exists named: " + name)
	}

	pipe, err := lang.GlobalPipes.Get(name)
	if err != nil {
		return err
	}

	shell.History.Writer = pipe

	return nil
}*/
