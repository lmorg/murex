package data

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"regexp"
	"strings"
)

var (
	ReadIndexes  map[string]func(p *proc.Process, params []string) error         = make(map[string]func(*proc.Process, []string) error)
	Unmarshal    map[string]func(p *proc.Process) (interface{}, error)           = make(map[string]func(*proc.Process) (interface{}, error))
	Marshal      map[string]func(p *proc.Process, v interface{}) ([]byte, error) = make(map[string]func(*proc.Process, interface{}) ([]byte, error))
	mimes        map[string]string                                               = make(map[string]string)
	fileExts     map[string]string                                               = make(map[string]string)
	rxMimePrefix *regexp.Regexp                                                  = regexp.MustCompile(`(^[-0-9a-zA-Z]+)/.*$`)
)

// Define MIME(s) and assign it a murex data type
func SetMime(dt string, mime ...string) {
	for i := range mime {
		mimes[mime[i]] = dt
	}
}

// Get the murex data type for a corresponding MIME
func MimeToMurex(mimeType string) string {
	mime := strings.ToLower(mimeType)
	mime = strings.Replace(mime, "; charset=utf-8", "", -1) //TODO: do this dynamically

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

// Define file extension(s) and assign it a murex data type
func SetFileExtensions(dt string, extension ...string) {
	for i := range extension {
		fileExts[extension[i]] = strings.ToLower(dt)
	}
}

// Get the murex data type for a corresponding file extension
func GetExtType(extension string) (dt string) {
	dt = fileExts[strings.ToLower(extension)]
	return
}
