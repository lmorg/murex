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
		{
			Expression: `bob=5`,
			Expected:   nil,
		},
		{
			Expression: `bob =5`,
			Expected:   nil,
		},
		{
			Expression: `bob= 5`,
			Expected:   nil,
		},
		{
			Expression: `bob = 5`,
			Expected:   nil,
		},
		{
			Expression: `bob = (5==5)`,
			Expected:   nil,
		},
		{
			Expression: `bob=(5==5)`,
			Expected:   nil,
		},
		{
			Expression: `bob=("5"=="5")`,
			Expected:   nil,
		},
		{
			Expression: `bob = 5==5`,
			Expected:   nil,
		},
		{
			Expression: `bob=5==5`,
			Expected:   nil,
		},
		{
			Expression: `bob="5"=="5"`,
			Expected:   nil,
		},
		{
			Expression: `bob=(5*5)`,
			Expected:   nil,
		},
	}

	testExpression(t, tests)
}

func TestStupidOffByOneErrorsInSubExpressions(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `bob = 1+1`,
			Expected:   nil,
		},
		{
			Expression: `10*(1+2)*10`,
			Expected:   float64(300),
		},
		{
			Expression: `2*(10*(1+2)*10)`,
			Expected:   float64(600),
		},
		{
			Expression: `(10*(1+2)*10)*2`,
			Expected:   float64(600),
		},
		{
			Expression: `2*((10*(1+2)*10)+2)`,
			Expected:   float64(604),
		},
		{
			Expression: `2*((10*(1+2)*10)+2)*2`,
			Expected:   float64(1208),
		},
		{
			Expression: `(2*((10*(1+2)*10)+2)*2)`,
			Expected:   float64(1208),
		},
	}

	testExpression(t, tests)
}
