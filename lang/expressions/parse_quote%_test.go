package expressions

import (
	"testing"

	"github.com/lmorg/murex/lang/expressions/symbols"
)

func TestParseQuoteParen(t *testing.T) {
	tests := expTestsT{
		symbol: symbols.QuoteParenthesis,
		tests: []expTestT{
			{
				input:    `%()`,
				expected: ``,
				pos:      1,
			},
			{
				input: `%(foobar`,
				error: true,
			},
			{
				input:    `%(foobar)`,
				expected: `foobar`,
				pos:      1,
			},
			{
				input:    `%(foobar)  `,
				expected: `foobar`,
				pos:      1,
			},
			{
				input:    ` %(foobar)`,
				expected: `foobar`,
				pos:      2,
			},
			{
				input:    `  %(foobar)`,
				expected: `foobar`,
				pos:      3,
			},
			{
				input:    `   %(foobar)`,
				expected: `foobar`,
				pos:      4,
			},
			{
				input:    "\t%(foobar)",
				expected: `foobar`,
				pos:      2,
			},
			{
				input:    "\t %(foobar)",
				expected: `foobar`,
				pos:      3,
			},
			{
				input:    "\t\t  %(foobar)",
				expected: `foobar`,
				pos:      5,
			},
			{
				input:    `  %(foobar)  `,
				expected: `foobar`,
				pos:      3,
			},
			{
				input:    `%(foo bar)`,
				expected: `foo bar`,
				pos:      1,
			},
			{
				input:    `%(foobar')`,
				expected: `foobar'`,
				pos:      1,
			},
			{
				input:    `%(foo-$bar-bar)`,
				expected: `foo--bar`,
				pos:      5,
			},
			{
				input:    `%(foo-$(bar)-bar)`,
				expected: `foo--bar`,
				pos:      7,
			},
			{
				input:    `%(foo-$b-bar)`,
				expected: `foo--bar`,
				pos:      3,
			},
			{
				input:    `%(foo-\$bar-bar)`,
				expected: `foo-\-bar`,
				pos:      5,
			},
			{
				input:    `%(foo-\$(bar)-bar)`,
				expected: `foo-\-bar`,
				pos:      7,
			},
			{
				input:    `%(\foo-\$bar-\bar)`,
				expected: `\foo-\-\bar`,
				pos:      5,
			},
			{
				input:    `%(foo-@bar-bar)`,
				expected: `foo-@bar-bar`,
				pos:      1,
			},
			{
				input:    `%(\s)`,
				expected: `\s`,
				pos:      1,
			},
			{
				input:    `%(\t)`,
				expected: `\t`,
				pos:      1,
			},
			{
				input:    `%(\r)`,
				expected: `\r`,
				pos:      1,
			},
			{
				input:    `%(\n)`,
				expected: `\n`,
				pos:      1,
			},
		},
	}

	testParserSymbol(t, tests)
}
