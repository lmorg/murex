//go:build deprecated_builtins
// +build deprecated_builtins

package io

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineMethod("tread", cmdTread, types.String, types.Null)
}

func cmdTread(p *lang.Process) error {
	//lang.FeatureDeprecatedBuiltin(p)

	dt, err := p.Parameters.String(0)
	if err != nil {
		return err
	}
	return read(p, dt, 1)
}
