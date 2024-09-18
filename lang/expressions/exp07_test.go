package expressions

import (
	"testing"

	"github.com/lmorg/murex/test"
)

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
			Expression: `1.0 == "1"`,
			Expected:   true,
		},
		{
			Expression: `1.0 == "2"`,
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
			Error:      true,
		},
		{
			Expression: `1 != "2"`,
			Error:      true,
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
		///
		{
			Expression: `1.0 != "1"`,
			Expected:   false,
		},
		{
			Expression: `1.0 != "2"`,
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
			Expression: `"foobar" =~ 'foo$'`,
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
			Expression: `"foobar" !~ 'foo$'`,
			Expected:   true,
		},
	}

	testExpression(t, tests, true)
}

// https://github.com/lmorg/murex/issues/831
func TestExprEquBugFixes(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:   `%[z y x] == %[z y x]`,
			Stdout:  `true`,
			ExitNum: 0,
		},
		{
			Block:   `"%[z y x]" == %[z y x]`,
			Stdout:  `false`,
			ExitNum: 1,
		},
		{
			Block:   `%[z y x] == "%[z y x]"`,
			Stdout:  `false`,
			ExitNum: 1,
		},
		{
			Block:   `"%[z y x]" == "%[z y x]"`,
			Stdout:  `true`,
			ExitNum: 0,
		},
		{
			Block:   `%[z y x] == '["z","y","x"]'`,
			Stdout:  `true`,
			ExitNum: 0,
		},
		/////
		{
			Block:   `%{z:3, y:2, x:1} == %{z:3, y:2, x:1}`,
			Stdout:  `true`,
			ExitNum: 0,
		},
		{
			Block:   `"%{z:3, y:2, x:1}" == %{z:3, y:2, x:1}`,
			Stdout:  `false`,
			ExitNum: 1,
		},
		{
			Block:   `%{z:3, y:2, x:1} == "%{z:3, y:2, x:1}"`,
			Stdout:  `false`,
			ExitNum: 1,
		},
		{
			Block:   `"%{z:3, y:2, x:1}" == "%{z:3, y:2, x:1}"`,
			Stdout:  `true`,
			ExitNum: 0,
		},
		{
			Block:   `%{z:3, y:2, x:1} == '{"x":1,"y":2,"z":3}'`,
			Stdout:  `true`,
			ExitNum: 0,
		},
	}

	test.RunMurexTestsRx(tests, t)
}
