package typemgmt

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/core/io"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test/count"
)

func TestLetFunctionPositive(t *testing.T) {
	lang.InitEnv()

	set := []Test{
		{
			Block:    "let: f=`b`",
			Name:     "f",
			Value:    "b",
			DataType: "*",
		},
		{
			Block:    "let: foo=`b`",
			Name:     "foo",
			Value:    "b",
			DataType: "*",
		},
		{
			Block:    "let: f=`bar`",
			Name:     "f",
			Value:    "bar",
			DataType: "*",
		},
		{
			Block:    "let: foo=`bar`",
			Name:     "foo",
			Value:    "bar",
			DataType: "*",
		},
		{
			Block:    "let: _b=`foobar`",
			Name:     "_b",
			Value:    "foobar",
			DataType: "*",
		},
		{
			Block:    "let: f_=`foobar`",
			Name:     "f_",
			Value:    "foobar",
			DataType: "*",
		},
		{
			Block:    "let: f_b=`foobar`",
			Name:     "f_b",
			Value:    "foobar",
			DataType: "*",
		},
		{
			Block:    "let: foo_b=`foobar`",
			Name:     "foo_b",
			Value:    "foobar",
			DataType: "*",
		},
		{
			Block:    "let: f_bar=`foobar`",
			Name:     "f_bar",
			Value:    "foobar",
			DataType: "*",
		},
		{
			Block:    "let: foo_bar=`foobar`",
			Name:     "foo_bar",
			Value:    "foobar",
			DataType: "*",
		},
		{
			Block:    "let: foobar=`foobar`",
			Name:     "foobar",
			Value:    "foobar",
			DataType: "*",
		},
	}

	unset := []string{
		"f",
		"foo",
		"_b",
		"f_",
		"f_b",
		"foo_b",
		"f_bar",
		"foo_bar",
		"foobar",
	}

	VariableTests(set, t)
	UnSetTests("!set", unset, t)
}

func TestLetFunctionNegative(t *testing.T) {
	lang.InitEnv()

	tests := []Test{
		{
			Block: "let: =foobar",
			Fail:  true,
		},
		{
			Block: "let: -=foobar",
			Fail:  true,
		},
		{
			Block: "let: _=foobar",
			Fail:  true,
		},
		{
			Block: "let: foo-bar=foobar",
			Fail:  true,
		},
		{
			Block: "let: foo\\-bar=foobar",
			Fail:  true,
		},
	}

	VariableTests(tests, t)
}

func TestLetFunctionDataTypes(t *testing.T) {
	lang.InitEnv()

	set := []Test{
		{
			Block:    "let: foobar=123",
			Name:     "foobar",
			Value:    "123",
			DataType: "num",
		},
		{
			Block:    "let: foobar=123.456",
			Name:     "foobar",
			Value:    "123.456",
			DataType: "num",
		},
		{
			Block:    "let: foobar=true",
			Name:     "foobar",
			Value:    "true",
			DataType: "bool",
		},
		{
			Block:    "let: foobar=false",
			Name:     "foobar",
			Value:    "false",
			DataType: "bool",
		},
	}

	unset := []string{
		"foobar",
	}

	VariableTests(set, t)
	UnSetTests("!set", unset, t)
}

func TestLetFunctionEvaluation(t *testing.T) {
	lang.InitEnv()

	set := []Test{
		{
			Block:    "let: foobar=2+2",
			Name:     "foobar",
			Value:    "4",
			DataType: "num",
		},
		{
			Block:    "let: foobar=5/2",
			Name:     "foobar",
			Value:    "2.5",
			DataType: "num",
		},
		{
			Block:    "let: foobar=1==1",
			Name:     "foobar",
			Value:    "true",
			DataType: "bool",
		}, {
			Block:    "let: foobar=1!=1",
			Name:     "foobar",
			Value:    "false",
			DataType: "bool",
		},
	}

	unset := []string{
		"foobar",
	}

	VariableTests(set, t)
	UnSetTests("!set", unset, t)
}

func TestLetBuilder(t *testing.T) {
	type TestLetBuilderT struct {
		Params     string
		Variable   string
		Expression string
		Error      bool
	}

	tests := []TestLetBuilderT{
		{
			Params:     "a--",
			Variable:   "a",
			Expression: "a-1",
		},
		{
			Params:     "a++",
			Variable:   "a",
			Expression: "a+1",
		},
		// a -=
		{
			Params:     "a-=1",
			Variable:   "a",
			Expression: "a-1",
		},
		{
			Params:     "a-=2",
			Variable:   "a",
			Expression: "a-2",
		},
		{
			Params:     "a-=-1",
			Variable:   "a",
			Expression: "a--1",
		},
		//
		{
			Params:     "a -=1",
			Variable:   "a",
			Expression: "a-1",
		},
		{
			Params:     "a -=2",
			Variable:   "a",
			Expression: "a-2",
		},
		{
			Params:     "a -=-1",
			Variable:   "a",
			Expression: "a--1",
		},
		//
		{
			Params:     "a-= 1",
			Variable:   "a",
			Expression: "a-1",
		},
		{
			Params:     "a-= 2",
			Variable:   "a",
			Expression: "a-2",
		},
		{
			Params:     "a-= -1",
			Variable:   "a",
			Expression: "a--1",
		},
		//
		{
			Params:     "a -= 1",
			Variable:   "a",
			Expression: "a-1",
		},
		{
			Params:     "a -= 2",
			Variable:   "a",
			Expression: "a-2",
		},
		{
			Params:     "a -= -1",
			Variable:   "a",
			Expression: "a--1",
		},
		// a +=
		{
			Params:     "a+=1",
			Variable:   "a",
			Expression: "a+1",
		},
		{
			Params:     "a+=2",
			Variable:   "a",
			Expression: "a+2",
		},
		{
			Params:     "a+=-1",
			Variable:   "a",
			Expression: "a+-1",
		},
		//
		{
			Params:     "a +=1",
			Variable:   "a",
			Expression: "a+1",
		},
		{
			Params:     "a +=2",
			Variable:   "a",
			Expression: "a+2",
		},
		{
			Params:     "a +=-1",
			Variable:   "a",
			Expression: "a+-1",
		},
		//
		{
			Params:     "a+= 1",
			Variable:   "a",
			Expression: "a+1",
		},
		{
			Params:     "a+= 2",
			Variable:   "a",
			Expression: "a+2",
		},
		{
			Params:     "a+= -1",
			Variable:   "a",
			Expression: "a+-1",
		},
		//
		{
			Params:     "a += 1",
			Variable:   "a",
			Expression: "a+1",
		},
		{
			Params:     "a += 2",
			Variable:   "a",
			Expression: "a+2",
		},
		{
			Params:     "a += -1",
			Variable:   "a",
			Expression: "a+-1",
		},
		// a /=
		{
			Params:     "a/=1",
			Variable:   "a",
			Expression: "a/1",
		},
		{
			Params:     "a/=2",
			Variable:   "a",
			Expression: "a/2",
		},
		{
			Params:     "a/=-1",
			Variable:   "a",
			Expression: "a/-1",
		},
		//
		{
			Params:     "a /=1",
			Variable:   "a",
			Expression: "a/1",
		},
		{
			Params:     "a /=2",
			Variable:   "a",
			Expression: "a/2",
		},
		{
			Params:     "a /=-1",
			Variable:   "a",
			Expression: "a/-1",
		},
		//
		{
			Params:     "a/= 1",
			Variable:   "a",
			Expression: "a/1",
		},
		{
			Params:     "a/= 2",
			Variable:   "a",
			Expression: "a/2",
		},
		{
			Params:     "a/= -1",
			Variable:   "a",
			Expression: "a/-1",
		},
		//
		{
			Params:     "a /= 1",
			Variable:   "a",
			Expression: "a/1",
		},
		{
			Params:     "a /= 2",
			Variable:   "a",
			Expression: "a/2",
		},
		{
			Params:     "a /= -1",
			Variable:   "a",
			Expression: "a/-1",
		},
		// a *=
		{
			Params:     "a*=1",
			Variable:   "a",
			Expression: "a*1",
		},
		{
			Params:     "a*=2",
			Variable:   "a",
			Expression: "a*2",
		},
		{
			Params:     "a*=-1",
			Variable:   "a",
			Expression: "a*-1",
		},
		//
		{
			Params:     "a *=1",
			Variable:   "a",
			Expression: "a*1",
		},
		{
			Params:     "a *=2",
			Variable:   "a",
			Expression: "a*2",
		},
		{
			Params:     "a *=-1",
			Variable:   "a",
			Expression: "a*-1",
		},
		//
		{
			Params:     "a*= 1",
			Variable:   "a",
			Expression: "a*1",
		},
		{
			Params:     "a*= 2",
			Variable:   "a",
			Expression: "a*2",
		},
		{
			Params:     "a*= -1",
			Variable:   "a",
			Expression: "a*-1",
		},
		//
		{
			Params:     "a *= 1",
			Variable:   "a",
			Expression: "a*1",
		},
		{
			Params:     "a *= 2",
			Variable:   "a",
			Expression: "a*2",
		},
		{
			Params:     "a *= -1",
			Variable:   "a",
			Expression: "a*-1",
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		variable, expression, err := letBuilder(test.Params)

		var errStr string

		if variable != test.Variable {
			errStr += ". Variable mismatch"
		}

		if expression != test.Expression {
			errStr += ". Expression mismatch"
		}

		if (err == nil) == test.Error {
			errStr += ". Error mismatch"
		}

		if errStr != "" {
			t.Errorf("Test %d failed%s", i, errStr)
			t.Logf("  Params:  '%s'", test.Params)
			t.Logf("  exp var: '%s'", test.Variable)
			t.Logf("  act var: '%s'", variable)
			t.Logf("  exp exp: '%s'", test.Expression)
			t.Logf("  act exp: '%s'", expression)
			t.Logf("  exp err:  %v", test.Error)
			t.Logf("  act err:  %s", err)
		}
	}
}
