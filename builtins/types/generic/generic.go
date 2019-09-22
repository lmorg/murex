package generic

import (
	"regexp"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	// Register data type
	lang.ReadIndexes[types.Generic] = index
	lang.ReadNotIndexes[types.Generic] = index
	lang.Marshallers[types.Generic] = marshal
	lang.Unmarshallers[types.Generic] = unmarshal

	stdio.RegesterReadArray(types.Generic, readArray)
	stdio.RegesterReadMap(types.Generic, readMap)
	stdio.RegesterWriteArray(types.Generic, newArrayWriter)
}

var rxWhitespace = regexp.MustCompile(`\s+`)
