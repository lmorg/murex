package shell

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	rxHistIndex *regexp.Regexp = regexp.MustCompile(`(\^[0-9]+)`)
	rxHistRegex *regexp.Regexp = regexp.MustCompile(`\^m/(.*?[^\\])/`) // Scratchpad: https://play.golang.org/p/Iya2Hx1uxb
)

func expandHistory(line []rune) []rune {
	s := string(line)
	matchHI := rxHistIndex.FindAllString(s, -1)
	for i := range matchHI {
		val, _ := strconv.Atoi(matchHI[i][1:])
		if val > len(History.List) {
			continue
		}
		s = rxHistIndex.ReplaceAllString(s, noColon(History.List[val].Block))
		return []rune(s)
	}

	matchHR := rxHistRegex.FindAllStringSubmatch(s, -1)
	for i := range matchHR {
		rx, err := regexp.Compile(matchHR[i][1])
		if err != nil {
			continue
		}

		for h := len(History.List) - 1; h > -1; h-- {
			if rx.MatchString(History.List[h].Block) {
				s = rxHistRegex.ReplaceAllString(s, noColon(History.List[h].Block))
				return []rune(s)
			}
		}
	}

	s = strings.Replace(s, "^!!", noColon(History.Last), -1)

	return []rune(s)
}

func noColon(line string) string {
	var escape, qSingle, qDouble bool

	for i := range line {
		switch line[i] {
		case '#':
			return line
		case '\\':
			switch {
			case escape:
				escape = false
			case qSingle:
			// do nothing
			default:
				escape = true
			}
		case '\'':
			switch {
			case qDouble, escape:
				escape = false
			default:
				qSingle = !qSingle
			}
		case '"':
			switch {
			case qSingle, escape:
				escape = false
			default:
				qDouble = !qDouble
			}
		case '{':
			if !escape && !qSingle && !qDouble {
				return line
			}
		case '\r', '\n', '\t', ' ':
			if !escape && !qSingle && !qDouble {
				return line
			}
		case ':':
			if !escape && !qSingle && !qDouble {
				return line[:i] + line[i+1:]
			}
		}
	}

	return line
}
