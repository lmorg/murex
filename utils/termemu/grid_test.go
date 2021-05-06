package termemu

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestWriteCell(t *testing.T) {
	count.Tests(t, 1)

	test := "Hello world!"
	exp := "Hello\n worl\nd!   \n"

	grid := NewGrid(5, 3)

	for _, r := range []rune(test) {
		grid.writeCell(r)
	}
	act := grid.Export()

	if exp != act {
		t.Error("Expected output does not match actual output")
		t.Logf("  Expected: '%s'", exp)
		t.Logf("  Actual:   '%s'", act)
		t.Logf("  exp bytes: %v", []byte(exp))
		t.Logf("  act bytes: %v", []byte(act))
	}
}

func TestMoveGridUp(t *testing.T) {
	count.Tests(t, 1)

	test := "1234554321"
	exp := "54321\n     \n     \n"

	grid := NewGrid(5, 3)

	for _, r := range []rune(test) {
		grid.writeCell(r)
	}
	grid.moveGridUp()

	act := grid.Export()

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

	grid := NewGrid(5, 3)
	grid.cells[1][4].char = 'T'
	grid.curPos.X = 4
	grid.curPos.Y = 1

	c := grid.Cell().char
	if c != 'T' {
		t.Errorf("*Grid.Cell look up appears to fail. %d", c)
	}
}

func TestMoveCursorForwards(t *testing.T) {
	count.Tests(t, 1)

	grid := NewGrid(5, 3)
	overflow := grid.moveCursorForwards(5)

	if grid.curPos.X != grid.size.X-1 {
		t.Errorf("curPos.X not where expected:")
		t.Logf("  Expected: %d", grid.size.X-1)
		t.Logf("  Actual:   %d", grid.curPos.X)
	}

	if overflow != 1 {
		t.Errorf("overflow value does not match expected:")
		t.Logf("  Expected: %d", 1)
		t.Logf("  Actual:   %d", overflow)
	}
}

func TestMoveCursorDownwards(t *testing.T) {
	count.Tests(t, 1)

	grid := NewGrid(5, 3)
	overflow := grid.moveCursorDownwards(3)

	if grid.curPos.Y != grid.size.Y-1 {
		t.Errorf("curPos.Y not where expected:")
		t.Logf("  Expected: %d", grid.size.Y-1)
		t.Logf("  Actual:   %d", grid.curPos.Y)
	}

	if overflow != 1 {
		t.Errorf("overflow value does not match expected:")
		t.Logf("  Expected: %d", 1)
		t.Logf("  Actual:   %d", overflow)
	}
}
