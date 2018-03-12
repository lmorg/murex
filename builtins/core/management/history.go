package management

import (
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/readline"
)

func init() {
	proc.GoFunctions["history"] = cmdHistory
	proc.GoFunctions["^"] = cmdHistCmd
	//proc.GoFunctions["history-set-write-pipe"] = cmdHistPipe
}

func cmdHistory(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Json)
	if !shell.Interactive {
		return errors.New("This is only designed to be run when the shell is in interactive mode.")
	}

	var list []string
	for i := 0; i < readline.History.Len(); i++ {
		line, err := readline.History.GetLine(i)
		if err == nil {
			return err
		}
		list = append(list, line)
	}
	b, err := utils.JsonMarshal(list, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Writeln(b)
	return err
}

func cmdHistCmd(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)
	return errors.New("Invalid usage of history variable!")
}

/*func cmdHistPipe(p *proc.Process) error {
	if !shell.Interactive {
		return errors.New("This is only designed to be run when the shell is in interactive mode.")
	}

	p.Stdout.SetDataType(types.Null)

	name, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	if proc.GlobalPipes.Dump()[name] == "" {
		return errors.New("No pipe exists named: " + name)
	}

	pipe, err := proc.GlobalPipes.Get(name)
	if err != nil {
		return err
	}

	shell.History.Writer = pipe

	return nil
}*/
