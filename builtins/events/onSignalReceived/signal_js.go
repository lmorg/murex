//go:build js
// +build js

package signaltrap

import (
	"fmt"

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

	return fmt.Errorf("`%s` is not supported on WASM", commandName)
}

func autocompleteSignals(p *lang.Process) error {
	p.Stdout.SetDataType(types.Json)

	signals := make(map[string]string, 0)

	b, err := json.Marshal(signals, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}
