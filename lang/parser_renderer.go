package lang

// https://play.golang.com/p/yaQY5kAJCv5

const (
	prEscaped = 1 << iota
	prQuoteSingle
	prQuoteDouble
	prQuoteBrace
	prNestedBlock
	prSquareBracket
	prExpectFunc
	prExpectParam

	prUnsafe
)

type ParserRenderer struct {
	Syntax               []rune
	PreviousFunctionName string
	FunctionName         string
	// Parameters []string
	// Variable string
	// LastFlowToken int
	//PipeToken     PipeToken

	state int
}

func (pr *ParserRenderer) Excaped() bool          { return pr.state&prEscaped != 0 }
func (pr *ParserRenderer) QuoteSingle() bool      { return pr.state&prQuoteSingle != 0 }
func (pr *ParserRenderer) QuoteDouble() bool      { return pr.state&prQuoteDouble != 0 }
func (pr *ParserRenderer) QuoteBrace() bool       { return pr.state&prQuoteBrace != 0 }
func (pr *ParserRenderer) NestedBlock() bool      { return pr.state&prNestedBlock != 0 }
func (pr *ParserRenderer) SquareBracket() bool    { return pr.state&prSquareBracket != 0 }
func (pr *ParserRenderer) ExpectFunction() bool   { return pr.state&prExpectFunc != 0 }
func (pr *ParserRenderer) ExpectParameters() bool { return pr.state&prExpectParam != 0 }
func (pr *ParserRenderer) Unsafe() bool           { return pr.state&prUnsafe != 0 }
