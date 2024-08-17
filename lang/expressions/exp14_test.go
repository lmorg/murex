package expressions

import (
	"testing"

	"github.com/lmorg/murex/test"
	"github.com/lmorg/murex/test/count"
)

// https://github.com/lmorg/murex/issues/793
func TestScalarNameDetokenised(t *testing.T) {
	// this just tests for panics

	test := []rune("$(test)")

	count.Tests(t, len(test))

	for i := range test {
		scalarNameDetokenised(test[:i])
	}
}

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

	testExpression(t, tests, true)
}

func TestExpAssignAdd(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `TestExpAssignAdd0 += 5`,
			Expected:   nil,
		},
		{
			Expression: `TestExpAssignAdd1 += "bar"`,
			Error:      true,
		},
		{
			Expression: `TestExpAssignAdd2 += bar`,
			Error:      true,
		},
		{
			Expression: `TestExpAssignAdd3 += true`,
			Error:      true,
		},
		{
			Expression: `TestExpAssignAdd4 += >`,
			Error:      true,
		},
	}

	testExpression(t, tests, true)
}

func TestExpAssignSubtract(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `TestExpAssignSubtract0 -= 5`,
			Expected:   nil,
		},
		{
			Expression: `TestExpAssignSubtract1 -= "bar"`,
			Error:      true,
		},
		{
			Expression: `TestExpAssignSubtract2 -= bar`,
			Error:      true,
		},
		{
			Expression: `TestExpAssignSubtract3 -= true`,
			Error:      true,
		},
		{
			Expression: `TestExpAssignSubtract4 -= >`,
			Error:      true,
		},
	}

	testExpression(t, tests, true)
}

func TestExpAssignMultiply(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `TestExpAssignMultiply0 *= 5`,
			Expected:   nil,
		},
		{
			Expression: `TestExpAssignMultiply1 *= "bar"`,
			Error:      true,
		},
		{
			Expression: `TestExpAssignMultiply2 *= bar`,
			Error:      true,
		},
		{
			Expression: `TestExpAssignMultiply3 *= true`,
			Error:      true,
		},
		{
			Expression: `TestExpAssignMultiply4 *= >`,
			Error:      true,
		},
	}

	testExpression(t, tests, true)
}

func TestExpAssignDivide(t *testing.T) {
	tests := []expressionTestT{
		{
			Expression: `TestExpAssignDivide0 /= 5`,
			Expected:   nil,
		},
		{
			Expression: `TestExpAssignDivide1 /= "bar"`,
			Error:      true,
		},
		{
			Expression: `TestExpAssignDivide2 /= bar`,
			Error:      true,
		},
		{
			Expression: `TestExpAssignDivide3 /= true`,
			Error:      true,
		},
		{
			Expression: `TestExpAssignDivide4 /= >`,
			Error:      true,
		},
	}

	testExpression(t, tests, true)
}

func TestLazyAssigns(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  "TestLazyAssignsAdd1 += 5; $TestLazyAssignsAdd1 += 5; $TestLazyAssignsAdd1",
			Stdout: "10",
		},
		{
			Block:  "TestLazyAssignsSubtract1 -= 5; $TestLazyAssignsSubtract1 -= 5; $TestLazyAssignsSubtract1",
			Stdout: "-10",
		},
		{
			Block:  "TestLazyAssignsMultiply1 *= 5; $TestLazyAssignsMultiply1 *= 5; $TestLazyAssignsMultiply1",
			Stdout: "0",
		},
		{
			Block:  "TestLazyAssignsDivide1 /= 5; $TestLazyAssignsDivide1 /= 5; $TestLazyAssignsDivide1",
			Stdout: "0",
		},
	}

	test.RunMurexTests(tests, t)
}
