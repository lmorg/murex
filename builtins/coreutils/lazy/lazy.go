package lazy

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/consts"
)

// These are really lazy aliases from POSIX to Windows. Since they don't intend to emulate the behavior of their POSIX
// counterparts we will always return a non-zero exit number so sensitive scripts can fail gracefully.

func init() {
	proc.GoFunctions["ps"] = alias("tasklist")
	proc.GoFunctions["ls"] = alias("dir")
	proc.GoFunctions["rm"] = alias("del")
	proc.GoFunctions["clear"] = alias("cls")
	proc.GoFunctions["cat"] = alias("type")
}

func alias(cmd string) func(p *proc.Process) error {
	return func(p *proc.Process) error {
		p.Stdout.SetDataType(types.String)

		if !p.IsMethod && p.Stdout.IsTTY() {
			p.Name = consts.CmdPty
		} else {
			p.Name = consts.CmdExec
		}

		p.Parameters.Params = append([]string{cmd}, p.Parameters.Params...)

		err := proc.External(p)

		p.ExitNum = 13

		return err
	}
}
