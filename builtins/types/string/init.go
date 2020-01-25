package string

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	// Register data type
	stdio.RegesterReadArray(types.String, readArray)
	stdio.RegesterReadMap(types.String, readMap)
	stdio.RegesterWriteArray(types.String, newArrayWriter)

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
	stdio.RegesterReadArray("string", readArray)
	stdio.RegesterReadMap("string", readMap)
	stdio.RegesterWriteArray("string", newArrayWriter)

	lang.ReadIndexes["string"] = index
	lang.ReadNotIndexes["string"] = index
	lang.Marshallers["string"] = marshal
	lang.Unmarshallers["string"] = unmarshal

}
