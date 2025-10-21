package boolean

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	// Register data types
	lang.RegisterDataType(types.Boolean, lang.DataTypeIsBoolean)
	lang.RegisterMarshaller(types.Boolean, marshal)
	lang.RegisterUnmarshaller(types.Boolean, unmarshal)
}
