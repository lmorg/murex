package stdio

import (
	"context"
	"fmt"
	"sort"

	"github.com/lmorg/murex/config"
)

var pipes = make(map[string]func(string) (Io, error))

// RegisterPipe is used by pipes (/builtins/) to regester themselves to murex.
// This function should only be called from a packages Init() func.
func RegisterPipe(name string, constructor func(string) (Io, error)) {
	if pipes[name] != nil {
		panic("Pipe already registered with the name: " + name)
	}

	pipes[name] = constructor
}

// CreatePipe returns an stdio.Io interface for a specified pipe type or errors if
// the pipe type is invalid.
func CreatePipe(pipeType, arguments string) (Io, error) {
	if pipes[pipeType] == nil {
		return nil, fmt.Errorf("`%s` is not a supported pipe type", pipeType)
	}

	return pipes[pipeType](arguments)
}

// DumpPipes returns a sorted array of regestered pipes.
func DumpPipes() (dump []string) {
	for name := range pipes {
		dump = append(dump, name)
	}

	sort.Strings(dump)
	return
}

// RegisterReadArray is used by types (/builtins/types) to regester themselves to murex.
// This function should only be called from a packages Init() func.
func RegisterReadArray(dataType string, function func(ctx context.Context, read Io, callback func([]byte)) error) {
	if readArray[dataType] != nil {
		panic("readArray already registered for the data type: " + dataType)
	}

	readArray[dataType] = function
}

// RegisterReadArrayWithType is used by types (/builtins/types) to regester themselves to murex.
// This function should only be called from a packages Init() func.
func RegisterReadArrayWithType(dataType string, function func(ctx context.Context, read Io, callback func([]byte, string)) error) {
	if readArrayWithType[dataType] != nil {
		panic("readArrayWithType already registered for the data type: " + dataType)
	}

	readArrayWithType[dataType] = function
}

// RegisterReadMap is used by pipes (/builtins/) to regester themselves to murex.
// This function should only be called from a packages Init() func.
func RegisterReadMap(dataType string, function func(read Io, config *config.Config, callback func(key, value string, last bool)) error) {
	if readMap[dataType] != nil {
		panic("readMap already registered for the data type: " + dataType)
	}

	readMap[dataType] = function
}

// RegisterWriteArray is used by pipes (/builtins/) to regester themselves to murex.
// This function should only be called from a packages Init() func.
func RegisterWriteArray(dataType string, function func(read Io) (ArrayWriter, error)) {
	if writeArray[dataType] != nil {
		panic("writeArray already registered for the data type: " + dataType)
	}

	writeArray[dataType] = function
}
