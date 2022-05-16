package coreutils

import (
	"os"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineFunction("pwd", pwd, types.String)
}

func pwd(p *lang.Process) error {
	p.Stdout.SetDataType(types.String)

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	_, err = p.Stdout.Writeln([]byte(dir))
	return err
}
