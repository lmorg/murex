package expressions

import "testing"

func TestExpAssign(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `foo = 5`,
			Expected:   nil,
		},
		{
			Expression: `foo = "bar"`,
			Expected:   nil,
		},
		{
			Expression: `foo = bar`,
			Error:      true,
		},
		{
			Expression: `foo = >`,
			Error:      true,
		},
	}

	testExpression(t, tests)
}
