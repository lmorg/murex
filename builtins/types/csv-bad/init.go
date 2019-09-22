package csvbad

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/stdio"
)

const typeName = "csvbad"

func init() {
	//stdio.RegesterReadArray(typeName, readArray)
	stdio.RegesterReadMap(typeName, readMap)

	lang.ReadIndexes[typeName] = readIndex
	lang.ReadNotIndexes[typeName] = readIndex

	lang.Marshallers[typeName] = marshal
	lang.Unmarshallers[typeName] = unmarshal

	// The following syntax is defined in the `csv` type.
	// Uncomment it if - for whatever reason - you want to disable the default
	// CSV marshaller. You may also want to change the `typeName` const (above)
	// to just `csv`.
	/*
		// `application/csv` and `text/csv` are the common ones. `x-csv` is added just in case anyone decides to use
		// something non-standard.
		lang.SetMime(typeName,
			"application/csv",
			"application/x-csv",
			"text/csv",
			"text/x-csv",
		)

		//lang.SetFileExtensions(typeName, "csv")

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

		lang.InitConf.Define("csv", "headings", config.Properties{
			Description: "CSV files include headings when queried in formap.",
			Default:     true,
			DataType:    types.Boolean,
		})
	*/
}
