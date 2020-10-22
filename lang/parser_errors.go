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
	ErrUnexpectedPipeTokenPipe
	ErrUnexpectedPipeTokenEqGt
	ErrUnexpectedPipeTokenQm
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

	ErrPipingToNothing
)

var errMessages = []string{
	"No errors. Block successfully parsed",
	"Unquoted or unescaped colon located in function parameters",
	"Pipe token, `|`, preceding function name",
	"Pipe token, `=>`, preceding function name",
	"Pipe token, `?`, preceding function name",
	"Unquoted or unescaped opening curly brace in function parameters",
	"Unquoted or unescaped closinging curly brace in function parameters",
	"Unexpected closing curly brace, `}`, as no matching opening curly brace found",
	"Unexpected closing quotation brace, `)`, as no matching opening quotation brace found",
	"Unexpected end of script. Escape token used but with no character escaped",
	"Unexpected end of script. Single quotes not closed",
	"Unexpected end of script. Double quotes not closed",
	"Unexpected end of script. More open curly braces, `{`, than closed",
	"Unexpected end of script. More open quotation braces, `(`, than closed",
	"Unexpected end of script. Variable index used, `[`, but missing closing bracket: `]`",
	"Unexpected parsing error",
	"Unexpected opening curly brace. Code blocks cannot be used as function names",

	"Piping out to nothing. Commands should not be terminated by a pipe token (`|`, `->`, `=>`, or ` ?`)",
}

func raiseErr(code, endByte int) ParserError {
	return ParserError{
		Message: errMessages[code],
		Code:    code,
		EndByte: endByte,
	}
}
