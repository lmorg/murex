package generic

import (
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/lang/types/define"
)

func init() {
	// Register data type
	define.ReadIndexes[types.Generic] = index
	define.ReadNotIndexes[types.Generic] = index
	define.Marshallers[types.Generic] = marshal
	define.Unmarshallers[types.Generic] = unmarshal
	streams.ReadArray[types.Generic] = readArray
	streams.ReadMap[types.Generic] = readMap
}
