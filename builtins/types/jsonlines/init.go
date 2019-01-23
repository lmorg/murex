package jsonlines

import (
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types/define"
)

const name = "jsonl"

func init() {
	// Register data type
	define.Marshallers[name] = marshal
	define.Unmarshallers[name] = unmarshal
	define.ReadIndexes[name] = index
	define.ReadNotIndexes[name] = index

	stdio.RegesterReadArray(name, readArray)
	//stdio.RegesterReadMap(name, readMap)
	stdio.RegesterWriteArray(name, newArrayWriter)

	define.SetMime(name, "application/jsonl", "application/jsonlines")

	define.SetFileExtensions(name, "jsonl", "jsonlines", "murex_history")
}
