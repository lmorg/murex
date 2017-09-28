package define

import (
	"github.com/lmorg/murex/lang/proc"
	"regexp"
	"strings"
)

var (
	// ReadIndexes defines the Go functions for the `[ Index ]` murex function
	ReadIndexes map[string]func(p *proc.Process, params []string) error = make(map[string]func(*proc.Process, []string) error)

	// Unmarshallers defines the Go functions for converting a murex data type into a Go interface
	Unmarshallers map[string]func(p *proc.Process) (interface{}, error) = make(map[string]func(*proc.Process) (interface{}, error))

	// Marshallers defines the Go functions for converting a Go interface into a murex data type
	Marshallers map[string]func(p *proc.Process, v interface{}) ([]byte, error) = make(map[string]func(*proc.Process, interface{}) ([]byte, error))
)

var (
	mimes        map[string]string = make(map[string]string)
	fileExts     map[string]string = make(map[string]string)
	rxMimePrefix *regexp.Regexp    = regexp.MustCompile(`(^[-0-9a-zA-Z]+)/.*$`)
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
