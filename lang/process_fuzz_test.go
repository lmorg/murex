//go:build go1.18
// +build go1.18

package lang_test

import (
	"testing"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test/count"
)

func FuzzParseBlock(f *testing.F) {
	tests := []string{"out: hello world", "", "bg { err: abc 123 }"}
	for _, tc := range tests {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		count.Tests(t, 1)
		lang.DontCacheAst = true
		lang.ParseBlock([]rune(orig))
		// we are just testing we can't cause an unhandled panic
	})
}
