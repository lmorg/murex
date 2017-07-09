package mkarray

import (
	"regexp"
	"strings"
)

const (
	caseLower = 1 + iota
	caseTitle
	caseUpper
)

var (
	rxCaseLower *regexp.Regexp = regexp.MustCompile(`^[a-z]+$`)
	rxCaseTitle *regexp.Regexp = regexp.MustCompile(`^[A-Z][a-z]+$`)
	rxCaseUpper *regexp.Regexp = regexp.MustCompile(`^[A-Z]+$`)
)

func getCase(s string) int {
	switch {
	case rxCaseLower.MatchString(s):
		return caseLower
	case rxCaseTitle.MatchString(s):
		return caseTitle
	case rxCaseUpper.MatchString(s):
		return caseUpper
	default:
		return 0
	}
}

func setCase(s string, c int) string {
	switch c {
	case caseLower:
		return s
	case caseTitle:
		return strings.ToTitle(s)
	case caseUpper:
		return strings.ToUpper(s)
	default:
		return s
	}
}
