package expressions

import (
	"testing"

	"github.com/lmorg/murex/lang/expressions/symbols"
)

func TestParseQuoteSingle(t *testing.T) {
	tests := expTestsT{
		symbol: symbols.QuoteSingle,
		tests: []expTestT{
			{
				input:    `''`,
				expected: ``,
			},
			{
				input: `'foobar`,
				error: true,
			},
			{
				input:    `'foobar'`,
				expected: `foobar`,
			},
			{
				input:    `'foobar'  `,
				expected: `foobar`,
			},
			{
				input:    ` 'foobar'`,
				expected: `foobar`,
				pos:      1,
			},
			{
				input:    `  'foobar'`,
				expected: `foobar`,
				pos:      2,
			},
			{
				input:    `   'foobar'`,
				expected: `foobar`,
				pos:      3,
			},
			{
				input:    "\t'foobar'",
				expected: `foobar`,
				pos:      1,
			},
			{
				input:    "\t 'foobar'",
				expected: `foobar`,
				pos:      2,
			},
			{
				input:    "\t\t  'foobar'",
				expected: `foobar`,
				pos:      4,
			},
			{
				input:    `  'foobar'  `,
				expected: `foobar`,
				pos:      2,
			},
			{
				input:    `'foo bar'`,
				expected: `foo bar`,
			},
			{
				input:    `'foobar"'`,
				expected: `foobar"`,
			},
			{
				input:    `'foobar\'`,
				expected: `foobar\`,
			},
			{
				input: `'foobar\''`,
				error: true,
			},
			{
				input:    `'foo-$bar-bar'`,
				expected: `foo-$bar-bar`,
			},
			{
				input:    `'foo-@bar-bar'`,
				expected: `foo-@bar-bar`,
			},
			{
				input:    `'\s\t\r\n'`,
				expected: `\s\t\r\n`,
			},
		},
	}

	testParserSymbol(t, tests)
}
