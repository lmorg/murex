package cmdpipe

import (
	"io"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/utils/consts"
)

func init() {
	lang.GoFunctions[consts.NamedPipeProcName] = cmdReadPipe
}

func cmdReadPipe(p *lang.Process) error {
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

	if p.IsMethod {
		pipe.SetDataType(p.Stdin.GetDataType())
		_, err = io.Copy(pipe, p.Stdin)
		return err
	}

	p.Stdout.SetDataType(pipe.GetDataType())
	_, err = io.Copy(p.Stdout, pipe)
	return err
}
