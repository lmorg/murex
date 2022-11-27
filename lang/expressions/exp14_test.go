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

func TestExpAssignAdd(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `foo += 5`,
			Expected:   nil,
		},
		{
			Expression: `foo += "bar"`,
			Expected:   nil,
		},
		{
			Expression: `foo += bar`,
			Error:      true,
		},
		{
			Expression: `foo += true`,
			Error:      true,
		},
		{
			Expression: `foo += >`,
			Error:      true,
		},
	}

	testExpression(t, tests)
}

func TestExpAssignSubtract(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `foo -= 5`,
			Expected:   nil,
		},
		{
			Expression: `foo -= "bar"`,
			Error:      true,
		},
		{
			Expression: `foo -= bar`,
			Error:      true,
		},
		{
			Expression: `foo -= true`,
			Error:      true,
		},
		{
			Expression: `foo -= >`,
			Error:      true,
		},
	}

	testExpression(t, tests)
}

func TestExpAssignMultiply(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `foo *= 5`,
			Expected:   nil,
		},
		{
			Expression: `foo *= "bar"`,
			Error:      true,
		},
		{
			Expression: `foo *= bar`,
			Error:      true,
		},
		{
			Expression: `foo *= true`,
			Error:      true,
		},
		{
			Expression: `foo *= >`,
			Error:      true,
		},
	}

	testExpression(t, tests)
}

func TestExpAssignDivide(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `foo /= 5`,
			Expected:   nil,
		},
		{
			Expression: `foo /= "bar"`,
			Error:      true,
		},
		{
			Expression: `foo /= bar`,
			Error:      true,
		},
		{
			Expression: `foo /= true`,
			Error:      true,
		},
		{
			Expression: `foo /= >`,
			Error:      true,
		},
	}

	testExpression(t, tests)
}
