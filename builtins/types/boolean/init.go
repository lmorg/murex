package boolean

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	// Register data types
	lang.Marshallers[types.Boolean] = marshal
	lang.Unmarshallers[types.Boolean] = unmarshal
}
