package csv

import (
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/lang/types/define"
)

const typeName = "csv"

func init() {
	//streams.ReadArray[typeName] = readArray
	streams.ReadMap[typeName] = readMap
	define.ReadIndexes[typeName] = readIndex
	define.ReadNotIndexes[typeName] = readIndex

	define.Marshallers[typeName] = marshal
	define.Unmarshallers[typeName] = unmarshal

	// `application/csv` and `text/csv` are the common ones. `x-csv` is added just in case anyone decides to use
	// something non-standard.
	define.SetMime(typeName,
		"application/csv",
		"application/x-csv",
		"text/csv",
		"text/x-csv",
	)

	define.SetFileExtensions(typeName, "csv")

	proc.InitConf.Define("csv", "separator", config.Properties{
		Description: "The delimiter for records in a CSV file.",
		Default:     `,`,
		DataType:    types.String,
	})

	proc.InitConf.Define("csv", "comment", config.Properties{
		Description: "The prefix token for comments in a CSV table.",
		Default:     `#`,
		DataType:    types.String,
	})

	proc.InitConf.Define("csv", "headings", config.Properties{
		Description: "CSV files include headings when queried in formap.",
		Default:     true,
		DataType:    types.Boolean,
	})
}
