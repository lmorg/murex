package apachelogs

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
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
	//lang.RegisterMarshaller(typeAccess,marshal)
	lang.RegisterUnmarshaller(typeAccess, unmarshal)
}
