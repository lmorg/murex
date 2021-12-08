package mkarray

import (
	"regexp"
	"strings"
)

const (
	caseLower = 1 + iota
	caseLDate
	caseFirst
	caseUpper
	caseTitle
	caseTDate
)

var (
	rxCaseLower = regexp.MustCompile(`^[- _/a-z]+$`)
	rxCaseLDate = regexp.MustCompile(`^[- _/a-z0-9]+$`)
	rxCaseFirst = regexp.MustCompile(`^[A-Z][- _/a-z0-9]+$`)
	rxCaseUpper = regexp.MustCompile(`^[- _/A-Z0-9]+$`)
	rxCaseTDate = regexp.MustCompile(`^[- _/a-zA-Z0-9]+$`)
)

func getCase(s string) int {
	switch {
	case rxCaseLower.MatchString(s):
		return caseLower
	case rxCaseLDate.MatchString(s):
		return caseLDate
	case rxCaseFirst.MatchString(s):
		return caseFirst
	case rxCaseUpper.MatchString(s):
		return caseUpper
	case s == strings.Title(s):
		return caseTitle
	case rxCaseTDate.MatchString(s):
		return caseTDate
	default:
		return 0
	}
}

func setCase(s string, c int) string {
	switch c {
	case caseLower:
		return s
	case caseLDate:
		return strings.ToLower(s)
	case caseFirst:
		return strings.ToUpper(s[:1]) + s[1:]
	case caseUpper:
		return strings.ToUpper(s)
	case caseTitle:
		return strings.Title(s)
	case caseTDate:
		return s
	default:
		return s
	}
}
