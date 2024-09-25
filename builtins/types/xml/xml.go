package xml

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
)

func init() {
	// Register data type
	lang.RegisterMarshaller(typeName, marshal)
	lang.RegisterUnmarshaller(typeName, UnmarshalFromProcess)
	lang.ReadIndexes[typeName] = index
	lang.ReadNotIndexes[typeName] = index

	stdio.RegisterReadArray(typeName, readArray)
	stdio.RegisterReadArrayWithType(typeName, readArrayWithType)
	stdio.RegisterReadMap(typeName, readMap)
	stdio.RegisterWriteArray(typeName, newArrayWriter)

	lang.SetMime(typeName,
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
	lang.SetFileExtensions(typeName,
		// xml documents
		"xml",
		"xhtml", "xht", // xhtml
		"rss", "atom", // rss
		"svg",   // xml-based images
		"plist", // manifest files

		// while valid SGML might not be valid XML, a lot of SGML documents are
		// ostensibly XML and since we don't have an SGML parser, lets include
		// SGMLs extension in `xml` because it might still work _some_ of the
		// time.
		"sgml",
	)
}

const typeName = "xml"

const (
	xmlDefaultRoot    = "xml"
	xmlDefaultElement = "list"
)
