package expressions

import (
	"testing"

	"github.com/lmorg/murex/lang/expressions/symbols"
)

func TestParseQuoteDouble(t *testing.T) {
	tests := expTestsT{
		symbol: symbols.QuoteDouble,
		tests: []expTestT{
			{
				input:    `""`,
				expected: ``,
			},
			{
				input: `"foobar`,
				error: true,
			},
			{
				input:    `"foobar"`,
				expected: `foobar`,
			},
			{
				input:    `"foobar"  `,
				expected: `foobar`,
			},
			{
				input:    ` "foobar"`,
				expected: `foobar`,
				pos:      1,
			},
			{
				input:    `  "foobar"`,
				expected: `foobar`,
				pos:      2,
			},
			{
				input:    `   "foobar"`,
				expected: `foobar`,
				pos:      3,
			},
			{
				input:    "\t\"foobar\"",
				expected: `foobar`,
				pos:      1,
			},
			{
				input:    "\t \"foobar\"",
				expected: `foobar`,
				pos:      2,
			},
			{
				input:    "\t\t  \"foobar\"",
				expected: `foobar`,
				pos:      4,
			},
			{
				input:    `  "foobar"  `,
				expected: `foobar`,
				pos:      2,
			},
			{
				input:    `"foo bar"`,
				expected: `foo bar`,
			},
			{
				input:    `"foobar'"`,
				expected: `foobar'`,
			},
			{
				input:    `"foo-$bar-bar"`,
				expected: `foo--bar`,
				pos:      4,
			},
			{
				input:    `"foo-$b-bar"`,
				expected: `foo--bar`,
				pos:      2,
			},
			{
				input:    `"foo-\$bar-bar"`,
				expected: `foo-$bar-bar`,
				pos:      1,
			},
			{
				input:    `"\foo-\$bar-\bar"`,
				expected: `foo-$bar-bar`,
				pos:      3,
			},
			{
				input:    `"foo-@bar-bar"`,
				expected: `foo-@bar-bar`,
			},
			{
				input:    `"\s\t\r\n"`,
				expected: " \t\r\n",
				pos:      4,
			},
		},
	}

	testParserSymbol(t, tests)
}
