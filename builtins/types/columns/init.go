package columns

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	// Register data type
	//lang.ReadIndexes[types.Columns] = index
	//lang.ReadNotIndexes[types.Columns] = index
	lang.Marshallers[types.Columns] = marshal
	//lang.Unmarshallers[types.Columns] = unmarshal

	//stdio.RegisterReadArray(types.Columns, readArray)
	//stdio.RegisterReadMap(types.Columns, readMap)
	//stdio.RegisterWriteArray(types.Columns, newArrayWriter)
}

//var rxWhitespace = regexp.MustCompile(`\s+`)
