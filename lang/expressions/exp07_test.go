package expressions

import "testing"

func TestExpEqualToStrict(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `"foobar" == "foobar"`,
			Expected:   true,
		},
		{
			Expression: `"foo" == "bar"`,
			Expected:   false,
		},
		///
		{
			Expression: `1 == 1`,
			Expected:   true,
		},
		{
			Expression: `1 == 2`,
			Expected:   false,
		},
		///
		{
			Expression: `1 == "1"`,
			Error:      true,
		},
		{
			Expression: `1 == "2"`,
			Error:      true,
		},
		///
		{
			Expression: `$variable == ""`,
			Error:      true,
		},
		{
			Expression: `$variable == null`,
			Expected:   true,
		},
	}

	testExpression(t, tests, true)
}

func TestExpEqualToWeak(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `"foobar" == "foobar"`,
			Expected:   true,
		},
		{
			Expression: `"foo" == "bar"`,
			Expected:   false,
		},
		///
		{
			Expression: `1 == 1`,
			Expected:   true,
		},
		{
			Expression: `1 == 2`,
			Expected:   false,
		},
		///
		{
			Expression: `1 == "1"`,
			Expected:   true,
		},
		{
			Expression: `1 == "2"`,
			Expected:   false,
		},
		///
		{
			Expression: `$variable == ""`,
			Expected:   true,
		},
		{
			Expression: `$variable == null`,
			Expected:   true,
		},
	}

	testExpression(t, tests, false)
}

func TestExpNotEqualToStrong(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `"foobar" != "foobar"`,
			Expected:   false,
		},
		{
			Expression: `"foo" != "bar"`,
			Expected:   true,
		},
		///
		{
			Expression: `1 != 1`,
			Expected:   false,
		},
		{
			Expression: `1 != 2`,
			Expected:   true,
		},
		///
		{
			Expression: `1 != "1"`,
			Error:   true,
		},
		{
			Expression: `1 != "2"`,
			Error:   true,
		},
	}

	testExpression(t, tests, true)
}

func TestExpNotEqualToWeak(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `"foobar" != "foobar"`,
			Expected:   false,
		},
		{
			Expression: `"foo" != "bar"`,
			Expected:   true,
		},
		///
		{
			Expression: `1 != 1`,
			Expected:   false,
		},
		{
			Expression: `1 != 2`,
			Expected:   true,
		},
		///
		{
			Expression: `1 != "1"`,
			Expected:   false,
		},
		{
			Expression: `1 != "2"`,
			Expected:   true,
		},
	}

	testExpression(t, tests, false)
}

func TestExpLike(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `"foobar" ~~ "foobar"`,
			Expected:   true,
		},
		{
			Expression: `"foobar" ~~ "FOOBAR"`,
			Expected:   true,
		},
		{
			Expression: `"foo" ~~ "bar"`,
			Expected:   false,
		},
		///
		{
			Expression: `1 ~~ 1`,
			Expected:   true,
		},
		{
			Expression: `1 ~~ 2`,
			Expected:   false,
		},
		///
		{
			Expression: `1 ~~ "1"`,
			Expected:   true,
		},
		{
			Expression: `1 ~~ "2"`,
			Expected:   false,
		},
		///
		{
			Expression: `(1==1) ~~ "true"`,
			Expected:   true,
		},
		{
			Expression: `(1==1) ~~ "false"`,
			Expected:   false,
		},
	}

	testExpression(t, tests, true)
}

func TestExpNotLike(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `"foobar" !! "foobar"`,
			Expected:   false,
		},
		{
			Expression: `"foobar" !! "FOOBAR"`,
			Expected:   false,
		},
		{
			Expression: `"foo" !! "bar"`,
			Expected:   true,
		},
		///
		{
			Expression: `1 !! 1`,
			Expected:   false,
		},
		{
			Expression: `1 !! 2`,
			Expected:   true,
		},
		///
		{
			Expression: `1 !! "1"`,
			Expected:   false,
		},
		{
			Expression: `1 !! "2"`,
			Expected:   true,
		},
		///
		{
			Expression: `(1==1) !! "true"`,
			Expected:   false,
		},
		{
			Expression: `(1==1) !! "false"`,
			Expected:   true,
		},
	}

	testExpression(t, tests, true)
}

func TestExpRegexp(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `"foobar" =~ "foobar"`,
			Expected:   true,
		},
		{
			Expression: `"foo" =~ "bar"`,
			Expected:   false,
		},
		///
		{
			Expression: `1 =~ 1`,
			Error:      true,
		},
		{
			Expression: `1 =~ 2`,
			Error:      true,
		},
		///
		{
			Expression: `1 =~ "1"`,
			Error:      true,
		},
		{
			Expression: `1 =~ "2"`,
			Error:      true,
		},
		///
		{
			Expression: `"foobar" =~ "foo"`,
			Expected:   true,
		},
		{
			Expression: `"foobar" =~ "foo$"`,
			Expected:   false,
		},
	}

	testExpression(t, tests, true)
}

func TestExpNotRegexp(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `"foobar" !~ "foobar"`,
			Expected:   false,
		},
		{
			Expression: `"foo" !~ "bar"`,
			Expected:   true,
		},
		///
		{
			Expression: `1 !~ 1`,
			Error:      true,
		},
		{
			Expression: `1 !~ 2`,
			Error:      true,
		},
		///
		{
			Expression: `1 !~ "1"`,
			Error:      true,
		},
		{
			Expression: `1 !~ "2"`,
			Error:      true,
		},
		///
		{
			Expression: `"foobar" !~ "foo"`,
			Expected:   false,
		},
		{
			Expression: `"foobar" !~ "foo$"`,
			Expected:   true,
		},
	}

	testExpression(t, tests, true)
}
