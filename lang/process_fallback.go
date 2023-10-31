//go:build windows || plan9 || js
// +build windows plan9 js

package lang

import (
	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang/types"
)

func ttys(p *Process) {
	if p.CCExists != nil && p.CCExists(p.Name.String()) {
		p.Stderr, p.CCErr = streams.NewTee(p.Stderr)
		p.CCErr.SetDataType(types.Generic)

		p.Stdout, p.CCOut = streams.NewTee(p.Stdout)
	}
}
