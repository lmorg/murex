package expressions

import "testing"

func TestExpressions(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `1 + 1 + 1 + 1 + 1 + 1`,
			Expected:   float64(6),
		},
		{
			Expression: `1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1`,
			Expected:   float64(10),
		},
		{
			Expression: `1 + 2 * 3`,
			Expected:   float64(7),
		},
		{
			Expression: `1 + 2 * 3 == 7`,
			Expected:   true,
		},
		{
			Expression: `(1 + 2) * 3`,
			Expected:   float64(9),
		},
		{
			Expression: `(1 + 2) * 3 == 7`,
			Expected:   false,
		},
		{
			Expression: `(1 + 2) * 3 == 9`,
			Expected:   true,
		},
		{
			Expression: `7 == 1 + 2 * 3`,
			Expected:   true,
		},
		{
			Expression: `7 == 4+4`,
			Expected:   false,
		},
		{
			Expression: `7 == 4 + 4`,
			Expected:   false,
		},
		{
			Expression: `7 == 4 + 4 - 1`,
			Expected:   true,
		},
		{
			Expression: `(7 == 7) == true`,
			Expected:   true,
		},
		{
			Expression: `(7 == 5) == true`,
			Expected:   false,
		},
		{
			Expression: `(7 == 7) == false`,
			Expected:   false,
		},
		{
			Expression: `(7 == 5) == false`,
			Expected:   true,
		},
		{
			Expression: `'bob' == "bob"`,
			Expected:   true,
		},
		{
			Expression: `"bob" == "bob"`,
			Expected:   true,
		},
		{
			Expression: `"bob" == 'bob'`,
			Expected:   true,
		},
		{
			Expression: `'bob' == 'bob'`,
			Expected:   true,
		},
	}

	testExpression(t, tests)
}
