package lang

import (
	"fmt"
	"strings"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

var (
	// ReadIndexes defines the Go functions for the `[ Index ]` murex function
	ReadIndexes = make(map[string]func(*Process, []string) error)

	// ReadNotIndexes defines the Go functions for the `![ Index ]` murex function
	ReadNotIndexes = make(map[string]func(*Process, []string) error)

	// Unmarshallers defines the Go functions for converting a murex data type into a Go interface
	Unmarshallers = make(map[string]func(*Process) (interface{}, error))

	// Marshallers defines the Go functions for converting a Go interface into a murex data type
	Marshallers = make(map[string]func(*Process, interface{}) ([]byte, error))

	MxInterfaces = make(map[string]MxInterface)
)

var (
	mimes    = make(map[string]string)
	fileExts = make(map[string]string)
)

// SetMime defines MIME(s) and assign it a murex data type
func SetMime(dt string, mime ...string) {
	for i := range mime {
		mimes[mime[i]] = dt
	}
}

// SetFileExtensions defines file extension(s) and assign it a murex data type
func SetFileExtensions(dt string, extension ...string) {
	for i := range extension {
		fileExts[extension[i]] = strings.ToLower(dt)
	}
}

// WriteMimes takes a JSON-encoded string and writes it to the mimes map.
// This is only intended to be used by `config.Properties.GoFunc.Write()`
func WriteMimes(v interface{}) error {
	switch v := v.(type) {
	case string:
		mimes = make(map[string]string)
		return json.Unmarshal([]byte(v), &mimes)

	default:
		return fmt.Errorf("invalid data-type. Expecting a %s encoded string", types.Json)
	}
}

// WriteFileExtensions takes a JSON-encoded string and writes it to the
// fileExts map.
// This is only intended to be used by `config.Properties.GoFunc.Write()`
func WriteFileExtensions(v interface{}) error {
	switch v := v.(type) {
	case string:
		fileExts = make(map[string]string)
		return json.Unmarshal([]byte(v), &fileExts)

	default:
		return fmt.Errorf("invalid data-type. Expecting a %s encoded string", types.Json)
	}
}
