package string

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/stdio"
)

const dataType = "qs"

func init() {
	// Register data type
	lang.Marshallers[dataType] = marshal
	lang.Unmarshallers[dataType] = unmarshal
	lang.ReadIndexes[dataType] = index

	stdio.RegesterReadArray(dataType, readArray)
	stdio.RegesterReadMap(dataType, readMap)
}
