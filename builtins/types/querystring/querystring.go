package string

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
)

const dataType = "qs"

func init() {
	// Register data type
	lang.RegisterDataType(dataType, lang.DataTypeIsKeyValue)
	lang.RegisterMarshaller(dataType, marshal)
	lang.RegisterUnmarshaller(dataType, unmarshal)
	lang.ReadIndexes[dataType] = index

	stdio.RegisterReadArray(dataType, readArray)
	stdio.RegisterReadArrayWithType(dataType, readArrayWithType)
	stdio.RegisterReadMap(dataType, readMap)
}
