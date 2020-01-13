package lang

import (
	"strings"

	"github.com/lmorg/murex/lang/types"
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

// ReadMimes returns an interface{} of mimes.
// This is only intended to be used by `config.Properties.GoFunc.Read()`
func ReadMimes() (interface{}, error) {
	return mimes, nil
}

// GetFileExts returns the file extension lookup table
func GetFileExts() map[string]string {
	return fileExts
}

// ReadFileExtensions returns an interface{} of fileExts.
// This is only intended to be used by `config.Properties.GoFunc.Read()`
func ReadFileExtensions() (interface{}, error) {
	return fileExts, nil
}
