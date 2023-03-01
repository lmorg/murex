package paths

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
)

const (
	typePath  = "path"
	typePaths = "paths"
)

func init() {
	// Register data type
	lang.Marshallers[typePath] = marshalPath
	lang.Marshallers[typePaths] = marshalPaths
	lang.Unmarshallers[typePath] = unmarshalPath
	lang.Unmarshallers[typePaths] = unmarshalPaths
	lang.ReadIndexes[typePath] = indexPath
	lang.ReadIndexes[typePaths] = indexPaths
	lang.ReadNotIndexes[typePath] = indexPath
	lang.ReadNotIndexes[typePaths] = indexPaths

	stdio.RegisterReadArray(typePath, readArrayPath)
	stdio.RegisterReadArray(typePaths, readArrayPaths)
	stdio.RegisterReadArrayWithType(typePath, readArrayWithTypePath)
	stdio.RegisterReadArrayWithType(typePaths, readArrayWithTypePaths)
	//stdio.RegisterReadMap(types.Json, readMap)
	//stdio.RegisterWriteArray(types.Json, newArrayWriter)
}
