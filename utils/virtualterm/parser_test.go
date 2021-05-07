package virtualterm_test

import (
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/virtualterm"
)

func TestParseBasic(t *testing.T) {
	count.Tests(t, 1)

	term := virtualterm.NewTerminal(5, 3)
	test := "Hello world!"
	exp := "Hello\n worl\nd!   \n"

	virtualterm.Parse(term, []rune(test))
	act := term.Export()

	if exp != act {
		t.Error("Expected output does not match actual output")
		t.Logf("  Expected: '%s'", exp)
		t.Logf("  Actual:   '%s'", act)
		t.Logf("  exp bytes: %v", []byte(exp))
		t.Logf("  act bytes: %v", []byte(act))
	}
}

func TestParseNewLineCrLf1(t *testing.T) {
	count.Tests(t, 1)

	term := virtualterm.NewTerminal(5, 3)
	test := "1\n2\n3\n4\n5"
	exp := "   4 \n    5\n     \n"

	virtualterm.Parse(term, []rune(test))
	act := term.Export()

	if exp != act {
		t.Error("Expected output does not match actual output")
		t.Logf("  Expected: '%s'", exp)
		t.Logf("  Actual:   '%s'", act)
		t.Logf("  exp bytes: %v", []byte(exp))
		t.Logf("  act bytes: %v", []byte(act))
	}
}

func TestParseNewLineCrLf2(t *testing.T) {
	count.Tests(t, 1)

	term := virtualterm.NewTerminal(5, 3)
	test := "1\r\n2\r\n3\r\n4\r\n5"
	exp := "3    \n4    \n5    \n"

	virtualterm.Parse(term, []rune(test))
	act := term.Export()

	if exp != act {
		t.Error("Expected output does not match actual output")
		t.Logf("  Expected: '%s'", exp)
		t.Logf("  Actual:   '%s'", act)
		t.Logf("  exp bytes: %v", []byte(exp))
		t.Logf("  act bytes: %v", []byte(act))
	}
}

func TestParseNewLineLf(t *testing.T) {
	count.Tests(t, 1)

	term := virtualterm.NewTerminal(5, 3)
	term.LfIncCr = true
	test := "1\n2\n3\n4\n5"
	exp := "3    \n4    \n5    \n"

	virtualterm.Parse(term, []rune(test))
	act := term.Export()

	if exp != act {
		t.Error("Expected output does not match actual output")
		t.Logf("  Expected: '%s'", exp)
		t.Logf("  Actual:   '%s'", act)
		t.Logf("  exp bytes: %v", []byte(exp))
		t.Logf("  act bytes: %v", []byte(act))
	}
}
