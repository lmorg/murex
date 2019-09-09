package jsonlines

import (
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/lang/types/define"
)

func init() {
	// Register data type
	define.Marshallers[types.JsonLines] = marshal
	define.Unmarshallers[types.JsonLines] = unmarshal
	define.ReadIndexes[types.JsonLines] = index
	define.ReadNotIndexes[types.JsonLines] = index

	stdio.RegesterReadArray(types.JsonLines, readArray)
	//stdio.RegesterReadMap(name, readMap)
	stdio.RegesterWriteArray(types.JsonLines, newArrayWriter)

	define.SetMime(types.JsonLines,
		"application/jsonl",
		"application/x-jsonl",
		"text/jsonl",
		"text/x-jsonl",

		"application/jsonlines",
		"application/x-jsonlines",
		"text/jsonlines",
		"text/x-jsonlines",
	)

	define.SetFileExtensions(types.JsonLines,
		"jsonl",
		"jsonlines",
		"murex_history",
	)
}
