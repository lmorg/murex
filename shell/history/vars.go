package history

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/utils/readline"
)

var (
	rxHistIndex  = regexp.MustCompile(`(\^[0-9]+)`)
	rxHistRegex  = regexp.MustCompile(`\^m/(.*?[^\\])/`) // Scratchpad: https://play.golang.org/p/Iya2Hx1uxb
	rxHistPrefix = regexp.MustCompile(`(\^\^[a-zA-Z]+)`)
	rxHistTag    = regexp.MustCompile(`(\^#[-_a-zA-Z0-9]+)`)
	//rxHistAllPs    = regexp.MustCompile(`\^\[([-]?[0-9]+)]\[([-]?[0-9]+)]`)
	rxHistParam   = regexp.MustCompile(`\^\[([-]?[0-9]+)]`)
	rxHistReplace = regexp.MustCompile(`\^s/(.*?[^\\])/(.*?[^\\])/`)
	//rxHistRepParam = regexp.MustCompile(`\^s([-]?[0-9]+)/(.*?[^\\])/(.*?[^\\])/`)
)

const (
	errCannotParsePrevCmd = "cannot parse previous command line to extract parameters for history variable"
)

func getLine(i int, rl *readline.Instance) (s string) {
	s, _ = rl.History.GetLine(i)
	return
}

// ExpandVariables finds all history variables and replaces them with the value
// of the variable
func ExpandVariables(line []rune, rl *readline.Instance) ([]rune, error) {
	return expandVariables(line, rl, false)
}

// ExpandVariablesInLine finds history variables in a line and replaces it with
// the value of the variable. It does not replace the line formatting variables.
func ExpandVariablesInLine(line []rune, rl *readline.Instance) ([]rune, error) {
	return expandVariables(line, rl, true)
}

func expandVariables(line []rune, rl *readline.Instance, skipFormatting bool) ([]rune, error) {
	s := string(line)

	if !skipFormatting {
		s = strings.Replace(s, `^\n`, "\n", -1) // Match new line
		s = strings.Replace(s, `^\t`, "\t", -1) // Match tab
	}

	funcs := []func(string, *readline.Instance) (string, error){
		expandHistBangBang,
		expandHistPrefix,
		expandHistIndex,
		expandHistRegex,
		expandHistHashtag,
		//expandHistAllPs,
		expandHistParam,
		expandHistReplace,
		//expandHistRepParam,
	}

	for f := range funcs {
		var err error
		s, err = funcs[f](s, rl)
		if err != nil {
			return nil, err
		}
	}

	return []rune(s), nil
}

// Match last command
func expandHistBangBang(s string, rl *readline.Instance) (string, error) {
	last := getLine(rl.History.Len()-1, rl)
	return strings.Replace(s, "^!!", noColon(last), -1), nil
}

// Match history index
func expandHistIndex(s string, rl *readline.Instance) (string, error) {
	mhIndex := rxHistIndex.FindAllString(s, -1)
	for i := range mhIndex {
		val, _ := strconv.Atoi(mhIndex[i][1:])
		if val > rl.History.Len() {
			return "", fmt.Errorf("(%s) Value greater than history length in ^%d", mhIndex[i], val)
		}
		s = rxHistIndex.ReplaceAllString(s, noColon(getLine(val, rl)))
		//return s, nil
	}
	return s, nil
}

// Match history by regexp
func expandHistRegex(s string, rl *readline.Instance) (string, error) {
	mhRegexp := rxHistRegex.FindAllStringSubmatch(s, -1)
	for i := range mhRegexp {
		rx, err := regexp.Compile(mhRegexp[i][1])
		if err != nil {
			return "", fmt.Errorf("(%s) Regexp error in history variable `^m/%s/`: %s", mhRegexp[i][0], mhRegexp[i][1], err.Error())
		}

		for h := rl.History.Len() - 1; h > -1; h-- {
			if rx.MatchString(getLine(h, rl)) {
				s = rxHistRegex.ReplaceAllString(s, noColon(getLine(h, rl)))
				goto next
			}
		}
		return "", fmt.Errorf("(%s) Cannot find a history item to match regexp: %s", mhRegexp[i][0], mhRegexp[i][1])
	next:
	}
	return s, nil
}

// Match history hashtag
func expandHistHashtag(s string, rl *readline.Instance) (string, error) {
	mhTag := rxHistTag.FindAllString(s, -1)

	for i := range mhTag {
		for h := rl.History.Len() - 1; h > -1; h-- {
			line := getLine(h, rl)
			if strings.HasSuffix(line, mhTag[i][1:]) {
				block := line[:len(line)-len(mhTag[i][1:])]
				s = strings.Replace(s, mhTag[i], noColon(block), 1)

				goto next
			}
		}

		return "", fmt.Errorf("(%s) Hashtag not found", mhTag[i])
	next:
	}

	return s, nil
}

/*// Match last params (all of block)
func expandHistAllPs(s string, rl *readline.Instance) (string, error) {
	mhParam := rxHistAllPs.FindAllStringSubmatch(s, -1)
	if len(mhParam) > 0 {
		last := getLine(rl.History.Len(), rl)
		nodes, pErr := lang.ParseBlock([]rune(last))
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
				return "", fmt.Errorf("(%s) Cannot extract parameter", mhParam[i][0])
			}

			p := parameters.Parameters{Tokens: nodes[cmd].ParamTokens}
			lang.ParseParameters(lang.ShellProcess, &p)
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
}*/

// Match last params (first command in block)
func expandHistParam(s string, rl *readline.Instance) (string, error) {
	mhParam := rxHistParam.FindAllStringSubmatch(s, -1)
	if len(mhParam) > 0 {
		last := getLine(rl.History.Len()-1, rl)
		nodes, pErr := lang.ParseBlock([]rune(last))
		if pErr.Code != lang.NoParsingErrors {
			return "", fmt.Errorf(errCannotParsePrevCmd)
		}
		p := parameters.Parameters{Tokens: nodes.Last().ParamTokens}
		err := lang.ParseParameters(lang.ShellProcess, &p)
		if err != nil {
			return s, err
		}

		for i := range mhParam {
			val, _ := strconv.Atoi(mhParam[i][1])
			if val < 0 {
				val += p.Len() + 1
			}

			switch {
			case val == 0:
				s = strings.Replace(s, mhParam[i][0], nodes.Last().Name, -1)
			case val > 0 && val-1 < p.Len():
				s = strings.Replace(s, mhParam[i][0], p.StringArray()[val-1], -1)
			default:
				s = strings.Replace(s, mhParam[i][0], "", -1)
				return s, fmt.Errorf("(%s) No parameter with index %s", mhParam[i][0], mhParam[i][1])
			}

		}

		//return s, nil //err
	}
	return s, nil
}

// Replace string from command buffer
func expandHistReplace(s string, rl *readline.Instance) (string, error) {
	sList := rxHistReplace.FindAllStringSubmatch(s, -1)
	var rxList []*regexp.Regexp
	var replaceList []string

	for i := range sList {
		rx, err := regexp.Compile(sList[i][1])
		if err != nil || len(sList[i]) != 3 {
			return "", fmt.Errorf("(%s) Regexp error in history variable `^s/%s/%s`: %s", sList[i][0], sList[i][1], sList[i][2], err.Error())
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

/*// Replace string from a parameter in the last command
func expandHistRepParam(s string, rl *readline.Instance) (string, error) {
	mhRepParam := rxHistRepParam.FindAllStringSubmatch(s, -1)
	if len(mhRepParam) > 0 {
		last := getLine(rl.History.Len()-1, rl)
		nodes, pErr := lang.ParseBlock([]rune(last))
		if pErr.Code != lang.NoParsingErrors {
			return "", errors.New(errCannotParsePrevCmd)
		}
		p := parameters.Parameters{Tokens: nodes.Last().ParamTokens}
		lang.ParseParameters(lang.ShellProcess, &p)

		for i := range mhRepParam {
			param, err := strconv.Atoi(mhRepParam[i][1])
			if err != nil {
				return "", fmt.Errorf("(%s) Unable to convert '%s' to int", mhRepParam[i][0], mhRepParam[i][1])
			}
			if param < 0 {
				param += p.Len() + 1
			}

			rx, err := regexp.Compile(mhRepParam[i][2])
			if err != nil {
				return "", fmt.Errorf("(%s) Error compiling regexp '%s': %s", mhRepParam[i][0], mhRepParam[i][2], err.Error())
			}
			var old string
			switch {
			case param == 0:
				old = nodes.Last().Name
			case param > 0 && param-1 < p.Len():
				old, err = p.String(param - 1)
				if err != nil {
					return "", fmt.Errorf("(%s) Parameter error for %d (derived from '%s'): %s", mhRepParam[i][0], param, mhRepParam[i][1], err.Error())
				}
			default:
				return "", fmt.Errorf("(%s) Parameter index out of bounds: %d (derived from '%s')", mhRepParam[i][0], param, mhRepParam[i][1])
			}
			new := rx.ReplaceAllString(old, mhRepParam[i][3])
			s = strings.Replace(s, mhRepParam[i][0], new, -1)
		}

	}
	return s, nil
}*/

// Match history prefix
func expandHistPrefix(s string, rl *readline.Instance) (string, error) {
	mhPrefix := rxHistPrefix.FindAllString(s, -1)
	for i := range mhPrefix {
		for h := rl.History.Len() - 1; h > -1; h-- {
			if strings.HasPrefix(getLine(h, rl), mhPrefix[i][2:]) {
				s = strings.Replace(s, mhPrefix[i], noColon(getLine(h, rl)), 1)

				goto next
			}
		}

		return "", fmt.Errorf("(%s) Cannot find a history item to match prefix", mhPrefix[i])
	next:
	}

	return s, nil
}
