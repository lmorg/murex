package types_test

import (
	"testing"

	"github.com/lmorg/murex/lang/types"
)

func TestConvertGoTypeNull(t *testing.T) {
	tests := []test{
		{
			In:       nil,
			DataType: types.Null,
			Out:      "",
		},
		{
			In:       "",
			DataType: types.Null,
			Out:      "",
		},
		{
			In:       "foobar",
			DataType: types.Null,
			Out:      "",
		},
		{
			In:       "0",
			DataType: types.Null,
			Out:      "",
		},
		{
			In:       "true",
			DataType: types.Null,
			Out:      "",
		},
		{
			In:       "false",
			DataType: types.Null,
			Out:      "",
		},
		{
			In:       0,
			DataType: types.Null,
			Out:      "",
		},
		{
			In:       float64(0),
			DataType: types.Null,
			Out:      "",
		},
		{
			In:       42,
			DataType: types.Null,
			Out:      "",
		},
		{
			In:       42.12345,
			DataType: types.Null,
			Out:      "",
		},
		{
			In:       true,
			DataType: types.Null,
			Out:      "",
		},
		{
			In:       false,
			DataType: types.Null,
			Out:      "",
		},
		{
			In:       `{ out: "testing" }`,
			DataType: types.Null,
			Out:      ``,
		},
		{
			In:       `{ "foo": "bar" }`,
			DataType: types.Null,
			Out:      ``,
		},
	}

	testConvertGoType(t, tests)
}
