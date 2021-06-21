package unicode

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
)

const dataType = "utf8"

func init() {
	// Register data type
	lang.ReadIndexes[dataType] = index
	lang.ReadNotIndexes[dataType] = index
	lang.Marshallers[dataType] = marshal
	lang.Unmarshallers[dataType] = unmarshal

	stdio.RegisterReadArray(dataType, readArray)
	stdio.RegisterWriteArray(dataType, newArrayWriter)
}
