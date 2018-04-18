package numeric

import (
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/lang/types/define"
)

func init() {
	// Register data types
	define.Marshallers[types.Integer] = marshalInt
	define.Unmarshallers[types.Integer] = unmarshalInt

	define.Marshallers[types.Float] = marshalFloat
	define.Unmarshallers[types.Float] = unmarshalFloat

	define.Marshallers[types.Number] = marshalNumber
	define.Unmarshallers[types.Number] = unmarshalNumber
}
