package expressions

import (
	"testing"

	"github.com/lmorg/murex/builtins/core/expressions/symbols"
)

func TestParseNumber(t *testing.T) {
	tests := expTestsT{
		symbol: symbols.Number,
		tests: []expTestT{
			{
				input:    `0`,
				expected: `0`,
			},
			{
				input:    `0  `,
				expected: `0`,
			},
			{
				input:    `12  `,
				expected: `12`,
			},
			{
				input:    ` 0`,
				expected: `0`,
				pos:      1,
			},
			{
				input:    ` 12`,
				expected: `12`,
				pos:      1,
			},
			{
				input:    `  0`,
				expected: `0`,
				pos:      2,
			},
			{
				input:    `   0`,
				expected: `0`,
				pos:      3,
			},
			{
				input:    "\t0",
				expected: `0`,
				pos:      1,
			},
			{
				input:    "\t 0",
				expected: `0`,
				pos:      2,
			},
			{
				input:    "\t\t  0",
				expected: `0`,
				pos:      4,
			},
			{
				input:    `  0  `,
				expected: `0`,
				pos:      2,
			},
			{
				input:    `  123  `,
				expected: `123`,
				pos:      2,
			},
			{
				input:    `0 0`,
				expected: `0`,
			},
			///
			{
				input:    `0.1`,
				expected: `0.1`,
			},
			{
				input:    `012`,
				expected: `012`,
			},
		},
	}

	testParserSymbol(t, tests)
}
