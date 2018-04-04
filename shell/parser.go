package shell

import (
	"github.com/lmorg/murex/utils/parser"
)

func syntaxHighlight(r []rune) string {
	_, highlighted := parse(r)
	return highlighted
}

func parse(line []rune) (pt parser.ParsedTokens, syntaxHighlighted string) {
	return parser.Parse(line, 0)
}
