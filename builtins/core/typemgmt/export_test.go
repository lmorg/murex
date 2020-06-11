package typemgmt

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/core/io"
)

func TestExportFunctionPositive(t *testing.T) {
	set := []Test{
		{
			Block:    "export: f=b",
			Name:     "f",
			Value:    "b",
			DataType: "str",
		},
		{
			Block:    "export: foo=b",
			Name:     "foo",
			Value:    "b",
			DataType: "str",
		},
		{
			Block:    "export: f=bar",
			Name:     "f",
			Value:    "bar",
			DataType: "str",
		},
		{
			Block:    "export: foo=bar",
			Name:     "foo",
			Value:    "bar",
			DataType: "str",
		},
		{
			Block:    "export: _=foobar",
			Name:     "_",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "export: _b=foobar",
			Name:     "_b",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "export: f_=foobar",
			Name:     "f_",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "export: f_b=foobar",
			Name:     "f_b",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "export: foo_b=foobar",
			Name:     "foo_b",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "export: f_bar=foobar",
			Name:     "f_bar",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "export: foo_bar=foobar",
			Name:     "foo_bar",
			Value:    "foobar",
			DataType: "str",
		},
	}

	VariableTests(set, t)

	unset := []string{
		"f",
		"foo",
		"_",
		"_b",
		"f_",
		"f_b",
		"foo_b",
		"f_bar",
		"foo_bar",
		"foobar",
	}

	UnSetTests("!export", unset, t)
}

func TestExportMethodPositive(t *testing.T) {
	set := []Test{
		{
			Block:    "out: b -> export: f",
			Name:     "f",
			Value:    "b",
			DataType: "str",
		},
		{
			Block:    "out: b -> export: foo",
			Name:     "foo",
			Value:    "b",
			DataType: "str",
		},
		{
			Block:    "out: bar -> export: f",
			Name:     "f",
			Value:    "bar",
			DataType: "str",
		},
		{
			Block:    "out: bar -> export: foo",
			Name:     "foo",
			Value:    "bar",
			DataType: "str",
		},
		{
			Block:    "out: foobar -> export: _",
			Name:     "_",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "out: foobar -> export: _b",
			Name:     "_b",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "out: foobar -> export: f_",
			Name:     "f_",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "out: foobar -> export: f_b",
			Name:     "f_b",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "out: foobar -> export: foo_b",
			Name:     "foo_b",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "out: foobar -> export: f_bar",
			Name:     "f_bar",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "out: foobar -> export: foo_bar",
			Name:     "foo_bar",
			Value:    "foobar",
			DataType: "str",
		},
	}

	VariableTests(set, t)

	unset := []string{
		"f",
		"foo",
		"_",
		"_b",
		"f_",
		"f_b",
		"foo_b",
		"f_bar",
		"foo_bar",
		"foobar",
	}

	UnSetTests("!export", unset, t)
}

func TestExportFunctionNegative(t *testing.T) {
	tests := []Test{
		{
			Block: "export: =foobar",
			Fail:  true,
		},
		{
			Block: "export: -=foobar",
			Fail:  true,
		},
		{
			Block: "export: foo-bar=foobar",
			Fail:  true,
		},
		{
			Block: "export: foo\\-bar=foobar",
			Fail:  true,
		},
		{
			Block: "export: foobar =foobar",
			Fail:  true,
		},
		{
			Block: "export: foobar = foobar",
			Fail:  true,
		},
	}

	VariableTests(tests, t)
}

func TestExportMethodNegative(t *testing.T) {
	tests := []Test{
		{
			Block: "out: foobar -> export",
			Fail:  true,
		},
		{
			Block: "out: foobar -> export: =",
			Fail:  true,
		},
		{
			Block: "out: foobar -> export: -",
			Fail:  true,
		},
		{
			Block: "out: foobar -> export: foo-bar",
			Fail:  true,
		},
		{
			Block: "out: foobar -> export: foo\\-bar",
			Fail:  true,
		},
		{
			Block: "out: foobar -> export: foo=",
			Fail:  true,
		},
		{
			Block: "out: foobar -> export: foo=bar",
			Fail:  true,
		},
	}

	VariableTests(tests, t)
}
