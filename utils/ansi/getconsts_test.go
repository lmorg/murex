package ansi_test

import (
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/ansi"
)

func TestGetConsts(t *testing.T) {
	tests := []string{
		"{ESC}", "{^C}", "{EOF}", "{BELL}", "{F12}", "{CURSOR-UP}",
		"{ESC}-1", "{ESC}a", "{ESC}$", "{ESC}£",
		"{ESC}12", "{ESC}ab", "{ESC}$%", "{ESC}😀😁",
		"1{ESC}1", "a{ESC}a", "${ESC}$", "£{ESC}£",
		"12{ESC}12", "ab{ESC}ab", "$%{ESC}$%", "😀😁{ESC}😀😁",
		"{ESC} ", "{ESC}  ",
		" ", "1", "a", "$", "£", "😀",
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		seq := ansi.ForceExpandConsts(tests[i], false)
		actual := ansi.GetConsts([]byte(seq))
		if test != actual {
			t.Errorf("Incorrect constants quoted in test %d:", i)
			t.Logf("  Expected: '%s'", test)
			t.Logf("  Actual:   '%s'", actual)
		}
	}
}
