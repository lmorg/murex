package structs

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	proc.GoFunctions["bg"] = cmdBackground
}

func cmdBackground(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Null)

	var block []rune

	if p.IsMethod {
		b, err := p.Stdin.ReadAll()
		if err != nil {
			return err
		}
		block = []rune(string(b))

	} else {
		block, err = p.Parameters.Block(0)
		if err != nil {
			return err
		}
	}

	p.IsBackground = true
	p.WaitForTermination <- false
	lang.RunBlockExistingConfigSpace(block, p.Stdin, p.Stdout, p.Stderr, p)

	return nil
}
