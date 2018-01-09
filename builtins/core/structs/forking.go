package structs

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"errors"
)

func init() {
	proc.GoFunctions["fork"] = cmdFork
}

func cmdFork(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Null)

	if p.Parameters.Len() ==0 {
		return errors.New("Nothing to fork.")
	}

	block, err := p.Parameters.Block(0)
	if err != nil {
		s := p.Parameters.StringAll()
		block = []rune(s)
	}

	p.IsBackground = true
	p.WaitForTermination <- false
	lang.ProcessNewBlock(block, p.Stdin, p.Stdout, p.Stderr, p)

	return nil
}
