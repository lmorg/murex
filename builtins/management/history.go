package management

import (
	"errors"
	"fmt"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/utils"
)

func init() {
	proc.GoFunctions["history"] = proc.GoFunction{Func: cmdHistory, TypeIn: types.Null, TypeOut: types.Json}
	proc.GoFunctions["^"] = proc.GoFunction{Func: cmdHistCmd, TypeIn: types.Null, TypeOut: types.Generic}
}

func cmdHistory(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Json)
	if shell.Instance == nil {
		return errors.New("This is only designed to be run when the shell is in interactive mode.")
	}

	b, err := utils.JsonMarshal(shell.History.List)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Writeln(b)
	return err
}

func cmdHistCmd(p *proc.Process) (err error) {
	if shell.Instance == nil {
		return errors.New("This is only designed to be run when the shell is in interactive mode.")
	}

	var block string
	i, piErr := p.Parameters.Int(0)
	s, _ := p.Parameters.String(0)

	switch {
	case p.Parameters.Len() == 0:
		return errors.New("Missing parameters.")
		//block = shell.History.List[len(shell.History.List)-2].Block

	case piErr == nil:
		if i < 0 || i > len(shell.History.List) {
			return errors.New("Not a valid history index. Use `history` to pick an item.")
		}
		block = shell.History.List[i].Block

	case s == "!!":
		block = shell.History.Last

	default:
		return errors.New("I do not understand your request." + utils.NewLineString + piErr.Error())
	}

	fmt.Println("» " + block)
	p.Stdin.MakeParent()
	p.ExitNum, err = lang.ProcessNewBlock([]rune(block), p.Stdin, p.Stdout, p.Stderr, "^")

	return err
}
