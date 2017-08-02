package streams

import (
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/types"
)

var ReadArray map[string]func(read Io, callback func([]byte)) error = make(map[string]func(read Io, callback func([]byte)) error)
var ReadMap map[string]func(read Io, config *config.Config, callback func(key, value string, last bool)) error = make(map[string]func(read Io, config *config.Config, callback func(key, value string, last bool)) error)

func init() {
	// ReadArray
	ReadArray[types.Generic] = readArrayDefault
	ReadArray[types.String] = readArrayDefault
	ReadArray[types.Json] = readArrayJson

	// ReadMap
	ReadMap[types.Generic] = readMapDefault
	ReadMap[types.String] = readMapDefault
	ReadMap[types.Json] = readMapJson
	ReadMap[types.Csv] = readMapCsv
}

func readArray(read Io, callback func([]byte)) error {
	dt := read.GetDataType()

	if ReadArray[dt] != nil {
		return ReadArray[dt](read, callback)
	}

	return ReadArray[types.Generic](read, callback)
}

func readMap(read Io, config *config.Config, callback func(key, value string, last bool)) error {
	dt := read.GetDataType()

	if ReadMap[dt] != nil {
		return ReadMap[dt](read, config, callback)
	}

	return ReadMap[types.Generic](read, config, callback)
}

func ListArrays() (s []string) {
	for name := range ReadArray {
		s = append(s, name)
	}
	return
}

func ListMaps() (s []string) {
	for name := range ReadMap {
		s = append(s, name)
	}
	return
}
