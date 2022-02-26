package lang

import (
	"regexp"
	"strings"

	"github.com/lmorg/murex/lang/types"
)

var rxMimePrefix = regexp.MustCompile(`(^[-0-9a-zA-Z]+)/.*$`)

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

	// Mime type not recognized so lets just make it a generic.
	return types.Generic
}
