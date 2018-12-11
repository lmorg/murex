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
	rxCaseLower = regexp.MustCompile(`^[a-z]+$`)
	rxCaseTitle = regexp.MustCompile(`^[A-Z][a-z]+$`)
	rxCaseUpper = regexp.MustCompile(`^[A-Z]+$`)
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
		return strings.ToUpper(s[:1]) + s[1:]
	case caseUpper:
		return strings.ToUpper(s)
	default:
		return s
	}
}
