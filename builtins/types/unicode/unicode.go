package unicode

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/stdio"
)

const dataType = "utf8"

func init() {
	// Register data type
	lang.ReadIndexes[dataType] = index
	lang.ReadNotIndexes[dataType] = index
	lang.Marshallers[dataType] = marshal
	lang.Unmarshallers[dataType] = unmarshal

	stdio.RegesterReadArray(dataType, readArray)
	stdio.RegesterWriteArray(dataType, newArrayWriter)
}
