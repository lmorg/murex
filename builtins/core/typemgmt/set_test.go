package typemgmt

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/core/io"
)

func TestSetFunctionPositive(t *testing.T) {
	tests := []Test{
		{
			Block:    "set: f=b",
			Name:     "f",
			Value:    "b",
			DataType: "str",
		},
		{
			Block:    "set: foo=b",
			Name:     "foo",
			Value:    "b",
			DataType: "str",
		},
		{
			Block:    "set: f=bar",
			Name:     "f",
			Value:    "bar",
			DataType: "str",
		},
		{
			Block:    "set: foo=bar",
			Name:     "foo",
			Value:    "bar",
			DataType: "str",
		},
		{
			Block:    "set: _=foobar",
			Name:     "_",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "set: _b=foobar",
			Name:     "_b",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "set: f_=foobar",
			Name:     "f_",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "set: f_b=foobar",
			Name:     "f_b",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "set: foo_b=foobar",
			Name:     "foo_b",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "set: f_bar=foobar",
			Name:     "f_bar",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "set: foo_bar=foobar",
			Name:     "foo_bar",
			Value:    "foobar",
			DataType: "str",
		},
	}

	VariableTests(tests, t)
}

func TestSetMethodPositive(t *testing.T) {
	tests := []Test{
		{
			Block:    "out: b -> set: f",
			Name:     "f",
			Value:    "b",
			DataType: "str",
		},
		{
			Block:    "out: b -> set: foo",
			Name:     "foo",
			Value:    "b",
			DataType: "str",
		},
		{
			Block:    "out: bar -> set: f",
			Name:     "f",
			Value:    "bar",
			DataType: "str",
		},
		{
			Block:    "out: bar -> set: foo",
			Name:     "foo",
			Value:    "bar",
			DataType: "str",
		},
		{
			Block:    "out: foobar -> set: _",
			Name:     "_",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "out: foobar -> set: _b",
			Name:     "_b",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "out: foobar -> set: f_",
			Name:     "f_",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "out: foobar -> set: f_b",
			Name:     "f_b",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "out: foobar -> set: foo_b",
			Name:     "foo_b",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "out: foobar -> set: f_bar",
			Name:     "f_bar",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "out: foobar -> set: foo_bar",
			Name:     "foo_bar",
			Value:    "foobar",
			DataType: "str",
		},
	}

	VariableTests(tests, t)
}

func TestSetFunctionNegative(t *testing.T) {
	tests := []Test{
		{
			Block: "set: =foobar",
			Fail:  true,
		},
		{
			Block: "set: -=foobar",
			Fail:  true,
		},
		{
			Block: "set: foo-bar=foobar",
			Fail:  true,
		},
		{
			Block: "set: foo\\-bar=foobar",
			Fail:  true,
		},
	}

	VariableTests(tests, t)
}

func TestSetMethodNegative(t *testing.T) {
	tests := []Test{
		{
			Block: "out: foobar -> set",
			Fail:  true,
		},
		{
			Block: "out: foobar -> set: =",
			Fail:  true,
		},
		{
			Block: "out: foobar -> set: -",
			Fail:  true,
		},
		{
			Block: "out: foobar -> set: foo-bar",
			Fail:  true,
		},
		{
			Block: "out: foobar -> set: foo\\-bar",
			Fail:  true,
		},
		{
			Block: "out: foobar -> set: foo=",
			Fail:  true,
		},
		{
			Block: "out: foobar -> set: foo=bar",
			Fail:  true,
		},
	}

	VariableTests(tests, t)
}
