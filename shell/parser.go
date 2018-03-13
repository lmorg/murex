package shell

import (
	"github.com/lmorg/murex/utils/ansi"
	"regexp"
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

	rxAllowedVarChars *regexp.Regexp = regexp.MustCompile(`^[_a-zA-Z0-9]$`)
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

func syntaxHighlight(r []rune) string {
	_, highlighted := parse(r)
	return highlighted
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
