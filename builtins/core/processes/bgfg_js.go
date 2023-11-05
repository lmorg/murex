//go:build js
// +build js

package processes

import (
	"errors"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func mkbg(p *lang.Process) error {
	return errors.New("Invalid parameters. `bg` only supports a code block in js/wasm because processes are running in a sandboxed VM")
}

func cmdForeground(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	return errors.New("This function is currently not supported on js/wasm")
}

func unstop(p *lang.Process) {}
