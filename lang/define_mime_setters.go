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

	MxInterfaces = make(map[string]MxInterface)
)

var (
	mimes        = make(map[string]string)
	mimeSuffixes = make(map[string]string)
	fileExts     = make(map[string]string)
	defaultMimes = make(map[string]string)
)

// SetMime defines MIME(s) and assign it a murex data type.
// The default MIME for any data type should be the first MIME passed.
func SetMime(dataType string, mime ...string) {
	defaultMimes[dataType] = mime[0] // default should always be first

	for i := range mime {
		switch {
		case len(mime[i]) == 0:
			continue

		case mime[i][0] == '+':
			mimeSuffixes[mime[i]] = dataType

		default:
			mimes[mime[i]] = dataType
		}
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
func WriteMimes(v any) error {
	switch v := v.(type) {
	case string:
		newMimes := make(map[string]string)
		err := json.Unmarshal([]byte(v), &newMimes)
		if err != nil {
			return err
		}
		mimes = newMimes
		return nil

	default:
		return fmt.Errorf("invalid data-type expecting a %s encoded string", types.Json)
	}
}

// WriteFileExtensions takes a JSON-encoded string and writes it to the
// fileExts map.
// This is only intended to be used by `config.Properties.GoFunc.Write()`
func WriteFileExtensions(v any) error {
	switch v := v.(type) {
	case string:
		newFileExts := make(map[string]string)
		err := json.Unmarshal([]byte(v), &newFileExts)
		if err != nil {
			return err
		}
		fileExts = newFileExts
		return nil

	default:
		return fmt.Errorf("invalid data-type. Expecting a %s encoded string", types.Json)
	}
}

// WriteDefaultMimes takes a JSON-encoded string and writes it to the default
// MIMEs map.
// This is only intended to be used by `config.Properties.GoFunc.Write()`
func WriteDefaultMimes(v any) error {
	switch v := v.(type) {
	case string:
		newDefaultMimes := make(map[string]string)
		err := json.Unmarshal([]byte(v), &newDefaultMimes)
		if err != nil {
			return err
		}
		defaultMimes = newDefaultMimes
		return nil

	default:
		return fmt.Errorf("invalid data-type. Expecting a %s encoded string", types.Json)
	}
}
