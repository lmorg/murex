package streams

import (
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc/streams/stdio"
	"github.com/lmorg/murex/lang/types"
	"sort"
)

// ReadArray is where custom data formats can define how to iterate through arrays (eg `foreach`).
// This should only be read from by stream.Io interfaces and written to inside an init() function.
var ReadArray map[string]func(read stdio.Io, callback func([]byte)) error = make(map[string]func(read stdio.Io, callback func([]byte)) error)

// ReadMap is where custom data formats can define how to iterate through structured data (eg `formap`).
// This should only be read from by stream.Io interfaces and written to inside an init() function.
var ReadMap map[string]func(read stdio.Io, config *config.Config, callback func(key, value string, last bool)) error = make(map[string]func(read stdio.Io, config *config.Config, callback func(key, value string, last bool)) error)

func readArray(read stdio.Io, callback func([]byte)) error {
	dt := read.GetDataType()

	if ReadArray[dt] != nil {
		return ReadArray[dt](read, callback)
	}

	return ReadArray[types.Generic](read, callback)
}

func readMap(read stdio.Io, config *config.Config, callback func(key, value string, last bool)) error {
	dt := read.GetDataType()

	if ReadMap[dt] != nil {
		return ReadMap[dt](read, config, callback)
	}

	return ReadMap[types.Generic](read, config, callback)
}

// DumpArray returns an array of compiled builtins supporting deserialization as an array
func DumpArray() (dump []string) {
	for name := range ReadArray {
		dump = append(dump, name)
	}
	sort.Strings(dump)
	return
}

// DumpMap returns an array of compiled builtins supporting deserialization as a key/value map (or hash)
func DumpMap() (dump []string) {
	for name := range ReadMap {
		dump = append(dump, name)
	}
	sort.Strings(dump)
	return
}
