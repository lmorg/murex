package ranges

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineMethod("@[", deprecatedRange, types.ReadArray, types.WriteArray)
}

func deprecatedRange(p *lang.Process) error {
	//lang.FeatureDeprecatedBuiltin(p)
	return CmdRange(p)
}
