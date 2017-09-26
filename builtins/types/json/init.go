package json

import (
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/lang/types/define"
)

func init() {
	// Register data type
	define.Marshal[types.Json] = marshal
	define.Unmarshal[types.Json] = unmarshal
	define.ReadIndexes[types.Json] = index
	streams.ReadArray[types.Json] = readArray
	streams.ReadMap[types.Json] = readMap

	define.SetMime(types.Json, "application/json")

	define.SetFileExtensions(types.Json, "json")
}
