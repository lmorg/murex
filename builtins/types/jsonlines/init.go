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

		"application/json-lines",
		"application/x-json-lines",
		"text/json-lines",
		"text/x-json-lines",

		"application/jsonseq",
		"application/x-jsonseq",
		"text/jsonseq",
		"text/x-jsonseq",

		"application/json-seq",
		"application/x-json-seq",
		"text/json-seq",
		"text/x-json-seq",

		"application/ldjson",
		"application/x-ldjson",
		"text/ldjson",
		"text/x-ldjson",

		"application/ndjson",
		"application/x-ndjson",
		"text/ndjson",
		"text/x-ndjson",
	)

	define.SetFileExtensions(types.JsonLines,
		"jsonl",
		"jsonlines",
		"json-lines",
		"jsons",
		"jsonseq",
		"json-seq",
		"ldjson",
		"ndjson",
		"murex_history",
	)
}
