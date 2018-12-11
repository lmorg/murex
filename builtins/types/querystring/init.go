package string

import (
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types/define"
)

const dataType = "qs"

func init() {
	// Register data type
	define.Marshallers[dataType] = marshal
	define.Unmarshallers[dataType] = unmarshal
	define.ReadIndexes[dataType] = index

	stdio.RegesterReadArray(dataType, readArray)
	stdio.RegesterReadMap(dataType, readMap)
}
