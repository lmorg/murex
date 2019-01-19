package coreutils

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

// These are really lazy aliases from POSIX to Windows. Since they don't intend to emulate the behavior of their POSIX
// counterparts we will always return a non-zero exit number so sensitive scripts can fail gracefully.

func init() {
	lang.GoFunctions["ps"] = alias("tasklist")
	lang.GoFunctions["ls"] = alias("dir")
	lang.GoFunctions["rm"] = alias("del")
	lang.GoFunctions["clear"] = alias("cls")
	lang.GoFunctions["cat"] = alias("type")
}

func alias(cmd string) func(p *lang.Process) error {
	return func(p *lang.Process) error {
		p.Stdout.SetDataType(types.String)

		p.Name = "exec"

		p.Parameters.Params = append([]string{cmd}, p.Parameters.Params...)

		err := lang.External(p)

		p.ExitNum = 13

		return err
	}
}
