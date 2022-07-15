package parser

import (
	"fmt"
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/ansi/codes"
)

func TestParserHighlightEscaped(t *testing.T) {
	type T struct {
		Block    string
		Expected string
	}

	tests := []T{
		{
			Block: `out: foo \n bar`,
			Expected: fmt.Sprintf(
				`%sout:%s foo %s\n%s bar%s`,
				hlFunction, codes.Reset, hlEscaped, codes.Reset, codes.Reset,
			),
		},
		{
			Block: `out: 'foo \n bar'`,
			Expected: fmt.Sprintf(
				`%sout:%s %s'foo \n bar'%s%s`,
				hlFunction, codes.Reset, hlSingleQuote, codes.Reset, codes.Reset,
			),
		},
		{
			Block: `out: "foo \n bar"`,
			Expected: fmt.Sprintf(
				`%sout:%s %s"foo %s\n%s bar"%s%s`,
				hlFunction, codes.Reset, hlDoubleQuote, hlEscaped, hlDoubleQuote, codes.Reset, codes.Reset,
			),
		},
		{
			Block: `out: (foo \n bar)`,
			Expected: fmt.Sprintf(
				`%sout:%s %s(foo \n bar)%s%s`,
				hlFunction, codes.Reset, hlBraceQuote, codes.Reset, codes.Reset,
			),
		},
		/////
		{
			Block: `out: foo\nbar`,
			Expected: fmt.Sprintf(
				`%sout:%s foo%s\n%sbar%s`,
				hlFunction, codes.Reset, hlEscaped, codes.Reset, codes.Reset,
			),
		},
		{
			Block: `out: 'foo\nbar'`,
			Expected: fmt.Sprintf(
				`%sout:%s %s'foo\nbar'%s%s`,
				hlFunction, codes.Reset, hlSingleQuote, codes.Reset, codes.Reset,
			),
		},
		{
			Block: `out: "foo\nbar"`,
			Expected: fmt.Sprintf(
				`%sout:%s %s"foo%s\n%sbar"%s%s`,
				hlFunction, codes.Reset, hlDoubleQuote, hlEscaped, hlDoubleQuote, codes.Reset, codes.Reset,
			),
		},
		{
			Block: `out: (foo\nbar)`,
			Expected: fmt.Sprintf(
				`%sout:%s %s(foo\nbar)%s%s`,
				hlFunction, codes.Reset, hlBraceQuote, codes.Reset, codes.Reset,
			),
		},
		/////
		{
			Block: `out: "foo\nb"`,
			Expected: fmt.Sprintf(
				`%sout:%s %s"foo%s\n%sb"%s%s`,
				hlFunction, codes.Reset, hlDoubleQuote, hlEscaped, hlDoubleQuote, codes.Reset, codes.Reset,
			),
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		_, highlight := Parse([]rune(test.Block), 0)
		if highlight != test.Expected {
			t.Errorf("Expected != Actual in test %d", i)
			t.Logf("  Block:    '%s'", test.Block)
			t.Logf("  Expected: '%s'", test.Expected)
			t.Logf("  Actual:   '%s'", highlight)
			t.Log("  exp byte: ", []byte(test.Expected))
			t.Log("  act byte: ", []byte(highlight))
		}
	}
}
