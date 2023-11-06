package expressions

import (
	"testing"

	"github.com/lmorg/murex/lang/expressions/symbols"
)

func TestParseLambdaScalar(t *testing.T) {
	tests := expTestsT{
		symbol: symbols.Scalar,
		tests: []expTestT{
			{
				input:    `$foo[{bar}]`,
				expected: `$foo[{bar}]`,
				pos:      -1,
			},
		},
	}

	testParserSymbol(t, tests)
}

/*func TestParseLambdaArray(t *testing.T) {
	tests := expTestsT{
		symbol: symbols.Calculated,
		tests: []expTestT{
			{
				input:    `@foo[{bar}]`,
				expected: `@foo[{bar}]`,
				pos:      4,
			},
		},
	}

	testParserSymbol(t, tests)
}
*/
