package expressions

import "testing"

func TestExpMultiplyStrict(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `"foo" * "bar"`,
			Error:      true,
		},
		{
			Expression: `"foo"* "bar"`,
			Error:      true,
		},
		///
		{
			Expression: `1 * 2`,
			Expected:   float64(2),
		},
		{
			Expression: `1* 2`,
			Expected:   float64(2),
		},
		{
			Expression: `1*-2`,
			Expected:   float64(-2),
		},
		{
			Expression: `1* -2`,
			Expected:   float64(-2),
		},
		{
			Expression: `1 *-2`,
			Expected:   float64(-2),
		},
		{
			Expression: `1 * -2`,
			Expected:   float64(-2),
		},
		///
		{
			Expression: `1 * "2"`,
			Error:      true,
		},
	}

	testExpression(t, tests, true)
}

func TestExpDivide(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `"foo" / "bar"`,
			Error:      true,
		},
		{
			Expression: `"foo"/ "bar"`,
			Error:      true,
		},
		///
		{
			Expression: `1 / 2`,
			Expected:   float64(0.5),
		},
		{
			Expression: `1/ 2`,
			Expected:   float64(0.5),
		},
		{
			Expression: `1/-2`,
			Expected:   float64(-0.5),
		},
		{
			Expression: `1 /-2`,
			Expected:   float64(-0.5),
		},
		{
			Expression: `1/ -2`,
			Expected:   float64(-0.5),
		},
		{
			Expression: `1 / -2`,
			Expected:   float64(-0.5),
		},
		///
		{
			Expression: `1 / "2"`,
			Error:      true,
		},
	}

	testExpression(t, tests, true)
}
