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

	testExpression(t, tests, true)
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

	testExpression(t, tests, true)
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

	testExpression(t, tests, true)
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

	testExpression(t, tests, true)
}

// TestExpGtLtArrayVsNumber is a regression test for
// https://github.com/lmorg/murex/issues/982 where comparing an
// (uncomparable) array against a number panicked with an interface
// conversion error instead of returning a clean result.
func TestExpGtLtArrayVsNumber(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `%[a b c] < 2`,
			Expected:   false,
		},
		{
			Expression: `%[a b c] > 2`,
			Expected:   true,
		},
		{
			Expression: `%[a b c] >= 2`,
			Expected:   true,
		},
	}

	testExpression(t, tests, false)
}
