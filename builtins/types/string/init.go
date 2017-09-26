package string

import (
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/lang/types/define"
)

func init() {
	// Register data type
	define.Marshal[types.String] = marshal
	define.Unmarshal[types.String] = unmarshal
	define.ReadIndexes[types.String] = index
	streams.ReadArray[types.Generic] = readArray
	streams.ReadMap[types.Generic] = readMap

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
