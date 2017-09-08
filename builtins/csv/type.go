package csv

import (
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types/data"
)

const typeName = "csv"

func init() {
	//streams.ReadArray[typeName] = readArray
	streams.ReadMap[typeName] = readMap
	data.ReadIndexes[typeName] = readIndex

	data.Marshal[typeName] = marshal
	data.Unmarshal[typeName] = unmarshal

	// `application/csv` and `text/csv` are the common ones. `x-csv` is added just in case anyone decides to use
	// something non-standard.
	data.SetMime(typeName,
		"application/csv",
		"application/x-csv",
		"text/csv",
		"text/x-csv",
	)

	data.SetFileExtensions(typeName, "csv")
}
