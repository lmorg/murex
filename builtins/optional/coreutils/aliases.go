package coreutils

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

// These are really lazy aliases from POSIX to Windows. Since they don't intend to emulate the behavior of their POSIX
// counterparts we will always return a non-zero exit number so sensitive scripts can fail gracefully.

func init() {
	lang.DefineFunction("ps", alias("tasklist"), types.Generic)
	lang.DefineFunction("ls", alias("dir"), types.Generic)
	lang.DefineFunction("rm", alias("del"), types.Null)
	lang.DefineFunction("clear", alias("cls"), types.Null)
	lang.DefineFunction("cat", alias("types"), types.Generic)
}

func alias(cmd string) func(p *lang.Process) error {
	return func(p *lang.Process) error {
		p.Stdout.SetDataType(types.Generic)

		p.Name.Set("exec")

		p.Parameters.Prepend([]string{cmd})

		err := lang.External(p)

		return err
	}
}
