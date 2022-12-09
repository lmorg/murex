package expressions

import (
	"testing"

	"github.com/lmorg/murex/lang/expressions/symbols"
)

func TestParseQuoteParen(t *testing.T) {
	tests := expTestsT{
		symbol: symbols.QuoteDouble,
		tests: []expTestT{
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
				input:    `%(\s\t\r\n)`,
				expected: `\s\t\r\n`,
				pos:      1,
			},
		},
	}

	testParserSymbol(t, tests)
}
