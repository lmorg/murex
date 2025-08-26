package parser_test

import (
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/parser"
)

func FuzzParser(f *testing.F) {
	tests := []string{"", "out: hello world", "bg { err: abc 123 }", "bob -> ? | =>"}
	for _, tc := range tests {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		count.Tests(t, 1)
		parser.Parse([]rune(orig), len(orig))
		// we are just testing we can't cause an unhandled panic
	})
}
