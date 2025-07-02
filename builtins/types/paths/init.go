package paths

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	// Register data type
	lang.RegisterMarshaller(types.Path, marshalPath)
	lang.RegisterMarshaller(types.Paths, marshalPaths)
	lang.RegisterUnmarshaller(types.Path, unmarshalPath)
	lang.RegisterUnmarshaller(types.Paths, unmarshalPaths)
	lang.ReadIndexes[types.Path] = indexPath
	lang.ReadIndexes[types.Paths] = indexPaths
	lang.ReadNotIndexes[types.Path] = indexPath
	lang.ReadNotIndexes[types.Paths] = indexPaths

	stdio.RegisterReadArray(types.Path, readArrayPath)
	stdio.RegisterReadArray(types.Paths, readArrayPaths)
	stdio.RegisterReadArrayWithType(types.Path, readArrayWithTypePath)
	stdio.RegisterReadArrayWithType(types.Paths, readArrayWithTypePaths)
	//stdio.RegisterReadMap(types.Json, readMap)
	stdio.RegisterWriteArray(types.Path, newArrayWriterPath)
	stdio.RegisterWriteArray(types.Paths, newArrayWriterPaths)
}
