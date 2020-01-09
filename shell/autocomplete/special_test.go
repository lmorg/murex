package autocomplete

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestIsSpecialBuiltin(t *testing.T) {
	trues := []string{">", ">>", "[", "![", "[[", "@[", "=", "(", "!", ".", "@g"}
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

func TestSortColon(t *testing.T) {
	test := []string{
		"cc", "c", "bb:", "b:", "dd", "d", "aa:", "a:",
	}

	expected := []string{
		"a:", "aa:", "b:", "bb:", "c", "cc", "d", "dd",
	}

	sortColon(test, 0, len(test)-1)

	passed := true
	for i := range test {
		passed = passed && test[i] == expected[i]
	}

	if !passed {
		t.Error("Expected splice does not match actual splice")
		t.Log("  Expected:", expected)
		t.Log("  Actual:  ", test)
	}
}
