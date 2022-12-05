package cmdpipe

import (
	"errors"
	"io"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/consts"
)

func init() {
	lang.DefineMethod("(murex named pipe)", cmdMurexNamedPipe, types.Null, types.Any)
	lang.DefineMethod(consts.NamedPipeProcName, cmdMurexNamedPipe, types.Any, types.Any)
}

func cmdMurexNamedPipe(p *lang.Process) error {
	name, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	var pipe stdio.Io

	if name == "stdin" {
		pipe = p.Scope.Stdin

	} else {
		pipe, err = lang.GlobalPipes.Get(name)
		if err != nil {
			return err
		}
	}

	if pipe == nil {
		return errors.New("STDIN is null")
	}

	if p.IsMethod {
		pipe.SetDataType(p.Stdin.GetDataType())
		_, err = io.Copy(pipe, p.Stdin)
		return err
	}

	p.Stdout.SetDataType(pipe.GetDataType())
	_, err = io.Copy(p.Stdout, pipe)
	return err
}
