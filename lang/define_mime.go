package lang

import (
	"path"
	"regexp"
	"strings"

	"github.com/lmorg/murex/lang/types"
)

var rxMimePrefix = regexp.MustCompile(`^([-0-9a-zA-Z]+)/.*$`)

// MimeToMurex gets the murex data type for a corresponding MIME
func MimeToMurex(mimeType string) string {
	mime := NormalizeMime(mimeType)

	// Check suffix
	for suffix := range mimeSuffixes {
		if strings.HasSuffix(mimeType, suffix) {
			return mimeSuffixes[suffix]
		}
	}

	// Find a direct match
	dt := mimes[mime]
	if dt != "" {
		return dt
	}

	// No matches found. Fall back to prefix
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
		return types.Generic

	default:
		// Mime type not recognized so lets just make it a generic
		return types.Generic
	}
}

// RequestToMurex attempts to pick a data type for the content of
// a downloaded file using the value of the Content-Type header
// or, if not, the filename. If neither can be used to infer a
// usable data type, types.Generic is returned.
func RequestMetadataToMurex(contentType string, filename string) string {
	contentType = NormalizeMime(contentType)

	if !(contentType == "" || contentType == "text/plain") {
		foundType := MimeToMurex(contentType)

		// MimeToMurex will return Generic when it can't find
		// any other matches, but in that case we still want
		// to check the filename ext if possible
		if foundType != types.Generic {
			return foundType
		}
	}

	if filename != "" {
		ext := strings.ToLower(strings.TrimPrefix(path.Ext(filename), "."))
		knownType, foundType := GetFileExts()[ext]

		if foundType {
			return knownType
		}
	}

	return types.Generic
}

// NormalizeMime prepares a mime type for processing by removing charset
// information, trimming leading/trailing spaces, and making it lower case
func NormalizeMime(rawMimeType string) string {
	mime := strings.Split(rawMimeType, ";")[0]
	mime = strings.TrimSpace(mime)
	mime = strings.ToLower(mime)

	return mime
}

// MurexToMime returns the default MIME for a given Murex data type.
// The intended use case for this is for GET and POST requests where the body
// is STDIN.
func MurexToMime(dataType string) string {
	return defaultMimes[dataType]
}
