package shell

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/parameters"
	"regexp"
	"strconv"
	"strings"
)

var (
	rxHistIndex  *regexp.Regexp = regexp.MustCompile(`(\^[0-9]+)`)
	rxHistRegex  *regexp.Regexp = regexp.MustCompile(`\^m/(.*?[^\\])/`) // Scratchpad: https://play.golang.org/p/Iya2Hx1uxb
	rxHistPrefix *regexp.Regexp = regexp.MustCompile(`(\^[a-zA-Z]+)`)
	rxHistTag    *regexp.Regexp = regexp.MustCompile(`(\^#[_a-zA-Z0-9]+)`)
	rxHistAllPs  *regexp.Regexp = regexp.MustCompile(`\^\[([-]?[0-9]+)]\[([-]?[0-9]+)]`)
	rxHistParam  *regexp.Regexp = regexp.MustCompile(`\^\[([-]?[0-9]+)]`)
)

func expandHistory(line []rune) []rune {
	s := string(line)

	// Match history index
	mhIndex := rxHistIndex.FindAllString(s, -1)
	for i := range mhIndex {
		val, _ := strconv.Atoi(mhIndex[i][1:])
		if val > len(History.List) {
			continue
		}
		s = rxHistIndex.ReplaceAllString(s, noColon(History.List[val].Block))
		return []rune(s)
	}

	// Match history by regexp
	mhRegexp := rxHistRegex.FindAllStringSubmatch(s, -1)
	for i := range mhRegexp {
		rx, err := regexp.Compile(mhRegexp[i][1])
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

	// Match history prefix
	mhPrefix := rxHistPrefix.FindAllString(s, -1)
	for i := range mhPrefix {
		for h := len(History.List) - 1; h > -1; h-- {
			if strings.HasPrefix(History.List[h].Block, mhPrefix[i][1:]) {
				s = strings.Replace(s, mhPrefix[i], noColon(History.List[h].Block), 1)
				return []rune(s)
			}
		}
		return []rune(s)
	}

	// Match history hashtag
	mhTag := rxHistTag.FindAllString(s, -1)
	for i := range mhTag {
		for h := len(History.List) - 1; h > -1; h-- {
			if strings.HasSuffix(History.List[h].Block, mhTag[i][1:]) {
				s = strings.Replace(s, mhTag[i], noColon(History.List[h].Block), 1)

				return []rune(s)
			}
		}
		return []rune(s)
	}

	// Match last params (all of block)
	mhParam := rxHistAllPs.FindAllStringSubmatch(s, -1)
	if len(mhParam) > 0 {
		nodes, pErr := lang.ParseBlock([]rune(History.Last))
		if pErr.Code != lang.NoParsingErrors {
			goto cannotParseLast
		}

		for i := range mhParam {
			cmd, _ := strconv.Atoi(mhParam[i][1])
			if cmd < 0 {
				cmd += len(nodes) + 1
			}
			val, _ := strconv.Atoi(mhParam[i][2])

			if cmd < 0 || cmd+1 > len(nodes) {
				continue
			}

			p := parameters.Parameters{Tokens: nodes[cmd].ParamTokens}
			lang.ParseParameters(proc.ShellProcess, &p, &proc.GlobalVars)
			if val < 0 {
				val += p.Len() + 1
			}

			if val == 0 {
				s = strings.Replace(s, mhParam[i][0], nodes[cmd].Name, -1)
			} else if val > 0 && val-1 < p.Len() {
				s = strings.Replace(s, mhParam[i][0], p.Params[val-1], -1)
			}

		}

		//return []rune(s)
	}

	// Match last params (first command in block)
	mhParam = rxHistParam.FindAllStringSubmatch(s, -1)
	if len(mhParam) > 0 {
		nodes, pErr := lang.ParseBlock([]rune(History.Last))
		if pErr.Code != lang.NoParsingErrors {
			goto cannotParseLast
		}
		p := parameters.Parameters{Tokens: nodes.Last().ParamTokens}
		lang.ParseParameters(proc.ShellProcess, &p, &proc.GlobalVars)

		for i := range mhParam {
			val, _ := strconv.Atoi(mhParam[i][1])
			if val < 0 {
				val += p.Len() + 1
			}

			if val == 0 {
				s = strings.Replace(s, mhParam[i][0], nodes.Last().Name, -1)
			} else if val > 0 && val-1 < p.Len() {
				s = strings.Replace(s, mhParam[i][0], p.Params[val-1], -1)
			}

		}

		return []rune(s)
	}
cannotParseLast:

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
