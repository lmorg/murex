package numeric

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	// Register data types
	lang.Marshallers[types.Integer] = marshalInt
	lang.Unmarshallers[types.Integer] = unmarshalInt

	lang.Marshallers[types.Float] = marshalFloat
	lang.Unmarshallers[types.Float] = unmarshalFloat

	lang.Marshallers[types.Number] = marshalNumber
	lang.Unmarshallers[types.Number] = unmarshalNumber
}
