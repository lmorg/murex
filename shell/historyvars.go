package shell

import (
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/parameters"
	"regexp"
	"strconv"
	"strings"
)

var (
	rxHistIndex   *regexp.Regexp = regexp.MustCompile(`(\^[0-9]+)`)
	rxHistRegex   *regexp.Regexp = regexp.MustCompile(`\^m/(.*?[^\\])/`) // Scratchpad: https://play.golang.org/p/Iya2Hx1uxb
	rxHistPrefix  *regexp.Regexp = regexp.MustCompile(`(\^\^[a-zA-Z]+)`)
	rxHistTag     *regexp.Regexp = regexp.MustCompile(`(\^#[_a-zA-Z0-9]+)`)
	rxHistAllPs   *regexp.Regexp = regexp.MustCompile(`\^\[([-]?[0-9]+)]\[([-]?[0-9]+)]`)
	rxHistParam   *regexp.Regexp = regexp.MustCompile(`\^\[([-]?[0-9]+)]`)
	rxHistReplace *regexp.Regexp = regexp.MustCompile(`\^s/(.*?[^\\])/(.*?[^\\])/`)
)

func expandHistory(line []rune) []rune {
	s := string(line)

	// Match new line
	s = strings.Replace(s, `^\n`, "\n", -1)
	s = strings.Replace(s, `^\t`, "\t", -1)

	// Match history index
	mhIndex := rxHistIndex.FindAllString(s, -1)
	for i := range mhIndex {
		val, _ := strconv.Atoi(mhIndex[i][1:])
		if val > len(History.List) {
			debug.Log("Value greater than history length.")
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
			debug.Log("Regexp err:", err)
			continue
		}

		for h := len(History.List) - 1; h > -1; h-- {
			if rx.MatchString(History.List[h].Block) {
				s = rxHistRegex.ReplaceAllString(s, noColon(History.List[h].Block))
				return []rune(s)
			}
		}
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
			goto cannotParserxHistAllPs
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
cannotParserxHistAllPs:

	// Match last params (first command in block)
	mhParam = rxHistParam.FindAllStringSubmatch(s, -1)
	if len(mhParam) > 0 {
		nodes, pErr := lang.ParseBlock([]rune(History.Last))
		if pErr.Code != lang.NoParsingErrors {
			goto cannotParserxHistParam
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
cannotParserxHistParam:

	// Match last command
	s = strings.Replace(s, "^!!", noColon(History.Last), -1)

	// Replace string from command buffer
	sList := rxHistReplace.FindAllStringSubmatch(s, -1)
	var rxList []*regexp.Regexp
	var replaceList []string
	//debug.Json("^s/...", sList)
	for i := range sList {
		rx, err := regexp.Compile(sList[i][1])
		if err != nil || len(sList[i]) != 3 {
			debug.Log("Regexp error.", err)
			continue
		}
		rxList = append(rxList, rx)
		replaceList = append(replaceList, sList[i][2])
		s = strings.Replace(s, sList[i][0], "", -1)
	}
	for i := range rxList {
		s = rxList[i].ReplaceAllString(s, replaceList[i])
	}

	// Match history prefix
	mhPrefix := rxHistPrefix.FindAllString(s, -1)
	for i := range mhPrefix {
		for h := len(History.List) - 1; h > -1; h-- {
			if strings.HasPrefix(History.List[h].Block, mhPrefix[i][2:]) {
				s = strings.Replace(s, mhPrefix[i], noColon(History.List[h].Block), 1)
				return []rune(s)
			}
		}
		return []rune(s)
	}

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
