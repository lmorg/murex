package string

import (
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	// Register data type
	stdio.RegesterWriteArray(types.Null, newArrayWriter)
}
