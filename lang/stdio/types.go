package stdio

import (
	"context"
	"sort"

	"github.com/lmorg/murex/config"
)

// readArray is where custom data formats can define how to iterate through arrays (eg `foreach`).
// This should only be read from by stream.Io interfaces and written to inside an init() function.
var readArray = make(map[string]func(ctx context.Context, read Io, callback func([]byte)) error)

// readArrayWithType is where custom data formats can define how to iterate through arrays (eg `foreach`).
// This should only be read from by stream.Io interfaces and written to inside an init() function.
var readArrayWithType = make(map[string]func(ctx context.Context, read Io, callback func(interface{}, string)) error)

// ReadMap is where custom data formats can define how to iterate through structured data (eg `formap`).
// This should only be read from by stream.Io interfaces and written to inside an init() function.
var readMap = make(map[string]func(read Io, config *config.Config, callback func(*Map)) error)

// WriteArray is where custom data formats can define how to do buffered writes
var writeArray = make(map[string]func(read Io) (ArrayWriter, error))

// DumpReadArray returns an array of compiled builtins supporting deserialization as an Array
func DumpReadArray() (dump []string) {
	for name := range readArray {
		dump = append(dump, name)
	}
	sort.Strings(dump)
	return
}

// DumpReadArrayWithType returns an array of compiled builtins supporting deserialization as an ArrayWithType
func DumpReadArrayWithType() (dump []string) {
	for name := range readArrayWithType {
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

// DumpWriteArray returns an array of compiled builtins supporting serialization as an Array
func DumpWriteArray() (dump []string) {
	for name := range writeArray {
		dump = append(dump, name)
	}
	sort.Strings(dump)
	return
}
