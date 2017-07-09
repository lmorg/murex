package mkarray

import (
	"regexp"
	"strings"
)

const (
	caseLower = 1 + iota
	caseTitle
	CaseUpper
)

var (
	rxCaseLower *regexp.Regexp = regexp.MustCompile(`^[a-z]+$`)
	rxCaseTitle *regexp.Regexp = regexp.MustCompile(`^[A-Z]{1}[a-z]+$`)
	rxCaseUpper *regexp.Regexp = regexp.MustCompile(`^[A-Z]+$`)
)

func getCase(s string) int {
	switch {
	case rxCaseLower.MatchString(s):
		return caseLower
	case rxCaseTitle.MatchString(s):
		return caseTitle
	case rxCaseUpper.MatchString(s):
		return CaseUpper
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
	case CaseUpper:
		return strings.ToUpper(s)
	default:
		return s
	}
}
