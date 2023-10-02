package shell

import (
	"github.com/lmorg/murex/utils/parser"
)

/*func SyntaxHighlight(r []rune) string {
	_, highlighted := parse(r)
	return highlighted
}*/

var SyntaxHighlight func([]rune) string

func parse(line []rune) (pt parser.ParsedTokens, syntaxHighlighted string) {
	return parser.Parse(line, 0)
}
