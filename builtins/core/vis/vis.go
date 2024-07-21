package vis

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineMethod("vis", vis, types.Any, types.Generic)
}

func vis(p *lang.Process) error {
	p.Stdout.SetDataType(types.Generic)

	return tree(p)
}

func tree(p *lang.Process) error {
	//stdio.
	return nil
}
