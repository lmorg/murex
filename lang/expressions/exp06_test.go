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

// TestExpCompareNonComparableStrict is a regression test for
// https://github.com/lmorg/murex/issues/982
// Comparing a non-comparable type (e.g. array) with a number
// should return an error, not panic.
func TestExpCompareNonComparableStrict(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `%[1,2,3] < 2`,
			Error:      true,
		},
		{
			Expression: `2 < %[1,2,3]`,
			Error:      true,
		},
		{
			Expression: `%[1,2,3] > 2`,
			Error:      true,
		},
		{
			Expression: `%[1,2,3] <= 2`,
			Error:      true,
		},
		{
			Expression: `%[1,2,3] >= 2`,
			Error:      true,
		},
	}

	testExpression(t, tests, true)
}

// TestExpCompareNonComparableNonStrict is a regression test for
// https://github.com/lmorg/murex/issues/982
// In non-strict mode, comparing a non-comparable type with a number
// previously panicked with "interface conversion: interface {} is float64,
// not string". After the fix, this should not panic.
func TestExpCompareNonComparableNonStrict(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `%[1,2,3] < 2`,
			Expected:   false,
		},
		{
			Expression: `2 > %[1,2,3]`,
			Expected:   true,
		},
		{
			Expression: `%[1,2,3] > 2`,
			Expected:   true,
		},
		{
			Expression: `2 < %[1,2,3]`,
			Expected:   false,
		},
	}

	testExpression(t, tests, false)
}
