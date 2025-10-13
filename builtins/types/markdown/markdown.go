package markdown

import (
	"github.com/lmorg/murex/lang"
)

const typeName = "md"

func init() {
	//stdio.RegisterReadArray(typeName, readArray)
	//stdio.RegisterReadMap(typeName, readMap)

	//lang.ReadIndexes[typeName] = readIndex
	//lang.ReadNotIndexes[typeName] = readIndex

	lang.RegisterMarshaller(typeName, marshal)
	//lang.RegisterUnmarshaller(typeName, unmarshal)

	/*lang.SetMime(typeName,
		"application/csv",
		"application/x-csv",
		"text/csv",
		"text/x-csv",
		"+csv",
	)*/

	lang.SetFileExtensions(typeName, "md")
}
