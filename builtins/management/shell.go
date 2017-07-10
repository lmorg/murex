package management

import (
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell"
	"io"
	"os"
)

func init() {
	proc.GoFunctions["history"] = proc.GoFunction{Func: cmdHistory, TypeIn: types.Null, TypeOut: types.Json}
}

func cmdHistory(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.String)
	if shell.Instance == nil {
		return errors.New("This is only designed to be run when the shell is in interactive mode.")
	}

	file, err := os.Open(shell.Instance.Config.HistoryFile)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(p.Stdout, file)
	return err
}
