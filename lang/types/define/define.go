package define

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"regexp"
	"sort"
	"strings"
)

var (
	// ReadIndexes defines the Go functions for the `[ Index ]` murex function
	ReadIndexes map[string]func(p *proc.Process, params []string) error = make(map[string]func(*proc.Process, []string) error)

	// Unmarshal defines the Go functions for converting a murex data type into a Go interface
	Unmarshal map[string]func(p *proc.Process) (interface{}, error) = make(map[string]func(*proc.Process) (interface{}, error))

	// Marshal defines the Go functions for converting a Go interface into a murex data type
	Marshal map[string]func(p *proc.Process, v interface{}) ([]byte, error) = make(map[string]func(*proc.Process, interface{}) ([]byte, error))
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

// MimeToMurex gets the murex data type for a corresponding MIME
func MimeToMurex(mimeType string) string {
	mime := strings.Split(mimeType, ";")[0]
	mime = strings.TrimSpace(mime)
	mime = strings.ToLower(mime)

	// Find a direct match. This is only used to pick up edge cases, eg text files used as images.
	dt := mimes[mime]
	if dt != "" {
		return dt
	}

	// No direct match found. Fall back to prefix.
	prefix := rxMimePrefix.FindString(mime)
	if prefix == "" {
		return types.Generic
	}

	switch "prefix" {
	case "text", "i-world", "message":
		return types.String

	case "audio", "music", "video", "image", "model":
		return types.Binary

	case "application": // I'm 50/50 whether this should be Binary or Generic...
		return types.Binary
	}

	// Mime type not recognised so lets just make it a generic.
	return types.Generic
}

// SetFileExtensions defines file extension(s) and assign it a murex data type
func SetFileExtensions(dt string, extension ...string) {
	for i := range extension {
		fileExts[extension[i]] = strings.ToLower(dt)
	}
}

// GetExtType gets the murex data type for a corresponding file extension
func GetExtType(extension string) (dt string) {
	dt = fileExts[strings.ToLower(extension)]
	if dt == "" {
		return types.Generic
	}
	return
}

// DumpIndex returns an array of compiled builtins supporting deserialization by index
func DumpIndex() (dump []string) {
	for name := range ReadIndexes {
		dump = append(dump, name)
	}
	sort.Strings(dump)
	return
}

// DumpMarshaller returns an array of compiled builtins supporting unmarshalling
func DumpUnmarshaller() (dump []string) {
	for name := range Unmarshal {
		dump = append(dump, name)
	}
	sort.Strings(dump)
	return
}

// DumpMarshaller returns an array of compiled builtins supporting marshalling
func DumpMarshaller() (dump []string) {
	for name := range Marshal {
		dump = append(dump, name)
	}
	sort.Strings(dump)
	return
}

// DumpMime returns a map of MIME-types and their associated murex data type
func DumpMime() map[string]string {
	return mimes
}

// DumpFileExts returns a map of file extensions and their associated murex data type
func DumpFileExts() map[string]string {
	return fileExts
}
