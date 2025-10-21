package string

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	// Register data type
	lang.RegisterDataType(types.String, lang.DataTypeIsList)
	stdio.RegisterReadArray(types.String, readArray)
	stdio.RegisterReadArrayWithType(types.String, readArrayWithType)
	stdio.RegisterReadMap(types.String, readMap)
	stdio.RegisterWriteArray(types.String, newArrayWriter)

	lang.ReadIndexes[types.String] = index
	lang.ReadNotIndexes[types.String] = index
	lang.RegisterMarshaller(types.String, marshal)
	lang.RegisterUnmarshaller(types.String, unmarshal)

	lang.SetMime(types.String,
		"application/x-latex",
		"www/mime",
		"application/base64",
		"application/postscript",
		"application/rtf", "application/x-rtf",
		"application/x-sh", "application/x-bsh", "application/x-shar",
		"application/plain",
		"application/x-tcl",
		"model/vrml", "x-world/x-vrml", "application/x-vrml",
		"image/svg+xml",
		"application/javascript", "application/x-javascript",
		"application/xml",
	)

	// descriptive name
	lang.RegisterDataType("string", lang.DataTypeIsList)
	stdio.RegisterReadArray("string", readArray)
	stdio.RegisterReadMap("string", readMap)
	stdio.RegisterWriteArray("string", newArrayWriter)

	lang.ReadIndexes["string"] = index
	lang.ReadNotIndexes["string"] = index
	lang.RegisterMarshaller("string", marshal)
	lang.RegisterUnmarshaller("string", unmarshal)

}
