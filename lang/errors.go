package lang

// ParserError is the error object used for the murex parser
type ParserError struct {
	Message string
	Code    int
	EndByte int // this is sometimes useful to know
}

// murex script parsing error codes:
const (
	NoParsingErrors = 0 + iota
	ErrUnexpectedColon
	ErrUnexpectedPipeToken
	ErrUnexpectedOpenBrace
	ErrUnexpectedCloseBrace
	ErrClosingBraceBlockNoOpen
	ErrClosingBraceQuoteNoOpen
	ErrUnterminatedEscape
	ErrUnterminatedQuotesSingle
	ErrUnterminatedQuotesDouble
	ErrUnterminatedBraceBlock
	ErrUnterminatedBraceQuote
	ErrUnclosedIndex
	ErrUnexpectedParsingError
	ErrUnexpectedOpenBraceFunc
)

var errMessages = map[int]string{
	0:  "No errors. Block successfully parsed.",
	1:  "Unquoted or unescaped colon located in function parameters.",
	2:  "Pipe token preceding function name.",
	3:  "Unquoted or unescaped opening curly brace in function parameters.",
	4:  "Unquoted or unescaped closinging curly brace in function parameters.",
	5:  "Unexpected closing curly brace, `}`, as no matching opening curly brace found.",
	6:  "Unexpected closing quotation brace, `)`, as no matching opening quotation brace found.",
	7:  "Unexpected end of script. Escape token used but with no character escaped.",
	8:  "Unexpected end of script. Single quotes not closed.",
	9:  "Unexpected end of script. Double quotes not closed.",
	10: "Unexpected end of script. More open curly braces, `{`, than closed.",
	11: "Unexpected end of script. More open quotation braces, `(`, than closed.",
	12: "Unexpected end of script. Variable index used, `[`, but missing closing bracket: `]`.",
	13: "Unexpected parsing error.",
	14: "Unexpected opening curly brace. Code blocks cannot be used as function names.",
}

func raiseErr(code, endByte int) ParserError {
	return ParserError{
		Message: errMessages[code],
		Code:    code,
		EndByte: endByte,
	}
}
