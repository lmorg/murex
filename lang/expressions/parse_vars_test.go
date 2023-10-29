package expressions

import (
	"testing"

	"github.com/lmorg/murex/lang/expressions/symbols"
)

func TestParseVarsScalarSymbol(t *testing.T) {
	tests := expTestsT{
		symbol: symbols.Scalar,
		tests: []expTestT{
			{
				input:    `$foo[1]`,
				expected: `$foo[1]`,
				pos:      -1,
			},
			{
				input:    `$foo[[/1]]`,
				expected: `$foo[[/1]]`,
				pos:      -1,
			},
		},
	}

	testParserSymbol(t, tests)
}

/*func TestParseVarsArraySymbol(t *testing.T) {
	tests := expTestsT{
		symbol: symbols.Calculated,
		tests: []expTestT{
			{
				input:    `@foo[1]`,
				expected: `@foo[1]`,
				pos:      6,
			},
			{
				input:    `@foo[[/1]]`,
				expected: `@foo[[/1]]`,
				pos:      6,
			},
		},
	}

	testParserSymbol(t, tests)
}
*/
