package parser

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestIsFuncUnsafe(t *testing.T) {
	trues := []string{">", ">>", "$var", "@g", "config"}
	falses := []string{"open", "regexp", "match", "cat", "cat", "format", "[", "[[", "runtime"}

	count.Tests(t, len(trues)+len(falses))

	for i := range trues {
		v := isFuncUnsafe(trues[i])
		if v != true {
			t.Errorf("Returned `%s` expected `%s`: '%s'", "false", "true", trues[i])
		}
	}

	for i := range falses {
		v := isFuncUnsafe(falses[i])
		if v != false {
			t.Errorf("Returned `%s` expected `%s`: '%s'", "true", "false", falses[i])
		}
	}
}
