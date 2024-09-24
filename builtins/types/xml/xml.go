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

		// while valid SGML might not be valid XML, a lot of SGML documents are
		// ostensibly XML and since we don't have an SGML parser, lets include
		// SGMLs MIMEs in `xml` because it might still work _some_ of the time.
		"application/sgml",
		"application/x-sgml",
		"text/sgml",
		"text/x-sgml",
		"+sgml",
	)
	lang.SetFileExtensions(dataType,
		// xml documents
		"xml", "svg", "xhtml", "xht", "rss", "atom",

		// while valid SGML might not be valid XML, a lot of SGML documents are
		// ostensibly XML and since we don't have an SGML parser, lets include
		// SGMLs extension in `xml` because it might still work _some_ of the
		// time.
		"sgml",
	)
}

const dataType = "xml"
