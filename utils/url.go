package utils

import "regexp"

var rxUrl *regexp.Regexp = regexp.MustCompile(`(?i)^http(s)?://`)

// IsURL checks the start of a string to see if it has a valid HTTP protocol
func IsURL(url string) bool {
	return rxUrl.MatchString(url)
}
