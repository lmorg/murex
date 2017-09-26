package csv

import (
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types/define"
)

const typeName = "csv"

func init() {
	//streams.ReadArray[typeName] = readArray
	streams.ReadMap[typeName] = readMap
	define.ReadIndexes[typeName] = readIndex

	define.Marshal[typeName] = marshal
	define.Unmarshal[typeName] = unmarshal

	// `application/csv` and `text/csv` are the common ones. `x-csv` is added just in case anyone decides to use
	// something non-standard.
	define.SetMime(typeName,
		"application/csv",
		"application/x-csv",
		"text/csv",
		"text/x-csv",
	)

	define.SetFileExtensions(typeName, "csv")
}
