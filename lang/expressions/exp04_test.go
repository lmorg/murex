package expressions

import "testing"

func TestExpAdd(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `"foo" + "bar"`,
			Expected:   `foobar`,
		},
		{
			Expression: `"foo"+ "bar"`,
			Expected:   `foobar`,
		},
		///
		{
			Expression: `1 + 2`,
			Expected:   float64(3),
		},
		{
			Expression: `1+ 2`,
			Expected:   float64(3),
		},
		///
		{
			Expression: `1 + "2"`,
			Error:      true,
		},
	}

	testExpression(t, tests)
}

func TestExpSubtract(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `"foo" - "bar"`,
			Error:      true,
		},
		///
		{
			Expression: `1 - 2`,
			Expected:   float64(-1),
		},
		{
			Expression: `-1 - 2`,
			Expected:   float64(-3),
		},
		{
			Expression: `1 - -2`,
			Expected:   float64(3),
		},
		{
			Expression: `-1 - -2`,
			Expected:   float64(1),
		},
		{
			Expression: `-1- -2`,
			Expected:   float64(1),
		},
		{
			Expression: `-1--2`,
			Expected:   float64(1),
		},
		///
		{
			Expression: `1 - "2"`,
			Error:      true,
		},
	}

	testExpression(t, tests)
}
