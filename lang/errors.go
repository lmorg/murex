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
	ErrPipingToNothing
	ErrUnknownParserErrorPipe
	ErrUnableToParseParametersInRunmode
	ErrInvalidParametersInRunmode
)

var errMessages = []string{
	"No errors. Block successfully parsed",
	"Piping out to nothing. Commands should not be terminated by a pipe token (`|`, `->`, `=>`, or ` ?`)",
	"Unexpected error parsing `|`. Reason unknown. Please file a bug at https://github.com/lmorg/murex/issues",
	"Unable to parse parameters in `runmode`",
	"Invalid parameters in `runmode`",
}
