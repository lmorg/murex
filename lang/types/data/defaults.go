package data

import "github.com/lmorg/murex/lang/types"

func init() {
	// Register builtin data types
	Marshal[types.String] = marshalString
	Unmarshal[types.String] = unmarshalString

	Marshal[types.Json] = marshalJson
	Unmarshal[types.Json] = unmarshalJson

	ReadIndexes[types.Json] = indexJson
	ReadIndexes[types.String] = indexString
	ReadIndexes[types.Generic] = indexGeneric

	SetMime(types.String,
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

	SetMime(types.Json, "application/json")

	SetMime(types.Binary, "multipart/x-zip")

	SetFileExtensions(types.Json, "json")
	SetFileExtensions(types.Binary, "bin")
}
