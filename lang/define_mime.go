package lang

import (
	"regexp"
	"strings"

	"github.com/lmorg/murex/lang/types"
)

var rxMimePrefix = regexp.MustCompile(`^([-0-9a-zA-Z]+)/.*$`)

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
	prefix := rxMimePrefix.FindStringSubmatch(mime)
	if len(prefix) != 2 {
		return types.Generic
	}

	switch prefix[1] {
	case "text", "i-world", "message":
		return types.String

	case "audio", "music", "video", "image", "model":
		return types.Binary

	case "application":
		if strings.HasSuffix(mime, "+json") {
			return types.Json
		}
		return types.Generic

	default:
		// Mime type not recognized so lets just make it a generic.
		return types.Generic
	}

}
