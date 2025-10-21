package json

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	// Register data type
	lang.RegisterDataType(types.Json, lang.DataTypeIsObject)
	lang.RegisterDataType(types.Json, lang.DataTypeIsObject)
	lang.RegisterMarshaller(types.Json, marshal)
	lang.RegisterUnmarshaller(types.Json, unmarshal)
	lang.ReadIndexes[types.Json] = index
	lang.ReadNotIndexes[types.Json] = index

	stdio.RegisterReadArray(types.Json, readArray)
	stdio.RegisterReadArrayWithType(types.Json, readArrayWithType)
	stdio.RegisterReadMap(types.Json, readMap)
	stdio.RegisterWriteArray(types.Json, newArrayWriter)

	lang.SetMime(types.Json,
		"application/json", // this is preferred, but we include the others incase a website sends a non-standard MIME time
		"application/x-json",
		"text/json",
		"text/x-json",
		"+json",
	)
	lang.SetFileExtensions(types.Json, "json", "tfstate")
}
