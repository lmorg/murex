package string

import (
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types/define"
)

const dataType = "qs"

func init() {
	// Register data type
	define.Marshallers[dataType] = marshal
	define.Unmarshallers[dataType] = unmarshal
	define.ReadIndexes[dataType] = index
	streams.ReadArray[dataType] = readArray
	streams.ReadMap[dataType] = readMap
}
