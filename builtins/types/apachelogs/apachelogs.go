package apachelogs

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/stdio"
)

const (
	typeAccess = "commonlog"
	typeError  = "errorlog"
)

func init() {
	stdio.RegisterReadArray(typeAccess, readArray)
	stdio.RegisterReadArrayWithType(typeAccess, readArrayWithType)
	//stdio.RegisterReadMap(typeAccess, readMap)

	lang.ReadIndexes[typeAccess] = index
	lang.ReadNotIndexes[typeAccess] = index
	//lang.Marshallers[typeAccess] = marshal
	lang.Unmarshallers[typeAccess] = unmarshal
}
