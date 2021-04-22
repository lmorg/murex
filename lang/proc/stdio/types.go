package stdio

import (
	"sort"

	"github.com/lmorg/murex/config"
)

// readArray is where custom data formats can define how to iterate through arrays (eg `foreach`).
// This should only be read from by stream.Io interfaces and written to inside an init() function.
var readArray = make(map[string]func(read Io, callback func([]byte)) error)

// readArrayByType is where custom data formats can define how to iterate through arrays (eg `foreach`).
// This should only be read from by stream.Io interfaces and written to inside an init() function.
var readArrayByType = make(map[string]func(read Io, callback func([]byte, string)) error)

// ReadMap is where custom data formats can define how to iterate through structured data (eg `formap`).
// This should only be read from by stream.Io interfaces and written to inside an init() function.
var readMap = make(map[string]func(read Io, config *config.Config, callback func(key, value string, last bool)) error)

// WriteArray is where custom data formats can define how to do buffered writes
var writeArray = make(map[string]func(read Io) (ArrayWriter, error))

// DumpArray returns an array of compiled builtins supporting deserialization as an array
func DumpArray() (dump []string) {
	for name := range readArray {
		dump = append(dump, name)
	}
	sort.Strings(dump)
	return
}

// DumpMap returns an array of compiled builtins supporting deserialization as a key/value map (or hash)
func DumpMap() (dump []string) {
	for name := range readMap {
		dump = append(dump, name)
	}
	sort.Strings(dump)
	return
}
