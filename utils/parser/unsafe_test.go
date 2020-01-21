package parser

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestIsCmdUnsafe(t *testing.T) {
	trues := []string{">", ">>", "$var", "@g", "config"}
	falses := append(safeCmds,
		"open", "regexp", "match",
		"cast", "format", "[", "[[",
		"runtime",
	)

	count.Tests(t, len(trues)+len(falses))

	for i := range trues {
		v := isCmdUnsafe(trues[i])
		if v != true {
			t.Errorf("Returned `%s` expected `%s`: '%s'", "false", "true", trues[i])
		}
	}

	for i := range falses {
		v := isCmdUnsafe(falses[i])
		if v != false {
			t.Errorf("Returned `%s` expected `%s`: '%s'", "true", "false", falses[i])
		}
	}
}
