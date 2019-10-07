package autocomplete

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestIsSpecialBuiltin(t *testing.T) {
	trues := []string{">", ">>", "[", "[[", "@[", "="}
	falses := []string{"", "and", "or", "if", "foobar", "0", "123"}

	count.Tests(t, len(trues)+len(falses))

	for i := range trues {
		v := isSpecialBuiltin(trues[i])
		if v != true {
			t.Errorf("Returned `%s` expected `%s`: '%s'", "false", "true", trues[i])
		}
	}

	for i := range falses {
		v := isSpecialBuiltin(falses[i])
		if v != false {
			t.Errorf("Returned `%s` expected `%s`: '%s'", "true", "false", falses[i])
		}
	}
}
