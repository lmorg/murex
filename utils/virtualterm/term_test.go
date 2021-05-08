package virtualterm

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestWriteCell(t *testing.T) {
	count.Tests(t, 1)

	test := "Hello world!"
	exp := "Hello\n worl\nd!   \n"

	term := NewTerminal(5, 3)

	for _, r := range []rune(test) {
		term.writeCell(r)
	}
	act := term.Export()

	if exp != act {
		t.Error("Expected output does not match actual output")
		t.Logf("  Expected: '%s'", exp)
		t.Logf("  Actual:   '%s'", act)
		t.Logf("  exp bytes: %v", []byte(exp))
		t.Logf("  act bytes: %v", []byte(act))
	}
}

func TestMoveContentsUp(t *testing.T) {
	count.Tests(t, 1)

	test := "1234554321"
	exp := "54321\n     \n     \n"

	term := NewTerminal(5, 3)

	for _, r := range []rune(test) {
		term.writeCell(r)
	}
	term.moveContentsUp()

	act := term.Export()

	if exp != act {
		t.Error("Expected output does not match actual output")
		t.Logf("  Expected: '%s'", exp)
		t.Logf("  Actual:   '%s'", act)
		t.Logf("  exp bytes: %v", []byte(exp))
		t.Logf("  act bytes: %v", []byte(act))
	}
}

func TestCell(t *testing.T) {
	count.Tests(t, 1)

	term := NewTerminal(5, 3)
	term.cells[1][4].char = 'T'
	term.curPos.X = 4
	term.curPos.Y = 1

	c := term.cell().char
	if c != 'T' {
		t.Errorf("*Term.Cell look up appears to fail. %d", c)
	}
}

func TestMoveCursorForwards(t *testing.T) {
	count.Tests(t, 1)

	term := NewTerminal(5, 3)
	overflow := term.moveCursorForwards(5)

	if term.curPos.X != term.size.X-1 {
		t.Errorf("curPos.X not where expected:")
		t.Logf("  Expected: %d", term.size.X-1)
		t.Logf("  Actual:   %d", term.curPos.X)
	}

	if overflow != 1 {
		t.Errorf("overflow value does not match expected:")
		t.Logf("  Expected: %d", 1)
		t.Logf("  Actual:   %d", overflow)
	}
}

func TestMoveCursorDownwards(t *testing.T) {
	count.Tests(t, 1)

	term := NewTerminal(5, 3)
	overflow := term.moveCursorDownwards(3)

	if term.curPos.Y != term.size.Y-1 {
		t.Errorf("curPos.Y not where expected:")
		t.Logf("  Expected: %d", term.size.Y-1)
		t.Logf("  Actual:   %d", term.curPos.Y)
	}

	if overflow != 1 {
		t.Errorf("overflow value does not match expected:")
		t.Logf("  Expected: %d", 1)
		t.Logf("  Actual:   %d", overflow)
	}
}
