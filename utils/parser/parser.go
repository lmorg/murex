package parser

import (
	"regexp"

	"github.com/lmorg/murex/utils/ansi"
)

// syntax highlighting
var (
	hlFunction    = ansi.Bold
	hlVariable    = ansi.FgGreen
	hlEscaped     = ansi.FgYellow
	hlSingleQuote = ansi.FgBlue
	hlDoubleQuote = ansi.FgBlue
	hlBraceQuote  = ansi.FgBlue
	hlBlock       = ansi.BgBlackBright
	hlPipe        = ansi.FgMagenta
	hlComment     = ansi.BgGreenBright
	hlError       = ansi.BgRed

	rxAllowedVarChars = regexp.MustCompile(`^[_a-zA-Z0-9]$`)
)

// ParsedTokens is a struct that returns a tokenized version of the selected command
type ParsedTokens struct {
	Source        []rune
	Loc           int
	VarLoc        int
	Escaped       bool
	QuoteSingle   bool
	QuoteDouble   bool
	QuoteBrace    int
	NestedBlock   int
	ExpectFunc    bool
	pop           *string
	FuncName      string
	Parameters    []string
	Variable      string
	Unsafe        bool // if the pipeline is estimated to be safe enough to dynamically preview
	LastFlowToken int
}

// SafeFunctions is a list of all the functions considered safe when using tab
// autocomplete with ExecCmdline
var SafeFunctions = []string{
	"open", "regexp", "match", "cast", "format", "[", "[[", "runtime",
	"cat", "ps", "grep", "ls",
}

// Parse a single line of code and return the tokens for a selected command
func Parse(block []rune, pos int) (pt ParsedTokens, syntaxHighlighted string) {
	var readFunc bool
	reset := []string{ansi.Reset, hlFunction}
	syntaxHighlighted = hlFunction
	pt.Loc = -1
	pt.ExpectFunc = true
	pt.pop = &pt.FuncName
	pt.Source = block

	ansiColour := func(colour string, r rune) {
		syntaxHighlighted += colour + string(r)
		reset = append(reset, colour)
	}

	ansiReset := func(r rune) {
		if len(reset) > 1 {
			reset = reset[:len(reset)-1]
		}
		syntaxHighlighted += string(r) + reset[len(reset)-1]
		if len(reset) == 1 && pt.NestedBlock > 0 {
			syntaxHighlighted += hlBlock
		}
	}

	ansiResetNoChar := func() {
		if len(reset) > 1 {
			reset = reset[:len(reset)-1]
		}
		syntaxHighlighted += reset[len(reset)-1]
		if len(reset) == 1 && pt.NestedBlock > 0 {
			syntaxHighlighted += hlBlock
		}
	}

	ansiChar := func(colour string, r rune) {
		syntaxHighlighted += colour + string(r) + reset[len(reset)-1]
		if len(reset) == 1 && pt.NestedBlock > 0 {
			syntaxHighlighted += hlBlock
		}
	}

	for i := range block {
		if pt.Variable != "" && !rxAllowedVarChars.MatchString(string(block[i])) {
			pt.Variable = ""
			ansiResetNoChar()
		}

		switch block[i] {
		case '#':
			pt.Loc = i
			switch {
			case pt.Escaped:
				pt.Escaped = false
				*pt.pop += `#`
				ansiReset(block[i])
			case pt.QuoteSingle, pt.QuoteDouble, pt.QuoteBrace > 0, pt.NestedBlock > 0:
				*pt.pop += `#`
				syntaxHighlighted += string(block[i])
			default:
				syntaxHighlighted += hlComment + string(block[i:]) + ansi.Reset
				return
			}

		case '\\':
			switch {
			case pt.Escaped:
				pt.Escaped = false
				*pt.pop += `\`
				ansiReset(block[i])
			case pt.QuoteSingle, pt.QuoteDouble, pt.QuoteBrace > 0:
				*pt.pop += `\`
				syntaxHighlighted += string(block[i])
			default:
				pt.Escaped = true
				ansiColour(hlEscaped, block[i])
			}

		case '\'':
			pt.Loc = i
			switch {
			case pt.Escaped:
				pt.Escaped = false
				*pt.pop += `'`
				ansiReset(block[i])
			case pt.QuoteDouble, pt.QuoteBrace > 0:
				*pt.pop += `'`
				syntaxHighlighted += string(block[i])
			case pt.QuoteSingle:
				pt.QuoteSingle = false
				ansiReset(block[i])
			default:
				pt.QuoteSingle = true
				ansiColour(hlSingleQuote, block[i])
			}

		case '"':
			pt.Loc = i
			switch {
			case pt.Escaped:
				pt.Escaped = false
				*pt.pop += `"`
				ansiReset(block[i])
			case pt.QuoteSingle, pt.QuoteBrace > 0:
				*pt.pop += `"`
				syntaxHighlighted += string(block[i])
			case pt.QuoteDouble:
				pt.QuoteDouble = false
				ansiReset(block[i])
			default:
				pt.QuoteDouble = true
				ansiColour(hlDoubleQuote, block[i])
			}

		case '(':
			pt.Loc = i
			switch {
			case pt.Escaped:
				pt.Escaped = false
				*pt.pop += `(`
				ansiReset(block[i])
			case pt.QuoteSingle, pt.QuoteDouble:
				*pt.pop += `(`
				syntaxHighlighted += string(block[i])
			case pt.ExpectFunc:
				pt.ExpectFunc = false
				ansiColour(hlBraceQuote, block[i])
				pt.FuncName = "("
				pt.Parameters = append(pt.Parameters, "")
				pt.pop = &pt.Parameters[0]
				pt.QuoteBrace++
			case pt.QuoteBrace == 0:
				ansiColour(hlBraceQuote, block[i])
				pt.QuoteBrace++
			default:
				*pt.pop += `(`
				syntaxHighlighted += string(block[i])
				pt.QuoteBrace++
			}

		case ')':
			pt.Loc = i
			switch {
			case pt.Escaped:
				pt.Escaped = false
				*pt.pop += `)`
				ansiReset(block[i])
			case pt.QuoteSingle, pt.QuoteDouble:
				*pt.pop += `)`
				syntaxHighlighted += string(block[i])
			case pt.QuoteBrace == 1:
				ansiReset(block[i])
				pt.QuoteBrace--
			case pt.QuoteBrace == 0:
				ansiColour(hlError, block[i])
				pt.QuoteBrace--
			default:
				*pt.pop += `)`
				syntaxHighlighted += string(block[i])
				pt.QuoteBrace--
			}

		case ' ':
			switch {
			case pt.Escaped:
				pt.Escaped = false
				*pt.pop += ` `
				ansiReset(block[i])
			case pt.QuoteSingle, pt.QuoteDouble, pt.QuoteBrace > 0:
				*pt.pop += ` `
				syntaxHighlighted += string(block[i])
			case readFunc:
				pt.Loc = i
				pt.ExpectFunc = false
				readFunc = false
				pt.Parameters = append(pt.Parameters, "")
				pt.pop = &pt.Parameters[0]
				pt.Unsafe = isFuncUnsafe(pt.FuncName) || pt.Unsafe
				ansiReset(block[i])
			case pt.ExpectFunc:
				pt.Loc = i
				syntaxHighlighted += string(block[i])
			default:
				pt.Loc = i
				pt.Parameters = append(pt.Parameters, "")
				pt.pop = &pt.Parameters[len(pt.Parameters)-1]
				syntaxHighlighted += string(block[i])
			}

		case '>':
			switch {
			case pt.Escaped:
				pt.Escaped = false
				ansiReset(block[i])
			case pt.QuoteSingle, pt.QuoteDouble, pt.QuoteBrace > 0:
				*pt.pop += ` `
				syntaxHighlighted += string(block[i])
			case i > 0 && block[i-1] == '-':
				if pos != 0 && pt.Loc >= pos {
					return
				}
				pt.Loc = i
				pt.LastFlowToken = i - 1
				pt.ExpectFunc = true
				pt.pop = &pt.FuncName
				//pt.FuncName = ""
				pt.Parameters = make([]string, 0)
				syntaxHighlighted = syntaxHighlighted[:len(syntaxHighlighted)-1]
				ansiColour(hlPipe, '-')
				ansiReset('>')
				syntaxHighlighted += hlFunction
			case pt.ExpectFunc, readFunc:
				readFunc = true
				*pt.pop += `>`
				fallthrough
			default:
				pt.Loc = i
				syntaxHighlighted += string(block[i])
			}

		case ';', '|':
			pt.Loc = i
			switch {
			case pt.Escaped:
				pt.Escaped = false
				*pt.pop += string(block[i])
				ansiReset(block[i])
			case pt.QuoteSingle, pt.QuoteDouble, pt.QuoteBrace > 0:
				*pt.pop += string(block[i])
				syntaxHighlighted += string(block[i])
			default:
				if pos != 0 && pt.Loc >= pos {
					return
				}
				pt.LastFlowToken = i
				pt.ExpectFunc = true
				pt.pop = &pt.FuncName
				//pt.FuncName = ""
				pt.Parameters = make([]string, 0)
				ansiChar(hlPipe, block[i])
				syntaxHighlighted += hlFunction
			}

		case '\n':
			pt.Loc = i
			switch {
			case pt.Escaped:
				pt.Escaped = false
				*pt.pop += string(block[i])
				ansiReset(block[i])
			case pt.QuoteSingle, pt.QuoteDouble, pt.QuoteBrace > 0:
				*pt.pop += string(block[i])
				syntaxHighlighted += string(block[i])
			default:
				if pos != 0 && pt.Loc >= pos {
					return
				}
				pt.LastFlowToken = i
				pt.Unsafe = true
				pt.ExpectFunc = true
				pt.pop = &pt.FuncName
				//pt.FuncName = ""
				pt.Parameters = make([]string, 0)
				ansiChar(hlPipe, block[i])
				syntaxHighlighted += hlFunction
			}

		case '?':
			pt.Loc = i
			switch {
			case pt.Escaped:
				pt.Escaped = false
				*pt.pop += `?`
				ansiReset(block[i])
			case pt.QuoteSingle, pt.QuoteDouble, pt.QuoteBrace > 0:
				*pt.pop += `?`
				syntaxHighlighted += string(block[i])
			case i > 0 && block[i-1] == ' ':
				if pos != 0 && pt.Loc >= pos {
					return
				}
				pt.LastFlowToken = i
				pt.ExpectFunc = true
				pt.pop = &pt.FuncName
				pt.Parameters = make([]string, 0)
				pt.Unsafe = true
				ansiChar(hlPipe, block[i])
				syntaxHighlighted += hlFunction
			default:
				*pt.pop += `?`
				syntaxHighlighted += string(block[i])
			}

		case '{':
			pt.Loc = i
			switch {
			case pt.Escaped:
				pt.Escaped = false
				*pt.pop += `{`
				ansiReset(block[i])
			case pt.QuoteSingle, pt.QuoteDouble, pt.QuoteBrace > 0:
				*pt.pop += `{`
				syntaxHighlighted += string(block[i])
			default:
				pt.NestedBlock++
				pt.ExpectFunc = true
				pt.pop = &pt.FuncName
				pt.Parameters = make([]string, 0)
				//pt.Unsafe = true
				syntaxHighlighted += hlBlock + string(block[i])
			}

		case '}':
			switch {
			case pt.Escaped:
				pt.Escaped = false
				*pt.pop += `}`
				ansiReset(block[i])
			case pt.Escaped, pt.QuoteSingle, pt.QuoteDouble, pt.QuoteBrace > 0:
				*pt.pop += `}`
				syntaxHighlighted += string(block[i])
			default:
				pt.NestedBlock--
				//pt.Unsafe = true
				syntaxHighlighted += string(block[i])
				if pt.NestedBlock == 0 {
					syntaxHighlighted += ansi.Reset + reset[len(reset)-1]
				}
			}

		case '$':
			pt.VarLoc = i
			switch {
			case pt.Escaped:
				pt.Escaped = false
				*pt.pop += string(block[i])
				ansiReset(block[i])
			case pt.QuoteSingle:
				*pt.pop += string(block[i])
				syntaxHighlighted += string(block[i])
			default:
				pt.Unsafe = true
				*pt.pop += string(block[i])
				pt.Variable = string(block[i])
				ansiColour(hlVariable, block[i])
			}

		case '@':
			pt.VarLoc = i
			switch {
			case pt.Escaped:
				pt.Escaped = false
				*pt.pop += string(block[i])
				ansiReset(block[i])
			case pt.QuoteSingle:
				*pt.pop += string(block[i])
				syntaxHighlighted += string(block[i])
			default:
				pt.Unsafe = true
				*pt.pop += string(block[i])

				if i > 0 && (block[i-1] == ' ' || block[i-1] == '\t') {
					pt.Variable = string(block[i])
					ansiColour(hlVariable, block[i])
				} else {
					syntaxHighlighted += string(block[i])
				}
			}

		case ':':
			switch {
			case pt.Escaped:
				pt.Escaped = false
				*pt.pop += `:`
				ansiReset(block[i])
			case pt.QuoteSingle, pt.QuoteDouble, pt.QuoteBrace > 0:
				*pt.pop += `:`
				syntaxHighlighted += string(block[i])
			case !pt.ExpectFunc:
				*pt.pop += `:`
				syntaxHighlighted += string(block[i])
			case readFunc:
				pt.Loc = i
				pt.ExpectFunc = false
				readFunc = false
				pt.Parameters = append(pt.Parameters, "")
				pt.pop = &pt.Parameters[0]
				pt.Unsafe = isFuncUnsafe(pt.FuncName) || pt.Unsafe
				ansiReset(block[i])
			default:
				syntaxHighlighted += string(block[i])
			}

		case '<':
			switch {
			case pt.Escaped:
				pt.Escaped = false
				ansiReset(block[i])
			case readFunc:
				*pt.pop += string(block[i])
				syntaxHighlighted += string(block[i])
			case pt.ExpectFunc:
				*pt.pop = string(block[i])
				readFunc = true
				syntaxHighlighted += string(block[i])
			default:
				pt.Unsafe = true
				*pt.pop += string(block[i])
				syntaxHighlighted += string(block[i])
			}

		default:
			switch {
			case pt.Escaped:
				pt.Escaped = false
				ansiReset(block[i])
			case readFunc:
				*pt.pop += string(block[i])
				syntaxHighlighted += string(block[i])
			case pt.ExpectFunc:
				*pt.pop = string(block[i])
				readFunc = true
				syntaxHighlighted += string(block[i])
			default:
				*pt.pop += string(block[i])
				syntaxHighlighted += string(block[i])
			}
		}
	}
	pt.Loc++
	pt.VarLoc++
	syntaxHighlighted += ansi.Reset
	return
}

func isFuncUnsafe(f string) bool {
	for _, sb := range SafeFunctions {
		if f == sb {
			return false
		}
	}

	return true
}
