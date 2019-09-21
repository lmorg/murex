package json

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	// Register data type
	lang.Marshallers[types.Json] = marshal
	lang.Unmarshallers[types.Json] = unmarshal
	lang.ReadIndexes[types.Json] = index
	lang.ReadNotIndexes[types.Json] = index

	stdio.RegesterReadArray(types.Json, readArray)
	stdio.RegesterReadMap(types.Json, readMap)
	stdio.RegesterWriteArray(types.Json, newArrayWriter)

	lang.SetMime(types.Json,
		"application/json", // this is preferred, but we include the others incase a website sends a non-standard MIME time
		"application/x-json",
		"text/json",
		"text/x-json",
	)
	lang.SetFileExtensions(types.Json, "json")
}
