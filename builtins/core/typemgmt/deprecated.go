//go:build deprecated_builtins
// +build deprecated_builtins

package typemgmt

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineFunction("die", cmdDie, types.Null)
}

func cmdDie(p *lang.Process) error {
	p.Stdout.SetDataType(types.Die)

	//lang.FeatureDeprecatedBuiltin(p)

	lang.Exit(1)
	return nil
}
