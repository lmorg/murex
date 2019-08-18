package define

import (
	"regexp"
	"strings"

	"github.com/lmorg/murex/lang"
)

var (
	// ReadIndexes defines the Go functions for the `[ Index ]` murex function
	ReadIndexes = make(map[string]func(*lang.Process, []string) error)

	// ReadNotIndexes defines the Go functions for the `![ Index ]` murex function
	ReadNotIndexes = make(map[string]func(*lang.Process, []string) error)

	// Unmarshallers defines the Go functions for converting a murex data type into a Go interface
	Unmarshallers = make(map[string]func(*lang.Process) (interface{}, error))

	// Marshallers defines the Go functions for converting a Go interface into a murex data type
	Marshallers = make(map[string]func(*lang.Process, interface{}) ([]byte, error))
)

var (
	mimes        = make(map[string]string)
	fileExts     = make(map[string]string)
	rxMimePrefix = regexp.MustCompile(`(^[-0-9a-zA-Z]+)/.*$`)
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
