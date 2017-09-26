package management

import (
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/utils"
)

func init() {
	proc.GoFunctions["history"] = cmdHistory
	proc.GoFunctions["^"] = cmdHistCmd
}

func cmdHistory(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Json)
	if shell.Instance == nil {
		return errors.New("This is only designed to be run when the shell is in interactive mode.")
	}

	b, err := utils.JsonMarshal(shell.History.List, p.Stdout.IsTTY())
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
