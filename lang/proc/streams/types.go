package streams

import (
	"sort"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc/streams/stdio"
)

// ReadArray is where custom data formats can define how to iterate through arrays (eg `foreach`).
// This should only be read from by stream.Io interfaces and written to inside an init() function.
var ReadArray = make(map[string]func(read stdio.Io, callback func([]byte)) error)

// ReadMap is where custom data formats can define how to iterate through structured data (eg `formap`).
// This should only be read from by stream.Io interfaces and written to inside an init() function.
var ReadMap = make(map[string]func(read stdio.Io, config *config.Config, callback func(key, value string, last bool)) error)

// WriteArray is where custom data formats can define how to do buffered writes
var WriteArray = make(map[string]func(read stdio.Io) (stdio.ArrayWriter, error))

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
