package lang

type ParserError struct {
	Message string
	Code    int
	EndByte int // this is sometimes useful to know
}

const (
	NoParsingErrors = 0 + iota
	ErrUnexpectedColon
	ErrUnexpectedPipeToken
	ErrUnexpectedOpenBrace
	ErrUnexpectedCloseBrace
	ErrClosingBraceNoOpen
	ErrUnterminatedEscape
	ErrUnterminatedQuotesSingle
	ErrUnterminatedQuotesDouble
	ErrUnterminatedBrace
)

var errMessages map[int]string = map[int]string{
	0: "No errors. Block successfully parsed.",
	1: "Unquoted or unescaped colon located in function parameters.",
	2: "Pipe token preceding function name.",
	3: "Unquoted or unescaped opening brace in function parameters.",
	4: "Unquoted or unescaped closinging brace in function parameters.",
	5: "Unexpected closing brace as no matching opening brace found.",
	6: "Unexpected end of script. Escape token used but with no character escaped.",
	7: "Unexpected end of script. Single quotes not closed.",
	8: "Unexpected end of script. Double quotes not closed.",
	9: "Unexpected end of script. More open braces than closed.",
}

func raiseErr(code, endByte int) ParserError {
	return ParserError{
		Message: errMessages[code],
		Code:    code,
		EndByte: endByte,
	}
}
