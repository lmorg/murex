package expressions_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

func TestExpressionsBuiltin(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `TestExpressionsBuiltin0=1+1+1+1+1+1+1+1+1+1`,
			Stdout: ``,
		},
		{
			Block:  `null;TestExpressionsBuiltin0=1+1+1+1+1+1+1+1+1+1`,
			Stdout: ``,
		},
		{
			Block:  `null;TestExpressionsBuiltin1=(1+1+1+1+1+1+1+1+1+1)`,
			Stdout: ``,
		},
		{
			Block:  `TestExpressionsBuiltin2=1+1+1+1+1+1+1+1+1+1;null`,
			Stdout: ``,
		},
		{
			Block:  `TestExpressionsBuiltin3=(1+1+1+1+1+1+1+1+1+1);null`,
			Stdout: ``,
		},
		{
			Block:  `null;TestExpressionsBuiltin4=1+1+1+1+1+1+1+1+1+1;null`,
			Stdout: ``,
		},
		{
			Block:  `null;TestExpressionsBuiltin5=(1+1+1+1+1+1+1+1+1+1);null`,
			Stdout: ``,
		},
		{
			Block:  `null;TestExpressionsBuiltin6=1+1+1+1+(1+1)+1+1+1+1;null`,
			Stdout: ``,
		},
		{
			Block:  `null;TestExpressionsBuiltin7=(1+1+1+1+(1+1)+1+1+1+1);null`,
			Stdout: ``,
		},
		/////
		{
			Block:  `null;TestExpressionsBuiltin0=1+1+1+1+1+1+1+1+1+1;out $TestExpressionsBuiltin0`,
			Stdout: "10\n",
		},
		{
			Block:  `null;TestExpressionsBuiltin1=(1+1+1+1+1+1+1+1+1+1);out $TestExpressionsBuiltin1`,
			Stdout: "10\n",
		},
		{
			Block:  `TestExpressionsBuiltin2=1+1+1+1+1+1+1+1+1+1;null;out $TestExpressionsBuiltin2`,
			Stdout: "10\n",
		},
		{
			Block:  `TestExpressionsBuiltin3=(1+1+1+1+1+1+1+1+1+1);null;out $TestExpressionsBuiltin3`,
			Stdout: "10\n",
		},
		{
			Block:  `null;TestExpressionsBuiltin4=1+1+1+1+1+1+1+1+1+1;null;out $TestExpressionsBuiltin4`,
			Stdout: "10\n",
		},
		{
			Block:  `null;TestExpressionsBuiltin5=(1+1+1+1+1+1+1+1+1+1);null;out $TestExpressionsBuiltin5`,
			Stdout: "10\n",
		},
		{
			Block:  `null;TestExpressionsBuiltin6=1+1+1+1+(1+1)+1+1+1+1;null;out $TestExpressionsBuiltin6`,
			Stdout: "10\n",
		},
		{
			Block:  `null;TestExpressionsBuiltin7=(1+1+1+1+(1+1)+1+1+1+1);null;out $TestExpressionsBuiltin7`,
			Stdout: "10\n",
		},
		/////
		{
			Block:  `null;1+1+1+1+1+1+1+1+1+1`,
			Stdout: `10`,
		},
		{
			Block:  `null;(1+1+1+1+1+1+1+1+1+1)`,
			Stdout: `1+1+1+1+1+1+1+1+1+1`,
		},
		{
			Block:  `1+1+1+1+1+1+1+1+1+1;null`,
			Stdout: `10`,
		},
		{
			Block:  `(1+1+1+1+1+1+1+1+1+1);null`,
			Stdout: `1+1+1+1+1+1+1+1+1+1`,
		},
		{
			Block:  `null;1+1+1+1+1+1+1+1+1+1;null`,
			Stdout: `10`,
		},
		{
			Block:  `null;(1+1+1+1+1+1+1+1+1+1);null`,
			Stdout: `1+1+1+1+1+1+1+1+1+1`,
		},
		{
			Block:  `null;1+1+1+1+(1+1)+1+1+1+1;null`,
			Stdout: `10`,
		},
		{
			Block:  `null;(1+1+1+1+(1+1)+1+1+1+1);null`,
			Stdout: `1+1+1+1+(1+1)+1+1+1+1`,
		},
		/////
		{
			Block:  `1+1+1+1+1+1+1+1+1+1`,
			Stdout: `10`,
		},
		{
			Block:  `null;1+1+1+1+1+1+1+1+1+1`,
			Stdout: `10`,
		},
		{
			Block:  `(1+1+1+1+1+1+1+1+1+1)`,
			Stdout: `1+1+1+1+1+1+1+1+1+1`,
		},
		{
			Block:  `1+1+1+1+1+1+1+1+1+1;null`,
			Stdout: `10`,
		},
		{
			Block:  `(1+1+1+1+1+1+1+1+1+1);null`,
			Stdout: `1+1+1+1+1+1+1+1+1+1`,
		},
		{
			Block:  `null;1+1+1+1+1+1+1+1+1+1;null`,
			Stdout: `10`,
		},
		{
			Block:  `null;(1+1+1+1+1+1+1+1+1+1);null`,
			Stdout: `1+1+1+1+1+1+1+1+1+1`,
		},
		{
			Block:  `null;1+1+1+1+(1+1)+1+1+1+1;null`,
			Stdout: `10`,
		},
		{
			Block:  `null;(1+1+1+1+(1+1)+1+1+1+1);null`,
			Stdout: `1+1+1+1+(1+1)+1+1+1+1`,
		},
		/////
		{
			Block:  `3*(3+1)`,
			Stdout: `12`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestExpressionsScalars(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `TestExpressionsScalars0="foobar";$TestExpressionsScalars0=="foobar"`,
			Stdout: `true`,
		},
		{
			Block:  `TestExpressionsScalars1="foobar";%[1,2,$TestExpressionsScalars1]`,
			Stdout: `[1,2,"foobar"]`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestExpressionsBuiltinSubExpr(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `(1+1);null`,
			Stdout: `1+1`,
		},
		{
			Block:  `(1+1) ;null`,
			Stdout: `1+1`,
		},
		{
			Block:  `(1+1) ; null`,
			Stdout: `1+1`,
		},
		/////
		{
			Block:  `bob=(1+1);null`,
			Stdout: ``,
		},
		{
			Block:  `bob=(1+1) ;null`,
			Stdout: ``,
		},
		{
			Block:  `bob=(1+1) ; null`,
			Stdout: ``,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestExpressionsBuiltinStrings(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `%(bob)`,
			Stdout: `bob`,
		},
	}

	test.RunMurexTests(tests, t)
}

// https://github.com/lmorg/murex/issues/827
func TestExpressionsMultipleParams(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `expr 1 + 2`,
			Stdout: `3`,
		},
	}

	test.RunMurexTests(tests, t)
}
