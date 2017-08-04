package data

import (
	"github.com/lmorg/murex/lang/types"
	"regexp"
	"strings"
)

var rxMimePrefix *regexp.Regexp = regexp.MustCompile(`(^[-0-9a-zA-Z]+)/.*$`)

func MimeToMurex(mimeType string) string {
	mime := strings.ToLower(mimeType)
	mime = strings.Replace(mime, "; charset=utf-8", "", -1) //TODO: do this dynamically

	// Find a direct match. This is only used to pick up edge cases, eg text files used as images.
	dt := Mimes[mime]
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
