package expressions

import (
	"testing"
)

func TestExpElvis(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `null ?: null`,
			Expected:   nil,
		},
		{
			Expression: `null ?: "null"`,
			Expected:   "null",
		},
		{
			Expression: `null ?: 3`,
			Expected:   3,
		},
		{
			Expression: `null ?: "3"`,
			Expected:   "3",
		},
		{
			Expression: `null ?: %[1..3]`,
			Expected:   "[1,2,3]",
		},
	}

	testExpression(t, tests, true)
}
