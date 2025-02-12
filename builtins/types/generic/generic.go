package generic

import (
	"regexp"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	// Register data type
	lang.ReadIndexes[types.Generic] = index
	lang.ReadNotIndexes[types.Generic] = index
	lang.RegisterMarshaller(types.Generic, marshal)
	lang.RegisterUnmarshaller(types.Generic, unmarshal)

	stdio.RegisterReadArray(types.Generic, readArray)
	stdio.RegisterReadArrayWithType(types.Generic, readArrayWithType)
	stdio.RegisterReadMap(types.Generic, readMap)
	stdio.RegisterWriteArray(types.Generic, newArrayWriter)

	// descriptive name
	lang.ReadIndexes["generic"] = index
	lang.ReadNotIndexes["generic"] = index
	lang.RegisterMarshaller("generic", marshal)
	lang.RegisterUnmarshaller("generic", unmarshal)

	stdio.RegisterReadArray("generic", readArray)
	stdio.RegisterReadMap("generic", readMap)
	stdio.RegisterWriteArray("generic", newArrayWriter)
}

var rxWhitespace = regexp.MustCompile(`\s+`)

// common tabwriter values
const (
	twMinWidth = 0
	twTabWidth = 0
	twPadding  = 2
	twPadChar  = ' '
	twFlags    = 0
)
