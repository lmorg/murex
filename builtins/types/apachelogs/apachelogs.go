package apachelogs

import (
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types/define"
)

const (
	typeAccess = "commonlog"
	typeError  = "errorlog"
)

func init() {
	stdio.RegesterReadArray(typeAccess, readArray)
	//stdio.RegesterReadMap(typeAccess, readMap)

	define.ReadIndexes[typeAccess] = index
	define.ReadNotIndexes[typeAccess] = index
	//define.Marshallers[typeAccess] = marshal
	define.Unmarshallers[typeAccess] = unmarshal
}
