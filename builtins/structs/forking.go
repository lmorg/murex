package structs

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	proc.GoFunctions["fork"] = proc.GoFunction{Func: cmdFork, TypeIn: types.Generic, TypeOut: types.Generic}
}

func cmdFork(p *proc.Process) (err error) {
	block, err := p.Parameters.Block(0)
	if err != nil {
		return err
	}

	p.IsBackground = true
	p.WaitForTermination <- false
	lang.ProcessNewBlock(block, p.Stdin, p.Stdout, p.Stderr, p)

	return
}
