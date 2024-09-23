package xml

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
)

func init() {
	// Register data type
	lang.Marshallers[dataType] = marshal
	lang.Unmarshallers[dataType] = UnmarshalFromProcess
	lang.ReadIndexes[dataType] = index
	lang.ReadNotIndexes[dataType] = index

	stdio.RegisterReadArray(dataType, readArray)
	stdio.RegisterReadArrayWithType(dataType, readArrayWithType)
	stdio.RegisterReadMap(dataType, readMap)
	stdio.RegisterWriteArray(dataType, newArrayWriter)

	lang.SetMime(dataType,
		"application/xml", // this is preferred, but we include the others incase a website sends a non-standard MIME time
		"application/x-xml",
		"text/xml",
		"text/x-xml",
		"application/xml-external-parsed-entity",
		"text/xml-external-parsed-entity",
		"application/xml-dtd",
		"+xml",
	)
	lang.SetFileExtensions(dataType, "xml", "svg", "xhtml", "xht", "rss", "atom")
}

const dataType = "xml"
