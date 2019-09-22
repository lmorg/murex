package types_test

import (
	"testing"

	"github.com/lmorg/murex/lang/types"
)

func TestConvertGoTypeBoolean(t *testing.T) {
	tests := []test{
		{
			In:       nil,
			DataType: types.Boolean,
			Out:      false,
		},
		{
			In:       "",
			DataType: types.Boolean,
			Out:      false,
		},
		{
			In:       "foobar",
			DataType: types.Boolean,
			Out:      true,
		},
		{
			In:       "0",
			DataType: types.Boolean,
			Out:      false,
		},
		{
			In:       "true",
			DataType: types.Boolean,
			Out:      true,
		},
		{
			In:       "false",
			DataType: types.Boolean,
			Out:      false,
		},
		{
			In:       "42",
			DataType: types.Boolean,
			Out:      true,
		},
		{
			In:       "42.12345",
			DataType: types.Boolean,
			Out:      true,
		},
		{
			In:       0,
			DataType: types.Boolean,
			Out:      false,
		},
		{
			In:       float64(0),
			DataType: types.Boolean,
			Out:      false,
		},
		{
			In:       42,
			DataType: types.Boolean,
			Out:      true,
		},
		{
			In:       42.12345,
			DataType: types.Boolean,
			Out:      true,
		},
		{
			In:       true,
			DataType: types.Boolean,
			Out:      true,
		},
		{
			In:       false,
			DataType: types.Boolean,
			Out:      false,
		},
		{
			In:       `{ out: "testing" }`,
			DataType: types.Boolean,
			Out:      true,
		},
		{
			In:       `{ "foo": "bar" }`,
			DataType: types.Boolean,
			Out:      true,
		},
	}

	testConvertGoType(t, tests)
}
