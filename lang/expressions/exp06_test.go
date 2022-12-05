package expressions

import "testing"

func TestExpGreaterThan(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `"foobar" > "foobar"`,
			Expected:   false,
		},
		{
			Expression: `"foo" > "bar"`,
			Expected:   true,
		},
		{
			Expression: `"bar" > "foo"`,
			Expected:   false,
		},
		///
		{
			Expression: `1 > 1`,
			Expected:   false,
		},
		{
			Expression: `1 > 2`,
			Expected:   false,
		},
		{
			Expression: `2 > 1`,
			Expected:   true,
		},
		///
		{
			Expression: `1 > "1"`,
			Error:      true,
		},
		{
			Expression: `1 > "2"`,
			Error:      true,
		},
		{
			Expression: `2 > "1"`,
			Error:      true,
		},
	}

	testExpression(t, tests)
}

func TestExpGreaterThanOrEqual(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `"foobar" >= "foobar"`,
			Expected:   true,
		},
		{
			Expression: `"foo" >= "bar"`,
			Expected:   true,
		},
		{
			Expression: `"bar" >= "foo"`,
			Expected:   false,
		},
		///
		{
			Expression: `1 >= 1`,
			Expected:   true,
		},
		{
			Expression: `1 >= 2`,
			Expected:   false,
		},
		{
			Expression: `2 >= 1`,
			Expected:   true,
		},
		///
		{
			Expression: `1 >= "1"`,
			Error:      true,
		},
		{
			Expression: `1 >= "2"`,
			Error:      true,
		},
		{
			Expression: `2 >= "1"`,
			Error:      true,
		},
	}

	testExpression(t, tests)
}

func TestExpLessThan(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `"foobar" < "foobar"`,
			Expected:   false,
		},
		{
			Expression: `"foo" < "bar"`,
			Expected:   false,
		},
		{
			Expression: `"bar" < "foo"`,
			Expected:   true,
		},
		///
		{
			Expression: `1 < 1`,
			Expected:   false,
		},
		{
			Expression: `1 < 2`,
			Expected:   true,
		},
		{
			Expression: `2 < 1`,
			Expected:   false,
		},
		///
		{
			Expression: `1 < "1"`,
			Error:      true,
		},
		{
			Expression: `1 < "2"`,
			Error:      true,
		},
		{
			Expression: `2 < "1"`,
			Error:      true,
		},
	}

	testExpression(t, tests)
}

func TestExpLessThanOrEqual(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `"foobar" <= "foobar"`,
			Expected:   true,
		},
		{
			Expression: `"foo" <= "bar"`,
			Expected:   false,
		},
		{
			Expression: `"bar" <= "foo"`,
			Expected:   true,
		},
		///
		{
			Expression: `1 <= 1`,
			Expected:   true,
		},
		{
			Expression: `1 <= 2`,
			Expected:   true,
		},
		{
			Expression: `2 <= 1`,
			Expected:   false,
		},
		///
		{
			Expression: `1 <= "1"`,
			Error:      true,
		},
		{
			Expression: `1 <= "2"`,
			Error:      true,
		},
		{
			Expression: `2 <= "1"`,
			Error:      true,
		},
	}

	testExpression(t, tests)
}
