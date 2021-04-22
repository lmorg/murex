package string

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	// Register data type
	stdio.RegisterReadArray(types.String, readArray)
	stdio.RegisterReadArrayByType(types.String, readArrayByType)
	stdio.RegisterReadMap(types.String, readMap)
	stdio.RegisterWriteArray(types.String, newArrayWriter)

	lang.ReadIndexes[types.String] = index
	lang.ReadNotIndexes[types.String] = index
	lang.Marshallers[types.String] = marshal
	lang.Unmarshallers[types.String] = unmarshal

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
	stdio.RegisterReadArray("string", readArray)
	stdio.RegisterReadMap("string", readMap)
	stdio.RegisterWriteArray("string", newArrayWriter)

	lang.ReadIndexes["string"] = index
	lang.ReadNotIndexes["string"] = index
	lang.Marshallers["string"] = marshal
	lang.Unmarshallers["string"] = unmarshal

}
