package types_test

import (
	"testing"

	"github.com/lmorg/murex/lang/types"
)

func TestConvertGoTypeString(t *testing.T) {
	tests := []test{
		{
			In:       nil,
			DataType: types.String,
			Out:      "",
		},
		{
			In:       "",
			DataType: types.String,
			Out:      "",
		},
		{
			In:       "foobar",
			DataType: types.String,
			Out:      "foobar",
		},
		{
			In:       "0",
			DataType: types.String,
			Out:      "0",
		},
		{
			In:       "true",
			DataType: types.String,
			Out:      "true",
		},
		{
			In:       "false",
			DataType: types.String,
			Out:      "false",
		},
		{
			In:       0,
			DataType: types.String,
			Out:      "0",
		},
		{
			In:       float64(0),
			DataType: types.String,
			Out:      "0",
		},
		{
			In:       42,
			DataType: types.String,
			Out:      "42",
		},
		{
			In:       42.12345,
			DataType: types.String,
			Out:      "42.12345",
		},
		{
			In:       true,
			DataType: types.String,
			Out:      "true",
		},
		{
			In:       false,
			DataType: types.String,
			Out:      "false",
		},
		{
			In:       `{ out: "testing" }`,
			DataType: types.String,
			Out:      `{ out: "testing" }`,
		},
		{
			In:       `{ "foo": "bar" }`,
			DataType: types.String,
			Out:      `{ "foo": "bar" }`,
		},
	}

	testConvertGoType(t, tests)
}
