package coreutils

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"os"
)

func init() {
	//proc.GoFunctions["ls"] = cmdLs
	proc.GoFunctions["pwd"] = pwd
}

func pwd(p *proc.Process) error {
	p.Stdout.SetDataType(types.String)

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	_, err = p.Stdout.Writeln([]byte(dir))
	return err
}
