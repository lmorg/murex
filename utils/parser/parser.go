package parser

//go:generate stringer -type=PipeToken

import (
	"regexp"

	"github.com/lmorg/murex/utils/ansi/codes"
)

// syntax highlighting
var (
	hlFunction    = codes.Bold
	hlVariable    = codes.FgGreen
	hlEscaped     = codes.FgYellow
	hlSingleQuote = codes.FgBlue
	hlDoubleQuote = codes.FgBlue
	hlBraceQuote  = codes.FgBlue
	hlBlock       = []string{codes.FgGreen, codes.FgMagenta, codes.FgBlue, codes.FgYellow}
	hlPipe        = codes.FgMagenta
	hlComment     = codes.FgGreen + codes.Invert
	hlError       = codes.FgRed + codes.Invert
	hlRedirect    = codes.FgGreen

	rxAllowedVarChars = regexp.MustCompile(`^[._a-zA-Z0-9]$`)
)

// ParsedTokens is a struct that returns a tokenized version of the selected command
type ParsedTokens struct {
	Source        []rune
	LastCharacter rune
	Loc           int
	VarLoc        int
	VarBrace      bool
	VarSigil      string
	Escaped       bool
	Comment       bool
	CommentMsg    string
	commentMsg    []rune
	QuoteSingle   bool
	QuoteDouble   bool
	QuoteBrace    int
	NestedBlock   int
	SquareBracket bool
	AngledBracket bool
	ExpectFunc    bool
	ExpectParam   bool
	pop           *string
	LastFuncName  string
	FuncName      string
	Parameters    []string
	Unsafe        bool // if the pipeline is estimated to be safe enough to dynamically preview
	LastFlowToken int
	PipeToken     PipeToken
}

// PipeToken stores an integer value for the pipe token used in a pipeline
type PipeToken int

// These are different pipe tokens
const (
	PipeTokenNone     PipeToken = 0    // No pipe token
	PipeTokenPosix    PipeToken = iota // `|`  (POSIX style pipe)
	PipeTokenArrow                     // `->` (murex style pipe)
	PipeTokenGeneric                   // `=>` (reformat to generic)
	PipeTokenRedirect                  // `?`  (STDERR redirected to STDOUT and vice versa)
	PipeTokenAppend                    // `>>` (append STDOUT to a file)
)

// Parse a single line of code and return the tokens for a selected command
func Parse(block []rune, pos int) (pt ParsedTokens, syntaxHighlighted string) {
	var readFunc bool
	reset := []string{codes.Reset, hlFunction}
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
	}

	ansiResetNoChar := func() {
		if len(reset) > 1 {
			reset = reset[:len(reset)-1]
		}
		syntaxHighlighted += reset[len(reset)-1]
	}

	ansiChar := func(colour string, r ...rune) {
		syntaxHighlighted += colour + string(r) + reset[len(reset)-1]
	}

	ansiStartFunction := func() {
		ansiResetNoChar()
		syntaxHighlighted += hlFunction
	}

	var i int

	expectParam := func() {
		pt.ExpectParam = false
		pt.Parameters = append(pt.Parameters, "")
		pt.pop = &pt.Parameters[len(pt.Parameters)-1]
	}

	escaped := func() {
		pt.Escaped = false
		*pt.pop += string(block[i])
		ansiReset(block[i])
	}

	next := func(r rune) bool {
		if i+1 < len(block) {
			return block[i+1] == r
		}
		return false
	}

	for ; i < len(block); i++ {
		if pt.Comment {
			pt.commentMsg = append(pt.commentMsg, block[i])
			continue
		}

		if !pt.Escaped {
			pt.LastCharacter = block[i]
		}

		if pt.VarSigil != "" {
			if !pt.VarBrace {
				if !rxAllowedVarChars.MatchString(string(block[i])) {
					pt.VarSigil = ""
					ansiResetNoChar()
				}
			} else {
				*pt.pop += string(block[i])
				syntaxHighlighted += string(block[i])
				if block[i] == ')' {
					pt.VarSigil = ""
					ansiResetNoChar()
				}
				continue
			}
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
				syntaxHighlighted += hlComment + string(block[i:]) + codes.Reset
				//return
				defer func() { pt.CommentMsg = string(pt.commentMsg) }()
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
				if pt.ExpectParam {
					expectParam()
				}
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
				pt.ExpectParam = true
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

		case '=':
			switch {
			case pt.Escaped:
				escaped()
			case pt.QuoteSingle, pt.QuoteDouble, pt.QuoteBrace > 0, readFunc:
				*pt.pop += `=`
				syntaxHighlighted += string(block[i])
			case pt.ExpectFunc:
				pt.Loc = i
				syntaxHighlighted += string(block[i])
			default:
				pt.Loc = i
				syntaxHighlighted += string(block[i])
				pt.ExpectParam = true
			}

		case ':':
			switch {
			case pt.Escaped:
				escaped()
			case pt.QuoteSingle, pt.QuoteDouble, pt.QuoteBrace > 0, pt.SquareBracket:
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
				pt.LastFuncName = pt.FuncName
				pt.Parameters = make([]string, 0)
				syntaxHighlighted = syntaxHighlighted[:len(syntaxHighlighted)-1]
				ansiColour(hlPipe, block[i-1])
				ansiReset('>')
				syntaxHighlighted += hlFunction
			case i > 0 && (block[i-1] == '\t' || block[i-1] == ' ') && next('>'):
				if pos != 0 && pt.Loc >= pos {
					return
				}
				i++
				pt.Loc = i
				pt.LastFlowToken = i - 1
				pt.Unsafe = true
				pt.ExpectFunc = false
				readFunc = false
				pt.ExpectParam = true
				pt.SquareBracket = false
				pt.PipeToken = PipeTokenAppend
				pt.FuncName = ">>"
				pt.Parameters = make([]string, 0)
				ansiColour(hlPipe, '>')
				ansiReset('>')
				syntaxHighlighted += hlRedirect
			case pt.ExpectFunc, readFunc:
				readFunc = true
				*pt.pop += `>`
				pt.Loc = i
				syntaxHighlighted += ">"
			case pt.AngledBracket:
				*pt.pop += `>`
				pt.Loc = i
				syntaxHighlighted += ">" + codes.Reset

			default:
				pt.Loc = i
				syntaxHighlighted += ">"
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
				pt.LastFuncName = pt.FuncName
				pt.Parameters = make([]string, 0)
				if next('>') {
					*pt.pop += `>`
					pt.LastFlowToken = i - 1
					pt.Unsafe = true
					pt.ExpectFunc = false
					readFunc = false
					pt.ExpectParam = true
					pt.SquareBracket = false
					i++
					if next('>') {
						pt.Loc = i
						i++
						ansiChar(hlPipe, '|', '>', '>')
						pt.PipeToken = PipeTokenAppend
					} else {
						ansiChar(hlPipe, '|', '>')
					}

					syntaxHighlighted += hlRedirect
				} else {
					ansiChar(hlPipe, block[i])
					ansiStartFunction()
				}
			}

		case '&':
			pt.Loc = i
			switch {
			case pt.Escaped:
				escaped()
			case pt.QuoteSingle, pt.QuoteDouble, pt.QuoteBrace > 0:
				*pt.pop += string(block[i])
				syntaxHighlighted += string(block[i])
			case next('&'):
				if pos != 0 && pt.Loc >= pos {
					return
				}
				pt.LastFlowToken = i
				pt.ExpectFunc = true
				pt.SquareBracket = false
				pt.PipeToken = PipeTokenNone
				pt.pop = &pt.FuncName
				pt.LastFuncName = pt.FuncName
				pt.Parameters = make([]string, 0)
				ansiChar(hlPipe, '&', '&')
				ansiStartFunction()
				i++
			default:
				*pt.pop += string(block[i])
				syntaxHighlighted += string(block[i])
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
				pt.LastFuncName = pt.FuncName
				pt.Parameters = make([]string, 0)
				ansiChar(hlPipe, block[i])
				ansiStartFunction()
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
				pt.LastFuncName = pt.FuncName
				pt.Parameters = make([]string, 0)
				ansiChar(hlPipe, block[i])
				ansiStartFunction()
			}

		case '?':
			pt.Loc = i
			switch {
			case pt.Escaped:
				escaped()
			case pt.QuoteSingle, pt.QuoteDouble, pt.QuoteBrace > 0:
				*pt.pop += "?"
				syntaxHighlighted += string(block[i])
			case next(':'):
				if pos != 0 && pt.Loc >= pos {
					return
				}
				pt.LastFlowToken = i
				pt.ExpectFunc = true
				pt.SquareBracket = false
				pt.PipeToken = PipeTokenNone
				pt.pop = &pt.FuncName
				pt.LastFuncName = pt.FuncName
				pt.Parameters = make([]string, 0)
				ansiChar(hlPipe, '?', ':')
				ansiStartFunction()
				i++
			case i > 0 && block[i-1] == ' ':
				if pos != 0 && pt.Loc >= pos {
					return
				}
				pt.LastFlowToken = i
				pt.ExpectFunc = true
				pt.SquareBracket = false
				pt.PipeToken = PipeTokenRedirect
				pt.pop = &pt.FuncName
				pt.LastFuncName = pt.FuncName
				pt.Parameters = make([]string, 0)
				pt.Unsafe = true
				ansiChar(hlPipe, block[i])
				syntaxHighlighted += hlFunction
			default:
				*pt.pop += `?`
				syntaxHighlighted += "?"
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
				if pt.NestedBlock >= 0 {
					i := pt.NestedBlock % len(hlBlock)
					syntaxHighlighted += hlBlock[i] + "{" + codes.Reset + hlFunction
				} else {
					syntaxHighlighted += hlError + "{"
				}
			}

		case '}':
			switch {
			case pt.Escaped:
				escaped()
			case pt.QuoteSingle, pt.QuoteDouble, pt.QuoteBrace > 0:
				*pt.pop += `}`
				syntaxHighlighted += "}"
			default:
				if pt.NestedBlock >= 1 {
					i := pt.NestedBlock % len(hlBlock)
					syntaxHighlighted += hlBlock[i] + "}" + codes.Reset
				} else {
					syntaxHighlighted += hlError + "}"
				}
				pt.NestedBlock--
				if pt.NestedBlock == 0 {
					syntaxHighlighted += reset[len(reset)-1]
				}
			}

		case '[':
			switch {
			case pt.Escaped:
				escaped()
			case readFunc:
				*pt.pop += string(block[i])
				syntaxHighlighted += string(block[i])
				pt.SquareBracket = true
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
			switch {
			case pt.Escaped:
				escaped()
			case pt.QuoteSingle:
				*pt.pop += string(block[i])
				syntaxHighlighted += string(block[i])
			case pt.ExpectParam:
				expectParam()
				fallthrough
			default:
				pt.Unsafe = true
				*pt.pop += string(block[i])
				pt.VarSigil = string(block[i])
				ansiColour(hlVariable, block[i])
				if next('(') {
					pt.VarLoc = i + 1
					pt.VarBrace = true
				} else {
					pt.VarLoc = i
					pt.VarBrace = false
				}
			}

		case '@':
			switch {
			case pt.Escaped:
				escaped()
			case pt.QuoteSingle, next(' '), next('\t'):
				*pt.pop += string(block[i])
				syntaxHighlighted += string(block[i])
			case pt.ExpectParam:
				expectParam()
				fallthrough
			default:
				pt.Unsafe = true
				*pt.pop += string(block[i])
				pt.VarSigil = string(block[i])
				ansiColour(hlVariable, block[i])
				if next('(') {
					pt.VarLoc = i + 1
					pt.VarBrace = true
				} else {
					pt.VarLoc = i
					pt.VarBrace = false
				}
			}

		case '<':
			switch {
			case pt.Escaped:
				escaped()
			case readFunc:
				*pt.pop += "<"
				syntaxHighlighted += "<"
			case pt.ExpectFunc:
				*pt.pop = "<"
				readFunc = true
				syntaxHighlighted += "<"
			default:
				pt.Unsafe = true
				*pt.pop += "<"
				syntaxHighlighted += hlRedirect + "<"
				pt.AngledBracket = true
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
	syntaxHighlighted += codes.Reset
	return
}
