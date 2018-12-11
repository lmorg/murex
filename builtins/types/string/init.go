package string

import (
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/lang/types/define"
)

func init() {
	// Register data type
	stdio.RegesterReadArray(types.String, readArray)
	stdio.RegesterReadMap(types.String, readMap)
	stdio.RegesterWriteArray(types.String, newArrayWriter)

	define.ReadIndexes[types.String] = index
	define.ReadNotIndexes[types.String] = index
	define.Marshallers[types.String] = marshal
	define.Unmarshallers[types.String] = unmarshal

	define.SetMime(types.String,
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
		"application/xml")

}
