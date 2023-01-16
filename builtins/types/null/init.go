package null

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	// Register data type
	lang.Marshallers[types.Null] = marshal
	lang.Unmarshallers[types.Null] = unmarshal
	stdio.RegisterWriteArray(types.Null, newArrayWriter)
}
