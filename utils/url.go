package utils

import (
	"net/url"
	"regexp"
	"strings"
)

var rxUrl *regexp.Regexp = regexp.MustCompile(`(?i)^http(s)?://`)

// IsURL checks the start of a string to see if it has a valid HTTP protocol
func IsURL(url string) bool {
	return rxUrl.MatchString(url)
}

// ExtractFileName returns the file name from a URL
func ExtractFileNameFromURL(address string) string {
	u, err := url.Parse(address)
	if err != nil {
		return address
	}

	if len(u.Path) == 0 || u.Path == "/" {
		return u.Host
	}

	split := strings.Split(u.Path, "/")
	for i := len(split) - 1; i > -1; i-- {
		if len(split[i]) != 0 && split[i] != "/" {
			return split[i]
		}
	}

	return u.Path
}
