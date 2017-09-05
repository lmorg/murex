package utils

import "regexp"

var rxUrl *regexp.Regexp = regexp.MustCompile(`(?i)^http(s)://`)

func IsURL(url string) bool {
	return rxUrl.MatchString(url)
}
