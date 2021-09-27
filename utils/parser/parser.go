package parser

//go:generate stringer -type=PipeToken

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
	LastCharacter rune
	Loc           int
	VarLoc        int
	Escaped       bool
	Comment       bool
	QuoteSingle   bool
	QuoteDouble   bool
	QuoteBrace    int
	NestedBlock   int
	SquareBracket bool
	ExpectFunc    bool
	ExpectParam   bool
	pop           *string
	LastFuncName  string
	FuncName      string
	Parameters    []string
	Variable      string
	Unsafe        bool // if the pipeline is estimated to be safe enough to dynamically preview
	LastFlowToken int
	PipeToken     PipeToken
}

// PipeToken stores an interger value for the pipe token used in a pipeline
type PipeToken int

// These are different pipe tokens
const (
	PipeTokenNone     PipeToken = 0    // No pipe token
	PipeTokenPosix    PipeToken = iota // `|`  (POSIX style pipe)
	PipeTokenArrow                     // `->` (murex style pipe)
	PipeTokenGeneric                   // `=>` (reformat to generic)
	PipeTokenRedirect                  // `?`  (STDERR redirected to STDOUT and vice versa)
)

// Parse a single line of code and return the tokens for a selected command
func Parse(block []rune, pos int) (pt ParsedTokens, syntaxHighlighted string) {
	var readFunc bool
	reset := []string{ansi.Reset, hlFunction}
	syntaxHighlighted = hlFunction
	pt.Loc = -1
	pt.ExpectFunc = true
	pt.pop = &pt.FuncName
	pt.Source = block
	pt.Parameters = []string{}

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

	var i int

	expectParam := func() {
		pt.ExpectParam = false
		pt.Parameters = append(pt.Parameters, "")
		pt.pop = &pt.Parameters[len(pt.Parameters)-1]
		//syntaxHighlighted += string(block[i])
	}

	escaped := func() {
		pt.Escaped = false
		*pt.pop += string(block[i])
		ansiReset(block[i])
	}

	for i = range block {
		if !pt.Escaped {
			pt.LastCharacter = block[i]
		}

		if pt.Variable != "" && !rxAllowedVarChars.MatchString(string(block[i])) {
			pt.Variable = ""
			ansiResetNoChar()
		}

		switch block[i] {
		case '#':
			pt.Loc = i
			switch {
			case pt.Escaped:
				escaped()
			case pt.QuoteSingle, pt.QuoteDouble, pt.QuoteBrace > 0, pt.NestedBlock > 0:
				*pt.pop += `#`
				syntaxHighlighted += string(block[i])
			case pt.ExpectParam:
				fallthrough
			default:
				pt.Comment = true
				syntaxHighlighted += hlComment + string(block[i:]) + ansi.Reset
				return
			}

		case '\\':
			switch {
			case pt.QuoteSingle, pt.QuoteBrace > 0:
				*pt.pop += `\`
				syntaxHighlighted += string(block[i])
			case pt.Escaped:
				escaped()
			case pt.ExpectParam:
				expectParam()
				fallthrough
			default:
				pt.Escaped = true
				ansiColour(hlEscaped, block[i])
			}

		case '\'':
			pt.Loc = i
			switch {
			case pt.Escaped:
				escaped()
			case pt.QuoteDouble, pt.QuoteBrace > 0:
				*pt.pop += `'`
				syntaxHighlighted += string(block[i])
			case pt.QuoteSingle:
				pt.QuoteSingle = false
				ansiReset(block[i])
			case pt.ExpectParam:
				expectParam()
				fallthrough
			default:
				pt.QuoteSingle = true
				ansiColour(hlSingleQuote, block[i])
			}

		case '"':
			pt.Loc = i
			switch {
			case pt.Escaped:
				escaped()
			case pt.QuoteSingle, pt.QuoteBrace > 0:
				*pt.pop += `"`
				syntaxHighlighted += string(block[i])
			case pt.QuoteDouble:
				pt.QuoteDouble = false
				ansiReset(block[i])
			case pt.ExpectParam:
				expectParam()
				fallthrough
			default:
				pt.QuoteDouble = true
				ansiColour(hlDoubleQuote, block[i])
			}

		case '(':
			pt.Loc = i
			switch {
			case pt.Escaped:
				escaped()
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
			case pt.ExpectParam:
				expectParam()
				fallthrough
			default:
				*pt.pop += `(`
				syntaxHighlighted += string(block[i])
				pt.QuoteBrace++
			}

		case ')':
			pt.Loc = i
			switch {
			case pt.Escaped:
				escaped()
			case pt.QuoteSingle, pt.QuoteDouble:
				*pt.pop += `)`
				syntaxHighlighted += string(block[i])
			case pt.QuoteBrace == 1:
				ansiReset(block[i])
				pt.QuoteBrace--
			case pt.QuoteBrace == 0:
				ansiColour(hlError, block[i])
				pt.QuoteBrace--
			case pt.ExpectParam:
				expectParam()
				fallthrough
			default:
				*pt.pop += `)`
				syntaxHighlighted += string(block[i])
				pt.QuoteBrace--
			}

		case ' ':
			switch {
			case pt.Escaped:
				escaped()
			case pt.QuoteSingle, pt.QuoteDouble, pt.QuoteBrace > 0:
				*pt.pop += ` `
				syntaxHighlighted += string(block[i])
			case readFunc:
				pt.Loc = i
				pt.ExpectFunc = false
				readFunc = false
				pt.Parameters = append(pt.Parameters, "")
				pt.pop = &pt.Parameters[0]
				pt.Unsafe = isCmdUnsafe(pt.FuncName) || pt.Unsafe
				ansiReset(block[i])
			case pt.ExpectFunc:
				pt.Loc = i
				syntaxHighlighted += string(block[i])
			case i > 0 && block[i-1] == ' ':
				pt.Loc = i
				syntaxHighlighted += " "
			case i > 0 && block[i-1] == ':' && len(pt.Parameters) == 1:
				pt.Loc = i
				syntaxHighlighted += " "
			default:
				pt.Loc = i
				syntaxHighlighted += string(block[i])
				pt.ExpectParam = true
			}

		case ':':
			switch {
			case pt.Escaped:
				escaped()
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
				pt.Unsafe = isCmdUnsafe(pt.FuncName) || pt.Unsafe
				ansiReset(block[i])
			default:
				syntaxHighlighted += string(block[i])
			}

		case '>':
			switch {
			case pt.Escaped:
				escaped()
			case pt.QuoteSingle, pt.QuoteDouble, pt.QuoteBrace > 0:
				*pt.pop += ` `
				syntaxHighlighted += string(block[i])
			case i > 0 && (block[i-1] == '-' || block[i-1] == '='):
				if pos != 0 && pt.Loc >= pos {
					return
				}
				pt.Loc = i
				pt.LastFlowToken = i - 1
				pt.ExpectFunc = true
				pt.SquareBracket = false
				if block[i-1] == '-' {
					pt.PipeToken = PipeTokenArrow
				} else {
					pt.PipeToken = PipeTokenGeneric
				}
				pt.pop = &pt.FuncName
				//pt.FuncName = ""
				pt.LastFuncName = pt.FuncName
				pt.Parameters = make([]string, 0)
				syntaxHighlighted = syntaxHighlighted[:len(syntaxHighlighted)-1]
				//ansiColour(hlPipe, '-')
				ansiColour(hlPipe, block[i-1])
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

		case '|':
			pt.Loc = i
			switch {
			case pt.Escaped:
				escaped()
			case pt.QuoteSingle, pt.QuoteDouble, pt.QuoteBrace > 0:
				*pt.pop += string(block[i])
				syntaxHighlighted += string(block[i])
			default:
				if pos != 0 && pt.Loc >= pos {
					return
				}
				pt.LastFlowToken = i
				pt.ExpectFunc = true
				pt.SquareBracket = false
				pt.PipeToken = PipeTokenPosix
				pt.pop = &pt.FuncName
				//pt.FuncName = ""
				pt.LastFuncName = pt.FuncName
				pt.Parameters = make([]string, 0)
				ansiChar(hlPipe, block[i])
				syntaxHighlighted += hlFunction
			}

		case ';':
			pt.Loc = i
			switch {
			case pt.Escaped:
				escaped()
			case pt.QuoteSingle, pt.QuoteDouble, pt.QuoteBrace > 0:
				*pt.pop += string(block[i])
				syntaxHighlighted += string(block[i])
			default:
				if pos != 0 && pt.Loc >= pos {
					return
				}
				pt.LastFlowToken = i
				pt.ExpectFunc = true
				pt.SquareBracket = false
				pt.PipeToken = PipeTokenNone
				pt.pop = &pt.FuncName
				//pt.FuncName = ""
				pt.LastFuncName = pt.FuncName
				pt.Parameters = make([]string, 0)
				ansiChar(hlPipe, block[i])
				syntaxHighlighted += hlFunction
			}

		case '\n':
			pt.Loc = i
			switch {
			case pt.Escaped:
				escaped()
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
				pt.SquareBracket = false
				pt.PipeToken = PipeTokenNone
				pt.pop = &pt.FuncName
				//pt.FuncName = ""
				pt.LastFuncName = pt.FuncName
				pt.Parameters = make([]string, 0)
				ansiChar(hlPipe, block[i])
				syntaxHighlighted += hlFunction
			}

		case '?':
			pt.Loc = i
			switch {
			case pt.Escaped:
				escaped()
			case pt.QuoteSingle, pt.QuoteDouble, pt.QuoteBrace > 0:
				*pt.pop += `?`
				syntaxHighlighted += string(block[i])
			case i > 0 && block[i-1] == ' ':
				if pos != 0 && pt.Loc >= pos {
					return
				}
				pt.LastFlowToken = i
				pt.ExpectFunc = true
				pt.SquareBracket = false
				pt.PipeToken = PipeTokenRedirect
				pt.pop = &pt.FuncName
				//pt.FuncName = ""
				pt.LastFuncName = pt.FuncName
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
				escaped()
			case pt.QuoteSingle, pt.QuoteDouble, pt.QuoteBrace > 0:
				*pt.pop += `{`
				syntaxHighlighted += string(block[i])
			default:
				pt.NestedBlock++
				pt.ExpectFunc = true
				pt.PipeToken = PipeTokenNone
				pt.pop = &pt.FuncName
				pt.Parameters = make([]string, 0)
				//pt.Unsafe = true
				syntaxHighlighted += hlBlock + string(block[i])
			}

		case '}':
			switch {
			case pt.Escaped:
				escaped()
			case pt.QuoteSingle, pt.QuoteDouble, pt.QuoteBrace > 0:
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

		case '[':
			switch {
			case pt.Escaped:
				escaped()
			case readFunc:
				*pt.pop += string(block[i])
				syntaxHighlighted += string(block[i])
				//if i > 0 && block[0] == '^' {
				pt.SquareBracket = true
				//}
			case pt.ExpectFunc:
				*pt.pop = string(block[i])
				readFunc = true
				syntaxHighlighted += string(block[i])
				pt.SquareBracket = true
			default:
				*pt.pop += string(block[i])
				syntaxHighlighted += string(block[i])
				pt.SquareBracket = true
			}

		case ']':
			switch {
			case pt.Escaped:
				escaped()
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
				pt.SquareBracket = true
			}

		case '$':
			pt.VarLoc = i
			switch {
			case pt.Escaped:
				escaped()
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
				escaped()
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

		case '<':
			switch {
			case pt.Escaped:
				escaped()
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
				switch block[i] {
				case 'r':
					*pt.pop = "\r"
				case 'n':
					*pt.pop = "\n"
				case 's':
					*pt.pop = " "
				case 't':
					*pt.pop = "\t"
				default:
					*pt.pop = string(block[i])
				}
				//syntaxHighlighted += string(block[i])
			case readFunc:
				*pt.pop += string(block[i])
				syntaxHighlighted += string(block[i])
			case pt.ExpectFunc:
				*pt.pop = string(block[i])
				readFunc = true
				syntaxHighlighted += string(block[i])
			case pt.ExpectParam:
				expectParam()
				fallthrough
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
