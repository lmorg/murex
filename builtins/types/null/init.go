package null

import (
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	// Register data type
	stdio.RegisterWriteArray(types.Null, newArrayWriter)
}
