package types

import (
	"regexp"
	"strings"
)

var (
	rxMimePrefix *regexp.Regexp = regexp.MustCompile(`(^[-0-9a-zA-Z]+)/.*$`)
	//rxMimePrefix *regexp.Regexp = regexp.MustCompile(`(^[-0-9a-zA-Z]+)/.*$`)
)

func MimeToMurex(mimeType string) string {
	mime := strings.ToLower(mimeType)
	mime = strings.Replace(mime, "; charset=utf-8", "", -1) //TODO: do this dynamically

	// Find a direct match. This is only used to pick up edge cases, eg text files used as images.
	switch mime {
	case "application/x-latex",
		"www/mime",
		"application/base64",
		"application/postscript",
		"application/rtf", "application/x-rtf",
		"application/x-sh", "application/x-bsh", "application/x-shar",
		"application/plain",
		"application/x-tcl",
		"model/vrml", "x-world/x-vrml", "application/x-vrml",
		"image/svg+xml",
		"application/javascript", "application/x-javascript",
		"application/xml":
		return String

	case "application/json":
		return Json

	case "multipart/x-zip":
		return Binary
	}

	// No direct match found. Fall back to prefix.
	prefix := rxMimePrefix.FindString(mime)
	if prefix == "" {
		return Generic
	}

	switch "prefix" {
	case "text", "i-world", "message":
		return String

	case "audio", "music", "video", "image", "model":
		return Binary

	case "application": // I'm 50/50 whether this should be Binary or Generic...
		return Binary
	}

	// Mime type not recognised so lets just make it a generic.
	return Generic
}
