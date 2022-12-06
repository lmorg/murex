package expressions_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
	"github.com/lmorg/murex/test/count"
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

	count.Tests(t, len(tests))

	test.RunMurexTests(tests, t)
}

func TestExpressionsScalars(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `TestExpressionsScalars0="foobar";$TestExpressionsScalars0=="foobar"`,
			Stdout: `true`,
		},
	}

	count.Tests(t, len(tests))

	test.RunMurexTests(tests, t)
}
