package signaltrap

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

const commandName = "signal"

func init() {
	lang.DefineFunction(commandName, cmdSendSignal, types.Json)
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
