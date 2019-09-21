package csv

import (
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types"
)

const typeName = "csv"

func init() {
	//stdio.RegesterReadArray(typeName, readArray)
	stdio.RegesterReadMap(typeName, readMap)

	lang.ReadIndexes[typeName] = readIndex
	lang.ReadNotIndexes[typeName] = readIndex

	lang.Marshallers[typeName] = marshal
	lang.Unmarshallers[typeName] = unmarshal

	// `application/csv` and `text/csv` are the common ones. `x-csv` is added just in case anyone decides to use
	// something non-standard.
	lang.SetMime(typeName,
		"application/csv",
		"application/x-csv",
		"text/csv",
		"text/x-csv",
	)

	lang.SetFileExtensions(typeName, "csv")

	lang.InitConf.Define("csv", "separator", config.Properties{
		Description: "The delimiter for records in a CSV file.",
		Default:     `,`,
		DataType:    types.String,
	})

	lang.InitConf.Define("csv", "comment", config.Properties{
		Description: "The prefix token for comments in a CSV table.",
		Default:     `#`,
		DataType:    types.String,
	})
}
