package coreutils

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"os"
)

func init() {
	//lang.GoFunctions["ls"] = cmdLs
	lang.GoFunctions["pwd"] = pwd
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
