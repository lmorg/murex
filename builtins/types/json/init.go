package json

import (
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/lang/types/define"
)

func init() {
	// Register data type
	define.Marshallers[types.Json] = marshal
	define.Unmarshallers[types.Json] = unmarshal
	define.ReadIndexes[types.Json] = index
	define.ReadNotIndexes[types.Json] = index

	stdio.RegesterReadArray(types.Json, readArray)
	stdio.RegesterReadMap(types.Json, readMap)
	stdio.RegesterWriteArray(types.Json, newArrayWriter)

	define.SetMime(types.Json, "application/json")

	define.SetFileExtensions(types.Json, "json")
}
