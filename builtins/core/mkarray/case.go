package mkarray

import (
	"regexp"
	"strings"
)

const (
	caseLower = 1 + iota
	caseFirst
	caseUpper
	caseTitle
)

var (
	rxCaseLower = regexp.MustCompile(`^[- _a-z]+$`)
	rxCaseFirst = regexp.MustCompile(`^[A-Z][- _a-z]+$`)
	rxCaseUpper = regexp.MustCompile(`^[- _A-Z]+$`)
)

func getCase(s string) int {
	switch {
	case rxCaseLower.MatchString(s):
		return caseLower
	case rxCaseFirst.MatchString(s):
		return caseFirst
	case rxCaseUpper.MatchString(s):
		return caseUpper
	case s == strings.Title(s):
		return caseTitle
	default:
		return 0
	}
}

func setCase(s string, c int) string {
	switch c {
	case caseLower:
		return s
	case caseFirst:
		return strings.ToUpper(s[:1]) + s[1:]
	case caseUpper:
		return strings.ToUpper(s)
	case caseTitle:
		return strings.Title(s)
	default:
		return s
	}
}
