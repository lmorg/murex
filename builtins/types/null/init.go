package null

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	// Register data type
	lang.RegisterMarshaller(types.Null, marshal)
	lang.RegisterUnmarshaller(types.Null, unmarshal)
	stdio.RegisterWriteArray(types.Null, newArrayWriter)
}
