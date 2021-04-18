package types

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestTrimSpaceRune(t *testing.T) {
	tests := map[string]string{
		"":      "",
		" ":     "",
		" a":    "a",
		"a ":    "a",
		"  a  ": "a",
	}

	count.Tests(t, len(tests))

	for in, exp := range tests {
		r := trimSpaceRune([]rune(in))
		act := string(r)

		if exp != act {
			t.Error("expected != actual")
			t.Logf("  Input:    '%s'", in)
			t.Logf("  Expected: '%s", exp)
			t.Logf("  Actual:   '%s'", act)
		}
	}
}

func TestIsBlockRune(t *testing.T) {
	tests := map[string]bool{
		"":         false,
		" {":       false,
		" }":       false,
		" { ":      false,
		" } ":      false,
		" {}":      true,
		" {} ":     true,
		" { } ":    true,
		"{ test }": true,
	}

	count.Tests(t, len(tests))

	for block, exp := range tests {
		act := IsBlockRune([]rune(block))

		if exp != act {
			t.Error("expected != actual")
			t.Logf("  Input:    '%s'", block)
			t.Logf("  Expected: '%v", exp)
			t.Logf("  Actual:   '%v'", act)
		}
	}
}
