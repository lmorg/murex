package lang

import (
	"strings"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

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

// GetExtType gets the murex data type for a corresponding file extension
func GetExtType(extension string) (dt string) {
	dt = fileExts[strings.ToLower(extension)]
	if dt == "" {
		return types.Generic
	}
	return
}

// GetMimes returns MIME lookup table
func GetMimes() map[string]string {
	return mimes
}

// ReadMimes returns a JSON-encoded string.
// This is only intended to be used by `config.Properties.GoFunc.Read()`
func ReadMimes() (interface{}, error) {
	b, err := json.Marshal(mimes, false)
	s := string(b)
	return s, err
}

// GetFileExts returns the file extension lookup table
func GetFileExts() map[string]string {
	return fileExts
}

// ReadFileExtensions returns a JSON-encoded string.
// This is only intended to be used by `config.Properties.GoFunc.Read()`
func ReadFileExtensions() (interface{}, error) {
	b, err := json.Marshal(fileExts, false)
	s := string(b)
	return s, err
}
