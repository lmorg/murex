//go:build js
// +build js

package signaltrap

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func cmdSendSignal(p *lang.Process) error {
	if p.Parameters.Len() == 0 {
		return autocompleteSignals(p)
	}

	p.Stdout.SetDataType(types.Null)

	return fmt.Errorf("`%s` is not supported on WASM", commandName)
}
