package termemu_test

import (
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/termemu"
)

func TestParseBasic(t *testing.T) {
	count.Tests(t, 1)

	grid := termemu.NewGrid(5, 3)
	test := "Hello world!"
	exp := "Hello\n worl\nd!   \n"

	termemu.Parse(grid, []rune(test))
	act := grid.Export()

	if exp != act {
		t.Error("Expected output does not match actual output")
		t.Logf("  Expected: '%s'", exp)
		t.Logf("  Actual:   '%s'", act)
		t.Logf("  exp bytes: %v", []byte(exp))
		t.Logf("  act bytes: %v", []byte(act))
	}
}
