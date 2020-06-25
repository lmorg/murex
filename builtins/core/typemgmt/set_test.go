package typemgmt

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/core/io"
	"github.com/lmorg/murex/lang"
)

func TestSetFunctionPositive(t *testing.T) {
	lang.InitEnv()

	set := []Test{
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
			Block:    "set: 2=bar",
			Name:     "2",
			Value:    "bar",
			DataType: "str",
		},
		{
			Block:    "set: foo2=bar",
			Name:     "foo2",
			Value:    "bar",
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
		{
			Block:    "set: foobar=foobar",
			Name:     "foobar",
			Value:    "foobar",
			DataType: "str",
		},
		/////
		{
			Block:    "set: gapped = foobar",
			Name:     "gapped",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "set: custard gapped2 = foobar",
			Name:     "gapped2",
			Value:    "foobar",
			DataType: "custard",
		},
		{
			Block: `set: json array = ([
			"a",
			"b",
			"c"
		])`,
			Name: "array",
			Value: `[
			"a",
			"b",
			"c"
		]`,
			DataType: "json",
		},
		{
			Block: `set: json map = {
			"a": "1",
			"b": "2",
			"c": "3"
		}`,
			Name: "map",
			Value: `{
			"a": "1",
			"b": "2",
			"c": "3"
		}`,
			DataType: "json",
		},
		{
			Block: `set: json map2 = ({
			"a": "1",
			"b": "2",
			"c": "3"
		})`,
			Name: "map2",
			Value: `{
			"a": "1",
			"b": "2",
			"c": "3"
		}`,
			DataType: "json",
		},
	}

	unset := []string{
		"f",
		"foo",
		"2",
		"foo2",
		"_b",
		"f_",
		"f_b",
		"foo_b",
		"f_bar",
		"foo_bar",
		"foobar",
		"gapped",
		"gapped2",
		"array",
		"map",
		"map2",
	}

	VariableTests(set, t)
	UnSetTests("!set", unset, t)
}

func TestSetMethodPositive(t *testing.T) {
	lang.InitEnv()

	set := []Test{
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
			Block:    "out: bar -> set: 2",
			Name:     "2",
			Value:    "bar",
			DataType: "str",
		},
		{
			Block:    "out: bar -> set: foo2",
			Name:     "foo2",
			Value:    "bar",
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
		{
			Block:    "out: foobar -> set: foobar",
			Name:     "foobar",
			Value:    "foobar",
			DataType: "str",
		},
	}

	unset := []string{
		"f",
		"foo",
		"2",
		"foo2",
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

func TestSetFunctionNegative(t *testing.T) {
	lang.InitEnv()

	tests := []Test{
		{
			Block: "set: =foobar",
			Fail:  true,
		},
		{
			Block: "set: _=foobar",
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
	lang.InitEnv()

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
			Block: "out: foobar -> set: _",
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
			Block: "out: foobar -> set: foo=bar",
			Fail:  true,
		},
	}

	VariableTests(tests, t)
}

func TestSetFunctionDataTypes(t *testing.T) {
	lang.InitEnv()

	set := []Test{
		{
			Block:    "set: foobar=123",
			Name:     "foobar",
			Value:    "123",
			DataType: "str",
		},
		{
			Block:    "set: foobar=123.456",
			Name:     "foobar",
			Value:    "123.456",
			DataType: "str",
		},
		{
			Block:    "set: foobar=true",
			Name:     "foobar",
			Value:    "true",
			DataType: "str",
		},
		{
			Block:    "set: foobar=false",
			Name:     "foobar",
			Value:    "false",
			DataType: "str",
		},
		{
			Block:    "set: int foobar=123",
			Name:     "foobar",
			Value:    "123",
			DataType: "int",
		},
		{
			Block:    "set: num foobar=123.456",
			Name:     "foobar",
			Value:    "123.456",
			DataType: "num",
		},
		{
			Block:    "set: bool foobar=true",
			Name:     "foobar",
			Value:    "true",
			DataType: "bool",
		},
		{
			Block:    "set: bool foobar=false",
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

func TestSetMethodDataTypes(t *testing.T) {
	lang.InitEnv()

	set := []Test{
		{
			Block:    "tout: int 123 -> set: foobar",
			Name:     "foobar",
			Value:    "123",
			DataType: "int",
		},
		{
			Block:    "tout: num 123.456 -> set: foobar",
			Name:     "foobar",
			Value:    "123.456",
			DataType: "num",
		},
		{
			Block:    "tout: bool true -> set: foobar",
			Name:     "foobar",
			Value:    "true",
			DataType: "bool",
		},
		{
			Block:    "tout: bool false -> set: foobar",
			Name:     "foobar",
			Value:    "false",
			DataType: "bool",
		},
		{
			Block:    "out: 123 -> set: int foobar",
			Name:     "foobar",
			Value:    "123",
			DataType: "int",
		},
		{
			Block:    "out: 123.456 -> set: num foobar",
			Name:     "foobar",
			Value:    "123.456",
			DataType: "num",
		},
		{
			Block:    "tout: int true -> set: bool foobar",
			Name:     "foobar",
			Value:    "true",
			DataType: "bool",
		},
		{
			Block:    "out: false -> set: bool foobar",
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
