//go:build !plan9
// +build !plan9

package signaltrap

import (
	"fmt"
	"os"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.DefineFunction(commandName, cmdSendSignal, types.Json)
}

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

	proc, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	if !isValidInterrupt(sig) {
		return fmt.Errorf("invalid signal name '%s'. Expecting something like 'SIGINT'", sig)
	}

	return proc.Signal(interrupts[sig])
}

func autocompleteSignals(p *lang.Process) error {
	p.Stdout.SetDataType(types.Json)

	signals := make(map[string]string, len(interrupts))
	for name := range interrupts {
		signals[name] = interrupts[name].String()
	}

	b, err := json.Marshal(signals, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}
