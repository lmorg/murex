package history

import (
	"fmt"
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

const (
	errCannotParsePrevCmd = "Cannot parse previous command line to extract parameters for history variable."
)

func ExpandVariables(line []rune, hist *History) ([]rune, error) {
	s := string(line)

	s = strings.Replace(s, `^\n`, "\n", -1)               // Match new line
	s = strings.Replace(s, `^\t`, "\t", -1)               // Match tab
	s = strings.Replace(s, "^!!", noColon(hist.Last), -1) // Match last command

	funcs := []func(string, *History) (string, error){
		expandHistPrefix,
		expandHistIndex,
		expandHistRegex,
		expandHistHashtag,
		expandHistAllPs,
		expandHistParam,
		expandHistReplace,
	}

	for f := range funcs {
		var err error
		s, err = funcs[f](s, hist)
		if err != nil {
			return nil, err
		}
	}

	return []rune(s), nil
}

// Match history index
func expandHistIndex(s string, hist *History) (string, error) {
	mhIndex := rxHistIndex.FindAllString(s, -1)
	for i := range mhIndex {
		val, _ := strconv.Atoi(mhIndex[i][1:])
		if val > len(hist.List) {
			return "", fmt.Errorf("Value greater than history length in ^%d", val)
		}
		s = rxHistIndex.ReplaceAllString(s, noColon(hist.List[val].Block))
		//return s, nil
	}
	return s, nil
}

// Match history by regexp
func expandHistRegex(s string, hist *History) (string, error) {
	mhRegexp := rxHistRegex.FindAllStringSubmatch(s, -1)
	for i := range mhRegexp {
		rx, err := regexp.Compile(mhRegexp[i][1])
		if err != nil {
			return "", fmt.Errorf("Regexp error in history variable `^m/%s/`: %s", mhRegexp[i][1], err.Error())
		}

		for h := len(hist.List) - 1; h > -1; h-- {
			if rx.MatchString(hist.List[h].Block) {
				s = rxHistRegex.ReplaceAllString(s, noColon(hist.List[h].Block))
				//return []rune(s)
				goto next
			}
		}
		return "", fmt.Errorf("Cannot find a history item to match regexp: %s", mhRegexp[i][1])
	next:
	}
	return s, nil
}

// Match history hashtag
func expandHistHashtag(s string, hist *History) (string, error) {
	mhTag := rxHistTag.FindAllString(s, -1)

	for i := range mhTag {
		for h := len(hist.List) - 1; h > -1; h-- {
			if strings.HasSuffix(hist.List[h].Block, mhTag[i][1:]) {
				block := hist.List[h].Block[:len(hist.List[h].Block)-len(mhTag[i][1:])]
				s = strings.Replace(s, mhTag[i], noColon(block), 1)

				//return s, nil
				goto next
			}
		}

		return "", fmt.Errorf("Hashtag not found: %s", mhTag[i])
	next:
	}

	return s, nil
}

// Match last params (all of block)
func expandHistAllPs(s string, hist *History) (string, error) {
	mhParam := rxHistAllPs.FindAllStringSubmatch(s, -1)
	if len(mhParam) > 0 {
		nodes, pErr := lang.ParseBlock([]rune(hist.Last))
		if pErr.Code != lang.NoParsingErrors {
			//goto cannotParserxHistAllPs
			return "", fmt.Errorf(errCannotParsePrevCmd)
		}

		for i := range mhParam {
			cmd, _ := strconv.Atoi(mhParam[i][1])
			if cmd < 0 {
				cmd += len(nodes) + 1
			}
			val, _ := strconv.Atoi(mhParam[i][2])

			if cmd < 0 || cmd+1 > len(nodes) {
				return "", fmt.Errorf("Cannot extract parameter for %s", mhParam[i][0])
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

	return s, nil
}

// Match last params (first command in block)
func expandHistParam(s string, hist *History) (string, error) {
	mhParam := rxHistParam.FindAllStringSubmatch(s, -1)
	if len(mhParam) > 0 {
		nodes, pErr := lang.ParseBlock([]rune(hist.Last))
		if pErr.Code != lang.NoParsingErrors {
			//goto cannotParserxHistParam
			return "", fmt.Errorf(errCannotParsePrevCmd)
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

		//return s, nil //err
	}
	return s, nil
}

// Replace string from command buffer
func expandHistReplace(s string, _ *History) (string, error) {
	sList := rxHistReplace.FindAllStringSubmatch(s, -1)
	var rxList []*regexp.Regexp
	var replaceList []string
	//debug.Json("^s/...", sList)
	for i := range sList {
		rx, err := regexp.Compile(sList[i][1])
		if err != nil || len(sList[i]) != 3 {
			//debug.Log("Regexp error.", err)
			//continue
			return "", fmt.Errorf("Regexp error in history variable `^s/%s/%s`: %s", sList[i][1], sList[i][2], err.Error())
		}
		rxList = append(rxList, rx)
		replaceList = append(replaceList, sList[i][2])
		s = strings.Replace(s, sList[i][0], "", -1)
	}
	for i := range rxList {
		s = rxList[i].ReplaceAllString(s, replaceList[i])
	}

	return s, nil
}

// Match history prefix
func expandHistPrefix(s string, hist *History) (string, error) {
	mhPrefix := rxHistPrefix.FindAllString(s, -1)
	for i := range mhPrefix {
		for h := len(hist.List) - 1; h > -1; h-- {
			if strings.HasPrefix(hist.List[h].Block, mhPrefix[i][2:]) {
				s = strings.Replace(s, mhPrefix[i], noColon(hist.List[h].Block), 1)
				//return s, nil //[]rune(s)
				goto next
			}
		}
		//return s, nil //[]rune(s)
		return "", fmt.Errorf("Cannot find a history item to match prefix: %s", mhPrefix[i])
	next:
	}

	return s, nil //[]rune(s)
}
