package shell

import (
	"regexp"
	"strings"
	"time"

	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/shell/history"
	"github.com/lmorg/murex/shell/variables"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/readline"
)

type murexCompleterIface struct{}

var (
	murexCompleter    *murexCompleterIface = new(murexCompleterIface)
	rxAllowedVarChars *regexp.Regexp       = regexp.MustCompile(`^[_a-zA-Z0-9]$`)
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
	Loc         int
	VarLoc      int
	Escaped     bool
	QuoteSingle bool
	QuoteDouble bool
	Bracket     int
	ExpectFunc  bool
	pop         *string
	FuncName    string
	Parameters  []string
	Variable    string
}

func parse(line []rune) (pt parseTokens, syntaxHighlighted string) {
	var readFunc bool
	reset := []string{ansi.Reset, hlFunction}
	syntaxHighlighted = hlFunction
	pt.Loc = -1
	pt.ExpectFunc = true
	pt.pop = &pt.FuncName

	ansiColour := func(colour string, r rune) {
		syntaxHighlighted += colour + string(r)
		reset = append(reset, colour)
	}

	ansiReset := func(r rune) {
		if len(reset) > 1 {
			reset = reset[:len(reset)-1]
		}
		syntaxHighlighted += string(r) + reset[len(reset)-1]
		if len(reset) == 1 && pt.Bracket > 0 {
			syntaxHighlighted += hlBlock
		}
	}

	ansiResetNoChar := func() {
		if len(reset) > 1 {
			reset = reset[:len(reset)-1]
		}
		syntaxHighlighted += reset[len(reset)-1]
		if len(reset) == 1 && pt.Bracket > 0 {
			syntaxHighlighted += hlBlock
		}
	}

	ansiChar := func(colour string, r rune) {
		syntaxHighlighted += colour + string(r) + reset[len(reset)-1]
		if len(reset) == 1 && pt.Bracket > 0 {
			syntaxHighlighted += hlBlock
		}
	}

	for i := range line {
		if pt.Variable != "" && !rxAllowedVarChars.MatchString(string(line[i])) {
			pt.Variable = ""
			ansiResetNoChar()
		}

		switch line[i] {
		case '#':
			pt.Loc = i
			switch {
			case pt.Escaped:
				pt.Escaped = false
				*pt.pop += `#`
				ansiReset(line[i])
			case pt.QuoteSingle, pt.QuoteDouble:
				*pt.pop += `#`
				syntaxHighlighted += string(line[i])
			default:
				syntaxHighlighted += hlComment + string(line[i:]) + ansi.Reset
				return
			}

		case '\\':
			switch {
			case pt.Escaped:
				pt.Escaped = false
				*pt.pop += `\`
				ansiReset(line[i])
			case pt.QuoteSingle, pt.QuoteDouble:
				*pt.pop += `\`
				syntaxHighlighted += string(line[i])
			default:
				pt.Escaped = true
				ansiColour(hlEscaped, line[i])
			}

		case '\'':
			pt.Loc = i
			switch {
			case pt.Escaped:
				pt.Escaped = false
				*pt.pop += `'`
				ansiReset(line[i])
			case pt.QuoteDouble:
				*pt.pop += `'`
				syntaxHighlighted += string(line[i])
			case pt.QuoteSingle:
				pt.QuoteSingle = false
				ansiReset(line[i])
			default:
				pt.QuoteSingle = true
				ansiColour(hlSingleQuote, line[i])
			}

		case '"':
			pt.Loc = i
			switch {
			case pt.Escaped:
				pt.Escaped = false
				*pt.pop += `"`
				ansiReset(line[i])
			case pt.QuoteSingle:
				*pt.pop += `"`
				syntaxHighlighted += string(line[i])
			case pt.QuoteDouble:
				pt.QuoteDouble = false
				ansiReset(line[i])
			default:
				pt.QuoteDouble = true
				ansiColour(hlDoubleQuote, line[i])
			}

		case ' ':
			switch {
			case pt.Escaped:
				pt.Escaped = false
				*pt.pop += ` `
				ansiReset(line[i])
			case pt.QuoteSingle, pt.QuoteDouble:
				*pt.pop += ` `
				syntaxHighlighted += string(line[i])
			case pt.ExpectFunc && readFunc:
				pt.Loc = i
				pt.ExpectFunc = false
				readFunc = false
				pt.Parameters = append(pt.Parameters, "")
				pt.pop = &pt.Parameters[0]
				ansiReset(line[i])
			default:
				pt.Loc = i
				pt.Parameters = append(pt.Parameters, "")
				pt.pop = &pt.Parameters[len(pt.Parameters)-1]
				syntaxHighlighted += string(line[i])
			}

		case '>':
			switch {
			case i > 0 && line[i-1] == '-':
				pt.Loc = i
				pt.ExpectFunc = true
				pt.pop = &pt.FuncName
				pt.Parameters = make([]string, 0)
				syntaxHighlighted = syntaxHighlighted[:len(syntaxHighlighted)-1]
				ansiColour(hlPipe, '-')
				ansiReset('>')
				syntaxHighlighted += hlFunction

			case pt.ExpectFunc, readFunc:
				readFunc = true
				*pt.pop += `>`
				fallthrough
			case pt.Escaped:
				pt.Escaped = false
				ansiReset(line[i])
			default:
				pt.Loc = i
				syntaxHighlighted += string(line[i])
			}

		case ';', '|':
			pt.Loc = i
			switch {
			case pt.Escaped:
				pt.Escaped = false
				*pt.pop += string(line[i])
				ansiReset(line[i])
			case pt.QuoteSingle, pt.QuoteDouble:
				*pt.pop += string(line[i])
				syntaxHighlighted += string(line[i])
			default:
				pt.ExpectFunc = true
				pt.pop = &pt.FuncName
				pt.Parameters = make([]string, 0)
				ansiChar(hlPipe, line[i])
				syntaxHighlighted += hlFunction
			}

		case '?':
			pt.Loc = i
			switch {
			case pt.Escaped:
				pt.Escaped = false
				*pt.pop += `?`
				ansiReset(line[i])
			case pt.QuoteSingle, pt.QuoteDouble:
				*pt.pop += `?`
				syntaxHighlighted += string(line[i])
			case i > 0 && line[i-1] == ' ':
				pt.ExpectFunc = true
				pt.pop = &pt.FuncName
				pt.Parameters = make([]string, 0)
				ansiChar(hlPipe, line[i])
				syntaxHighlighted += hlFunction
			default:
				*pt.pop += `?`
				syntaxHighlighted += string(line[i])
			}

		case '{':
			pt.Loc = i
			switch {
			case pt.Escaped:
				pt.Escaped = false
				*pt.pop += `{`
				ansiReset(line[i])
			case pt.QuoteSingle, pt.QuoteDouble:
				*pt.pop += `{`
				syntaxHighlighted += string(line[i])
			default:
				pt.Bracket++
				pt.ExpectFunc = true
				pt.pop = &pt.FuncName
				pt.Parameters = make([]string, 0)
				syntaxHighlighted += hlBlock + string(line[i])
			}

		case '}':
			switch {
			case pt.Escaped:
				pt.Escaped = false
				*pt.pop += `}`
				ansiReset(line[i])
			case pt.Escaped, pt.QuoteSingle, pt.QuoteDouble:
				*pt.pop += `}`
				syntaxHighlighted += string(line[i])
			default:
				pt.Bracket--
				syntaxHighlighted += string(line[i])
				if pt.Bracket == 0 {
					syntaxHighlighted += ansi.Reset + reset[len(reset)-1]
				}
			}

		case '$':
			pt.VarLoc = i
			switch {
			case pt.Escaped:
				pt.Escaped = false
				*pt.pop += string(line[i])
				ansiReset(line[i])
			case pt.QuoteSingle:
				*pt.pop += string(line[i])
				syntaxHighlighted += string(line[i])
			default:
				*pt.pop += string(line[i])
				pt.Variable = string(line[i])
				ansiColour(hlVariable, line[i])
			}

		case '@':
			pt.VarLoc = i
			switch {
			case pt.Escaped:
				pt.Escaped = false
				*pt.pop += string(line[i])
				ansiReset(line[i])
			case pt.QuoteSingle:
				*pt.pop += string(line[i])
				syntaxHighlighted += string(line[i])
			default:
				*pt.pop += string(line[i])

				if i > 0 && (line[i-1] == ' ' || line[i-1] == '\t') {
					pt.Variable = string(line[i])
					ansiColour(hlVariable, line[i])
				} else {
					syntaxHighlighted += string(line[i])
				}
			}

		case ':':
			switch {
			case pt.Escaped:
				pt.Escaped = false
				*pt.pop += `:`
				ansiReset(line[i])
			case pt.QuoteSingle, pt.QuoteDouble:
				*pt.pop += `:`
				syntaxHighlighted += string(line[i])
			case !pt.ExpectFunc:
				*pt.pop += `:`
				syntaxHighlighted += string(line[i])
			default:
				syntaxHighlighted += string(line[i])
			}

		default:
			switch {
			case pt.Escaped:
				pt.Escaped = false
				ansiReset(line[i])
			case readFunc:
				*pt.pop += string(line[i])
				syntaxHighlighted += string(line[i])
			case pt.ExpectFunc:
				*pt.pop = string(line[i])
				readFunc = true
				syntaxHighlighted += string(line[i])
			default:
				*pt.pop += string(line[i])
				syntaxHighlighted += string(line[i])
			}
		}
	}
	pt.Loc++
	pt.VarLoc++
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
	case pt.Variable != "":
		var s string
		if pt.VarLoc < len(line) {
			s = strings.TrimSpace(string(line[pt.VarLoc:]))
		}
		s = pt.Variable + s
		retPos = len(s)
		items = autocomplete.MatchVars(s)

	case pt.ExpectFunc:
		var s string
		if pt.Loc < len(line) {
			s = strings.TrimSpace(string(line[pt.Loc:]))
		}
		retPos = len(s)
		items = autocomplete.MatchFunction(s)

	default:
		var s string
		if len(pt.Parameters) > 0 {
			s = pt.Parameters[len(pt.Parameters)-1]
		}
		retPos = len(s)

		autocomplete.InitExeFlags(pt.FuncName)

		pIndex := 0
		items = autocomplete.MatchFlags(autocomplete.ExesFlags[pt.FuncName], s, pt.FuncName, pt.Parameters, &pIndex)
	}

	v, err := proc.ShellProcess.Config.Get("shell", "max-suggestions", types.Integer)
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
		if len(items[i]) == 0 {
			continue
		}

		if !pt.QuoteSingle && !pt.QuoteDouble && len(items[i]) > 1 && strings.Contains(items[i][:len(items[i])-1], " ") {
			items[i] = strings.Replace(items[i], " ", `\ `, -1)
		}

		if items[i][len(items[i])-1] == '/' || items[i][len(items[i])-1] == '=' {
			suggest[i] = []rune(items[i])
		} else {
			suggest[i] = []rune(items[i] + " ")
		}
	}

	return
}

func listener(line []rune, pos int, key rune) (newLine []rune, newPos int, ok bool) {
	typed := time.Now().After(keyPressTimer)
	keyPressTimer = time.Now().Add(20 * time.Millisecond)

	switch {
	case key == readline.CharEnter:
		return nil, 0, ok

	case forward == 2 && pos == len(line):
		//newLine = expandVariables(line)
		//newLine = expandHistory(newLine)
		//newPos = len(newLine)
		expanded, err := history.ExpandVariables(line, History)
		if err != nil {
			ansi.Stderrln(ansi.FgRed, utils.NewLineString+err.Error())
		} else {
			ansi.Stderrln(ansi.FgBlue, utils.NewLineString+string(variables.Expand(expanded)))
		}
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
		if !pt.Escaped && !pt.QuoteSingle && !pt.QuoteDouble {
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
		if !pt.Escaped && !pt.QuoteSingle && !pt.QuoteDouble {
			newLine = smooshLines(line, pos, ']')
			newPos = pos
		} else {
			newPos = pos
			newLine = line
		}

	case key == '\'' && typed:
		pt, _ := parse(line)
		forward = 0
		if !pt.Escaped && pt.QuoteSingle && !pt.QuoteDouble {
			newLine = smooshLines(line, pos, '\'')
			newPos = pos
		} else {
			newPos = pos
			newLine = line
		}

	case key == '"' && typed:
		pt, _ := parse(line)
		forward = 0
		if !pt.Escaped && !pt.QuoteSingle && pt.QuoteDouble {
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
		case pt.Bracket < 0:
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
		case pt.QuoteSingle:
			newLine, newPos = unsmooshLines(line, pos, '\'')
		case pt.QuoteDouble:
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
	// up stuff with slices sometimes due to them literally just being pointers.
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
