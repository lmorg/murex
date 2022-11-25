package expressions

import (
	"testing"

	"github.com/lmorg/murex/builtins/core/expressions/symbols"
)

func TestParseBoolean(t *testing.T) {
	tests := expTestsT{
		symbol: symbols.Boolean,
		tests: []expTestT{
			{
				input:    `true`,
				expected: `true`,
			},
			{
				input:    `true  `,
				expected: `true`,
			},
			{
				input:    ` true`,
				expected: `true`,
				pos:      1,
			},
			{
				input:    `  true`,
				expected: `true`,
				pos:      2,
			},
			{
				input:    `   true`,
				expected: `true`,
				pos:      3,
			},
			{
				input:    "\ttrue",
				expected: `true`,
				pos:      1,
			},
			{
				input:    "\t true",
				expected: `true`,
				pos:      2,
			},
			{
				input:    "\t\t  true",
				expected: `true`,
				pos:      4,
			},
			{
				input:    `  true  `,
				expected: `true`,
				pos:      2,
			},
			///
			{
				input:    `false`,
				expected: `false`,
			},
			{
				input:    `false  `,
				expected: `false`,
			},
			{
				input:    ` false`,
				expected: `false`,
				pos:      1,
			},
			{
				input:    `  false`,
				expected: `false`,
				pos:      2,
			},
			{
				input:    `   false`,
				expected: `false`,
				pos:      3,
			},
			{
				input:    "\tfalse",
				expected: `false`,
				pos:      1,
			},
			{
				input:    "\t false",
				expected: `false`,
				pos:      2,
			},
			{
				input:    "\t\t  false",
				expected: `false`,
				pos:      4,
			},
			{
				input:    `  false  `,
				expected: `false`,
				pos:      2,
			},
		},
	}

	testParserSymbol(t, tests)
}
