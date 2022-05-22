package jsonconcat

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
)

const name = "jsonc"

func init() {
	// Register data type
	lang.Marshallers[name] = marshal
	lang.Unmarshallers[name] = unmarshal
	lang.ReadIndexes[name] = index
	lang.ReadNotIndexes[name] = index

	stdio.RegisterReadArray(name, readArray)
	stdio.RegisterReadArrayWithType(name, readArrayWithType)
	//stdio.RegisterReadMap(name, readMap)
	stdio.RegisterWriteArray(name, newArrayWriter)

	lang.SetMime(name,
		"application/jsonc",
		"application/x-jsonc",
		"text/jsonc",
		"text/x-jsonc",

		"application/jsonconcat",
		"application/x-jsonconcat",
		"text/jsonconcat",
		"text/x-jsonconcat",

		"application/concatenated-json",
		"application/x-concatenated-json",
		"text/concatenated-json",
		"text/concatenated-json",

		"application/jsonseq",
		"application/x-jsonseq",
		"text/jsonseq",
		"text/x-jsonseq",

		"application/json-seq",
		"application/x-json-seq",
		"text/json-seq",
		"text/x-json-seq",
	)

	lang.SetFileExtensions(name,
		"jsonc",
		"jsonconcat",
		"concatenated-json",
		"jsons",
		"jsonseq",
		"json-seq",
	)
}
