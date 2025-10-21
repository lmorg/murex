package numeric

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	// Register data types
	lang.RegisterDataType(types.Integer, lang.DataTypeIsNumeric)
	lang.RegisterMarshaller(types.Integer, marshalInt)
	lang.RegisterUnmarshaller(types.Integer, unmarshalInt)

	lang.RegisterDataType(types.Float, lang.DataTypeIsNumeric)
	lang.RegisterMarshaller(types.Float, marshalFloat)
	lang.RegisterUnmarshaller(types.Float, unmarshalFloat)

	lang.RegisterDataType(types.Number, lang.DataTypeIsNumeric)
	lang.RegisterMarshaller(types.Number, marshalNumber)
	lang.RegisterUnmarshaller(types.Number, unmarshalNumber)
}
