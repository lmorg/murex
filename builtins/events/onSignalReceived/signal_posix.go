//go:build !plan9 && !js
// +build !plan9,!js

package signaltrap

import (
	"fmt"
	"os"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func cmdSendSignal(p *lang.Process) error {
	if p.Parameters.Len() == 0 {
		return autocompleteSignals(p)
	}

	p.Stdout.SetDataType(types.Null)

	pid, err := p.Parameters.Int(0)
	if err != nil {
		return err
	}

	sig, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	if !isValidInterrupt(sig) {
		return fmt.Errorf("invalid signal name '%s'. Expecting something like 'SIGINT'", sig)
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	return proc.Signal(interrupts[sig])
}
