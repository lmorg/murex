package shell

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/home"
	"github.com/lmorg/murex/utils/man"
	"github.com/lmorg/readline"
	"regexp"
	"strings"
	"time"
)

type murexCompleterIface struct{}

var (
	murexCompleter    *murexCompleterIface = new(murexCompleterIface)
	rxAllowedVarChars *regexp.Regexp       = regexp.MustCompile(`^[_a-zA-Z0-9]$`)
	rxVars            *regexp.Regexp       = regexp.MustCompile(`(\$[_a-zA-Z0-9]+)`)
	rxHome            *regexp.Regexp       = regexp.MustCompile(`(~[_\-.a-zA-Z0-9]+)`)
	keyPressTimer     time.Time
)

// syntax highlighting
var (
	hlFunction    string = ansi.Bold
	hlVariable    string = ansi.FgGreen
	hlEscaped     string = ansi.FgYellow
	hlSingleQuote string = ansi.FgBlue
	hlDoubleQuote string = ansi.FgBlue
	hlBlock       string = ansi.BgBlackBright
	hlPipe        string = ansi.FgMagenta
	hlComment     string = ansi.BgGreenBright
)

type parseTokens struct {
	loc        int
	vloc       int
	escaped    bool
	qSingle    bool
	qDouble    bool
	bracket    int
	expectFunc bool
	__pop      *string
	funcName   string
	parameters []string
	variable   string
}

func parse(line []rune) (pt parseTokens, syntaxHighlighted string) {
	var readFunc bool
	reset := []string{ansi.Reset, hlFunction}
	syntaxHighlighted = hlFunction
	pt.loc = -1
	pt.expectFunc = true
	pt.__pop = &pt.funcName

	ansiColour := func(colour string, r rune) {
		syntaxHighlighted += colour + string(r)
		reset = append(reset, colour)
	}

	ansiReset := func(r rune) {
		if len(reset) > 1 {
			reset = reset[:len(reset)-1]
		}
		syntaxHighlighted += string(r) + reset[len(reset)-1]
		if len(reset) == 1 && pt.bracket > 0 {
			syntaxHighlighted += hlBlock
		}
	}

	ansiResetNoChar := func() {
		if len(reset) > 1 {
			reset = reset[:len(reset)-1]
		}
		syntaxHighlighted += reset[len(reset)-1]
		if len(reset) == 1 && pt.bracket > 0 {
			syntaxHighlighted += hlBlock
		}
	}

	ansiChar := func(colour string, r rune) {
		syntaxHighlighted += colour + string(r) + reset[len(reset)-1]
		if len(reset) == 1 && pt.bracket > 0 {
			syntaxHighlighted += hlBlock
		}
	}

	for i := range line {
		if pt.variable != "" && !rxAllowedVarChars.MatchString(string(line[i])) {
			pt.variable = ""
			ansiResetNoChar()
		}

		switch line[i] {
		case '#':
			pt.loc = i
			switch {
			case pt.escaped:
				pt.escaped = false
				//if readFunc {
				*pt.__pop += `#`
				//}
				ansiReset(line[i])
			case pt.qSingle, pt.qDouble:
				//if readFunc {
				*pt.__pop += `#`
				//}
				//pt.escaped = false
				syntaxHighlighted += string(line[i])
			default:
				syntaxHighlighted += hlComment + string(line[i:]) + ansi.Reset
				return
			}

		case '\\':
			switch {
			case pt.escaped:
				pt.escaped = false
				//if readFunc {
				*pt.__pop += `\`
				//}
				ansiReset(line[i])
			case pt.qSingle, pt.qDouble:
				//if readFunc {
				*pt.__pop += `\`
				//}
				syntaxHighlighted += string(line[i])
			default:
				pt.escaped = true
				ansiColour(hlEscaped, line[i])
			}

		case '\'':
			pt.loc = i
			switch {
			case pt.escaped:
				pt.escaped = false
				//if readFunc {
				*pt.__pop += `'`
				//}
				ansiReset(line[i])
			case pt.qDouble:
				//if readFunc {
				*pt.__pop += `'`
				//}
				syntaxHighlighted += string(line[i])
			case pt.qSingle:
				pt.qSingle = false
				ansiReset(line[i])
			default:
				pt.qSingle = true
				ansiColour(hlSingleQuote, line[i])
			}

		case '"':
			pt.loc = i
			switch {
			case pt.escaped:
				pt.escaped = false
				//if readFunc {
				*pt.__pop += `"`
				//}
				ansiReset(line[i])
			case pt.qSingle:
				//if readFunc {
				*pt.__pop += `"`
				//}
				syntaxHighlighted += string(line[i])
			case pt.qDouble:
				pt.qDouble = false
				ansiReset(line[i])
			default:
				pt.qDouble = true
				ansiColour(hlDoubleQuote, line[i])
			}

		case ' ':
			//pt.loc = i
			switch {
			case pt.escaped:
				pt.escaped = false
				//if readFunc {
				*pt.__pop += ` `
				//}
				ansiReset(line[i])
			case pt.qSingle, pt.qDouble:
				//if readFunc {
				*pt.__pop += ` `
				//}
				syntaxHighlighted += string(line[i])
			case pt.expectFunc && readFunc:
				pt.loc = i
				pt.expectFunc = false
				readFunc = false
				pt.parameters = append(pt.parameters, "")
				pt.__pop = &pt.parameters[0]
				//syntaxHighlighted += string(line[i])
				ansiReset(line[i])
			default:
				pt.loc = i
				pt.parameters = append(pt.parameters, "")
				pt.__pop = &pt.parameters[len(pt.parameters)-1]
				syntaxHighlighted += string(line[i])
			}

		case '>':
			switch {
			case i > 0 && line[i-1] == '-':
				pt.loc = i
				pt.expectFunc = true
				pt.__pop = &pt.funcName
				pt.parameters = make([]string, 0)
				//syntaxHighlighted += string(line[i])
				syntaxHighlighted = syntaxHighlighted[:len(syntaxHighlighted)-1]
				ansiColour(hlPipe, '-')
				ansiReset('>')
				syntaxHighlighted += hlFunction

			case pt.expectFunc, readFunc:
				readFunc = true
				*pt.__pop += `>`
				fallthrough
			case pt.escaped:
				pt.escaped = false
				ansiReset(line[i])
			default:
				pt.loc = i
				syntaxHighlighted += string(line[i])
			}

		case ';', '|':
			pt.loc = i
			switch {
			case pt.escaped:
				pt.escaped = false
				//if readFunc {
				*pt.__pop += string(line[i])
				//}
				ansiReset(line[i])
			case pt.qSingle, pt.qDouble:
				//if readFunc {
				*pt.__pop += string(line[i])
				//}
				syntaxHighlighted += string(line[i])
			default:
				pt.expectFunc = true
				pt.__pop = &pt.funcName
				pt.parameters = make([]string, 0)
				//syntaxHighlighted += string(line[i])
				ansiChar(hlPipe, line[i])
				syntaxHighlighted += hlFunction
			}

		case '?':
			pt.loc = i
			switch {
			case pt.escaped:
				pt.escaped = false
				//if readFunc {
				*pt.__pop += `?`
				//}
				ansiReset(line[i])
			case pt.qSingle, pt.qDouble:
				//if readFunc {
				*pt.__pop += `?`
				//}
				syntaxHighlighted += string(line[i])
			case i > 0 && line[i-1] == ' ':
				pt.expectFunc = true
				pt.__pop = &pt.funcName
				pt.parameters = make([]string, 0)
				//syntaxHighlighted += string(line[i])
				ansiChar(hlPipe, line[i])
				syntaxHighlighted += hlFunction
			default:
				//if readFunc {
				*pt.__pop += `?`
				//}
				syntaxHighlighted += string(line[i])
			}

		case '{':
			pt.loc = i
			switch {
			case pt.escaped:
				pt.escaped = false
				//if readFunc {
				*pt.__pop += `{`
				//}
				ansiReset(line[i])
			case pt.qSingle, pt.qDouble:
				//if readFunc {
				*pt.__pop += `{`
				//}
				syntaxHighlighted += string(line[i])
			default:
				pt.bracket++
				pt.expectFunc = true
				pt.__pop = &pt.funcName
				pt.parameters = make([]string, 0)
				syntaxHighlighted += hlBlock + string(line[i])
				//ansiColour(ansi.BgBlackBright, line[i])
			}

		case '}':
			//loc = i
			switch {
			case pt.escaped:
				pt.escaped = false
				//if readFunc {
				*pt.__pop += `}`
				//}
				ansiReset(line[i])
			case pt.escaped, pt.qSingle, pt.qDouble:
				//if readFunc {
				*pt.__pop += `}`
				//}
				syntaxHighlighted += string(line[i])
			default:
				pt.bracket--
				syntaxHighlighted += string(line[i])
				if pt.bracket == 0 {
					syntaxHighlighted += ansi.Reset + reset[len(reset)-1]
				}
				//ansiReset(line[i])
			}

		case '$':
			//pt.loc = i
			pt.vloc = i
			switch {
			case pt.escaped:
				pt.escaped = false
				//if readFunc {
				*pt.__pop += string(line[i])
				//}
				ansiReset(line[i])
			case pt.qSingle:
				//if readFunc {
				*pt.__pop += string(line[i])
				//}
				syntaxHighlighted += string(line[i])
			default:
				*pt.__pop += string(line[i])
				pt.variable = string(line[i])
				ansiColour(hlVariable, line[i])
			}

		case '@':
			//pt.loc = i
			pt.vloc = i
			switch {
			case pt.escaped:
				pt.escaped = false
				//if readFunc {
				*pt.__pop += string(line[i])
				//}
				ansiReset(line[i])
			case pt.qSingle:
				//if readFunc {
				*pt.__pop += string(line[i])
				//}
				syntaxHighlighted += string(line[i])
			default:
				*pt.__pop += string(line[i])

				if i > 0 && (line[i-1] == ' ' || line[i-1] == '\t') {
					pt.variable = string(line[i])
					ansiColour(hlVariable, line[i])
				} else {
					syntaxHighlighted += string(line[i])
				}
			}

		case ':':
			switch {
			case pt.escaped:
				pt.escaped = false
				//if readFunc {
				*pt.__pop += `:`
				//}
				ansiReset(line[i])
			case pt.qSingle, pt.qDouble:
				//if readFunc {
				*pt.__pop += `:`
				//}
				syntaxHighlighted += string(line[i])
			case !pt.expectFunc:
				*pt.__pop += `:`
				syntaxHighlighted += string(line[i])
			default:
				syntaxHighlighted += string(line[i])
			}

		default:
			switch {
			case pt.escaped:
				pt.escaped = false
				ansiReset(line[i])
			case readFunc:
				*pt.__pop += string(line[i])
				syntaxHighlighted += string(line[i])
			case pt.expectFunc:
				*pt.__pop = string(line[i])
				readFunc = true
				syntaxHighlighted += string(line[i])
			default:
				*pt.__pop += string(line[i])
				syntaxHighlighted += string(line[i])
			}
		}
	}
	pt.loc++
	pt.vloc++
	syntaxHighlighted += ansi.Reset
	return
}

func (mc murexCompleterIface) Do(line []rune, pos int) (suggest [][]rune, retPos int) {
	var items []string
	if len(line) > pos-1 {
		line = line[:pos]
	}

	pt, _ := parse(line)

	switch {
	case pt.variable != "":
		var s string
		if pt.vloc < len(line) {
			s = strings.TrimSpace(string(line[pt.vloc:]))
		}
		s = pt.variable + s
		retPos = len(s)
		items = matchVars(s)

	case pt.expectFunc:
		var s string
		if pt.loc < len(line) {
			s = strings.TrimSpace(string(line[pt.loc:]))
		}
		retPos = len(s)
		switch {
		case isLocal(s):
			items = matchLocal(s, true)
			items = append(items, matchDirs(s)...)
		default:
			exes := allExecutables(true)
			items = matchExes(s, exes, true)
		}

	default:
		//items = []string{"{ ", "-> ", "| ", " ? ", "; "}
		//var s string
		//if pt.loc < len(line) {
		//	s = strings.TrimSpace(string(line[pt.loc:]))
		//}
		var s string
		if len(pt.parameters) > 0 {
			s = pt.parameters[len(pt.parameters)-1]
		}
		retPos = len(s)

		/*items = matchFlags(s, pt.funcName)
		items = append(items, matchDynamic(s, pt.funcName, pt.parameters)...)

		if ExesFlags[pt.funcName].IncExePath {
			pathexes := allExecutables(false)
			items = append(items, matchExes(s, pathexes, false)...)
		}

		switch {
		case !ExesFlags[pt.funcName].NoFiles:
			items = append(items, matchFilesAndDirs(s)...)
		case !ExesFlags[pt.funcName].NoDirs:
			items = append(items, matchDirs(s)...)
		}*/
		//items = matchFlags(0, s, pt.funcName, pt.parameters)

		if len(ExesFlags[pt.funcName]) == 0 {
			ExesFlags[pt.funcName] = []Flags{{
				Flags:         man.ScanManPages(pt.funcName),
				IncFiles:      true,
				AllowMultiple: true,
			}}
		}

		pIndex := 0
		items = matchFlags(ExesFlags[pt.funcName], s, pt.funcName, pt.parameters, &pIndex)
	}

	v, err := proc.GlobalConf.Get("shell", "max-suggestions", types.Integer)
	if err != nil {
		v = -1
	}

	limitSuggestions := v.(int)
	if len(items) < limitSuggestions || limitSuggestions < 0 {
		limitSuggestions = len(items)
	}

	Instance.Config.MaxCompleteLines = limitSuggestions

	suggest = make([][]rune, len(items))
	for i := range items {
		if !pt.qSingle && !pt.qDouble && len(items[i]) > 1 && strings.Contains(items[i][:len(items[i])-1], " ") {
			items[i] = strings.Replace(items[i][:len(items[i])-1], " ", `\ `, -1) + items[i][len(items[i])-1:]
		}
		suggest[i] = []rune(items[i])
	}

	return
}

func listener(line []rune, pos int, key rune) (newLine []rune, newPos int, ok bool) {
	typed := time.Now().After(keyPressTimer)
	keyPressTimer = time.Now().Add(20 * time.Millisecond)

	switch {
	/*case key == 77:
	line = expandVariables(line)
	line = expandHistory(line)
	os.Stderr.WriteString(string(line) + utils.NewLineString)
	return line, pos, ok*/

	case key == readline.CharEnter:
		return nil, 0, ok

	case forward == 2 && pos == len(line):
		//newLine = expandVariables(line)
		//newLine = expandHistory(newLine)
		//newPos = len(newLine)
		ansi.Stderrln(ansi.FgBlue, utils.NewLineString+string(expandVariables(expandHistory(line))))
		newLine = line
		newPos = pos
		forward = 0

	case forward == 1 && pos == len(line):
		/*s := string(line)
		if len(rxVars.FindAllString(s, -1)) > 0 || strings.Contains(s, "~") ||
			len(rxHistIndex.FindAllString(s, -1)) > 0 ||
			len(rxHistRegex.FindAllString(s, -1)) > 0 ||
			len(rxHistPrefix.FindAllString(s, -1)) > 0 ||
			len(rxHistTag.FindAllString(s, -1)) > 0 ||
			len(rxHistAllPs.FindAllString(s, -1)) > 0 ||
			len(rxHistParam.FindAllString(s, -1)) > 0 ||
			strings.Contains(s, "^!!") {
			//os.Stderr.WriteString(utils.NewLineString + "Tap forward again to expand $VARS, ~HOME and ^HISTORY." + utils.NewLineString)
		} else {
			forward = 0
		}*/
		if len(line) == 0 {
			forward = 0
		}
		newPos = pos
		newLine = line

	case key == '{' && typed:
		pt, _ := parse(line)
		forward = 0
		if !pt.escaped && !pt.qSingle && !pt.qDouble {
			//newLine = append(line, '}')
			//newPos = len(newLine) - 1
			newLine = smooshLines(line, pos, '}')
			newPos = pos
		} else {
			newPos = pos
			newLine = line
		}

	case key == '[' && typed:
		pt, _ := parse(line)
		forward = 0
		if !pt.escaped && !pt.qSingle && !pt.qDouble {
			newLine = smooshLines(line, pos, ']')
			newPos = pos
		} else {
			newPos = pos
			newLine = line
		}

	case key == '\'' && typed:
		pt, _ := parse(line)
		forward = 0
		if !pt.escaped && pt.qSingle && !pt.qDouble {
			newLine = smooshLines(line, pos, '\'')
			newPos = pos
		} else {
			newPos = pos
			newLine = line
		}

	case key == '"' && typed:
		pt, _ := parse(line)
		forward = 0
		if !pt.escaped && !pt.qSingle && pt.qDouble {
			newLine = smooshLines(line, pos, '"')
			newPos = pos
		} else {
			newPos = pos
			newLine = line
		}

	case key == readline.CharBackspace, key == readline.CharDelete:
		newLine = line
		newPos = pos
		forward = 0

		pt, _ := parse(line)
		switch {
		case pt.bracket < 0:
			for i := pos; i < len(line); i++ {
				if line[i] == '}' {
					newLine = line[:i]
					if i < len(line) {
						newLine = append(newLine, line[i+1:]...)
					} else {
						newPos = newPos - 1
					}
				}
			}
		case pt.qSingle:
			newLine, newPos = unsmooshLines(line, pos, '\'')
		case pt.qDouble:
			newLine, newPos = unsmooshLines(line, pos, '"')
		}

	default:
		forward = 0
		newPos = pos
		newLine = line
	}

	if newPos > len(newLine) {
		newPos = len(line) - 1
	} else if newPos < 0 {
		newPos = 0
	}

	return newLine, newPos, true
}

func smooshLines(line []rune, pos int, injectedChar rune) []rune {
	if pos == len(line) {
		return append(line, injectedChar)
	}

	// How the did this happen?
	if pos > len(line) {
		return line
	}

	var i int
	for i = pos; i < len(line); i++ {
		if line[i] == ' ' || line[i] == '\t' {
			continue
		} //else {
		break
		//}
	}

	// It might seem odd converting this into a string only to convert back to []rune but Go does some pretty fucked
	// up stuff with slices sometimes due to them literally just being pointers. I found this caused all kinds of
	// annoying little glitches in this routine, as simple as it seems.
	if i >= len(line) || line[i] == '}' || line[i] == '|' || line[i] == ';' {
		before := string(line[:pos])
		after := string(line[pos:])
		newLine := before + string(injectedChar) + after
		return []rune(newLine)
	}

	return line
}

func unsmooshLines(line []rune, pos int, injectedChar rune) ([]rune, int) {
	if pos > len(line) {
		return line, pos
	}

	if pos < 2 {
		if line[0] == injectedChar {
			if len(line) > 1 {
				return line[1:], pos
			} // else {
			return []rune{}, 0
			//}
		} //else {
		return line, pos
		//}
	}

	if pos == len(line) {
		if line[pos-1] == injectedChar {
			if len(line) > 1 {
				return line[:pos], pos
			} //else {
			return []rune{}, 0
			//}
		} //else {
		return line, pos
		//}
	}

	if line[pos] == injectedChar {
		before := string(line[:pos])
		after := string(line[pos+1:])
		return []rune(before + after), pos
	}

	return line, pos
}

func expandVariablesString(line string) string {
	match := rxVars.FindAllString(line, -1)
	for i := range match {
		line = strings.Replace(line, match[i], proc.GlobalVars.GetString(match[i][1:]), -1)
	}

	match = rxHome.FindAllString(line, -1)
	for i := range match {
		line = rxHome.ReplaceAllString(line, home.UserDir(match[i][1:]))
	}

	line = strings.Replace(line, "~", home.MyDir, -1)
	return line
}

func expandVariables(line []rune) []rune {
	return []rune(expandVariablesString(string(line)))
}
