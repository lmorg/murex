package expressions

import (
	"testing"

	"github.com/lmorg/murex/lang/expressions/symbols"
)

func TestParseBareword(t *testing.T) {
	tests := expTestsT{
		symbol: symbols.Bareword,
		tests: []expTestT{
			{
				input:    `foobar`,
				expected: `foobar`,
			},
			{
				input:    `foobar  `,
				expected: `foobar`,
			},
			{
				input:    ` foobar`,
				expected: `foobar`,
				pos:      1,
			},
			{
				input:    `  foobar`,
				expected: `foobar`,
				pos:      2,
			},
			{
				input:    `   foobar`,
				expected: `foobar`,
				pos:      3,
			},
			{
				input:    "\tfoobar",
				expected: `foobar`,
				pos:      1,
			},
			{
				input:    "\t foobar",
				expected: `foobar`,
				pos:      2,
			},
			{
				input:    "\t\t  foobar",
				expected: `foobar`,
				pos:      4,
			},
			{
				input:    `  foobar  `,
				expected: `foobar`,
				pos:      2,
			},
			{
				input:    `foo bar`,
				expected: `foo`,
			},
			{
				input:    `foo-bar`,
				expected: `foo`,
			},
			{
				input:    `foo_bar`,
				expected: `foo_bar`,
			},
			{
				input:    `foo0bar`,
				expected: `foo0bar`,
			},
		},
	}

	testParserSymbol(t, tests)
}
