package typemgmt

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/core/io"
)

func TestLetFunctionPositive(t *testing.T) {
	tests := []Test{
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
			Block:    "let: _=`foobar`",
			Name:     "_",
			Value:    "foobar",
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
	}

	VariableTests(tests, t)
}

func TestLetFunctionNegative(t *testing.T) {
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
	tests := []Test{
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
		}, {
			Block:    "let: foobar=false",
			Name:     "foobar",
			Value:    "false",
			DataType: "bool",
		},
	}

	VariableTests(tests, t)
}

func TestLetFunctionEvaluation(t *testing.T) {
	tests := []Test{
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

	VariableTests(tests, t)
}
