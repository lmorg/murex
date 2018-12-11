package generic

import (
	"regexp"

	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/lang/types/define"
)

func init() {
	// Register data type
	define.ReadIndexes[types.Generic] = index
	define.ReadNotIndexes[types.Generic] = index
	define.Marshallers[types.Generic] = marshal
	define.Unmarshallers[types.Generic] = unmarshal

	stdio.RegesterReadArray(types.Generic, readArray)
	stdio.RegesterReadMap(types.Generic, readMap)
	stdio.RegesterWriteArray(types.Generic, newArrayWriter)
}

var rxWhitespace = regexp.MustCompile(`\s+`)
