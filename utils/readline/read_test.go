package readline

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestRemoveNonPrintableChars(t *testing.T) {
	tests := []struct {
		Slice    string
		Expected string
	}{
		{
			Slice:    "",
			Expected: "",
		},
		{
			Slice:    "a",
			Expected: "a",
		},
		{
			Slice:    "abc",
			Expected: "abc",
		},
		{
			Slice:    "\t",
			Expected: "\t",
		},
		{
			Slice:    "\ta",
			Expected: "\ta",
		},
		{
			Slice:    "a\t",
			Expected: "a\t",
		},
		{
			Slice:    "a\tb",
			Expected: "a\tb",
		},
		{
			Slice:    "a\tb\tc",
			Expected: "a\tb\tc",
		},
		{
			Slice:    "a\t\tb\t\tc",
			Expected: "a\t\tb\t\tc",
		},

		// non printable

		{
			Slice:    "\x16",
			Expected: "",
		},
		{
			Slice:    "\x16a",
			Expected: "a",
		},
		{
			Slice:    "a\x16",
			Expected: "a",
		},
		{
			Slice:    "a\x16b",
			Expected: "ab",
		},
		{
			Slice:    "a\x16b\x16c",
			Expected: "abc",
		},
		{
			Slice:    "a\x16\x16b\x16\x16c",
			Expected: "abc",
		},

		// unicode

		{
			Slice:    "世界",
			Expected: "世界",
		},
		{
			Slice:    "\x16世\x16界\x16",
			Expected: "世界",
		},
		{
			Slice:    "\x16世界\x16世界\x16",
			Expected: "世界世界",
		},
		{
			Slice:    "\x16\x16世界\x16\x16世界\x16\x16",
			Expected: "世界世界",
		},
		{
			Slice:    "😀😁😂",
			Expected: "😀😁😂",
		},
		{
			Slice:    "\x16😀\x16😁\x16😂",
			Expected: "😀😁😂",
		},
		{
			Slice:    "\x16😀😁😂\x16😀😁😂\x16",
			Expected: "😀😁😂😀😁😂",
		},
		{
			Slice:    "\x16\x16😀😁😂\x16\x16😀😁😂\x16\x16",
			Expected: "😀😁😂😀😁😂",
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		s := []byte(test.Slice)
		actual := string(s[:removeNonPrintableChars(s)])

		if test.Expected != actual {
			t.Errorf("Expected does not match actual in test %d", i)
			t.Logf("  Slice:    '%s'", test.Slice)
			t.Logf("  Expected: '%s'", test.Expected)
			t.Logf("  Actual:   '%s'", actual)
			t.Logf("  s bytes:  '%v'", []byte(test.Slice))
			t.Logf("  e bytes:  '%v'", []byte(test.Expected))
			t.Logf("  a bytes:  '%v'", []byte(actual))
		}
	}
}
