package apachelogs

import (
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types/define"
)

const (
	typeAccess = "commonlog"
	typeError  = "errorlog"
)

func init() {
	streams.ReadArray[typeAccess] = readArray
	//streams.ReadMap[typeAccess] = readMap
	define.ReadIndexes[typeAccess] = index
	define.ReadNotIndexes[typeAccess] = index
	//define.Marshallers[typeAccess] = marshal
	define.Unmarshallers[typeAccess] = unmarshal
}
