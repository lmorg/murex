package numeric

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	// Register data types
	lang.RegisterMarshaller(types.Integer, marshalInt)
	lang.RegisterUnmarshaller(types.Integer, unmarshalInt)

	lang.RegisterMarshaller(types.Float, marshalFloat)
	lang.RegisterUnmarshaller(types.Float, unmarshalFloat)

	lang.RegisterMarshaller(types.Number, marshalNumber)
	lang.RegisterUnmarshaller(types.Number, unmarshalNumber)
}
