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
		///

		{
			In:       []string{"foo", "bar"},
			DataType: types.String,
			Out:      "foo\nbar",
		},
		{
			In:       []int{1, 2},
			DataType: types.String,
			Out:      "1\n2",
		},
		{
			In:       []float64{1.2, 1.3},
			DataType: types.String,
			Out:      "1.2\n1.3",
		},
		{
			In:       []bool{true, false},
			DataType: types.String,
			Out:      "true\nfalse",
		},
		{
			In:       []interface{}{"foo", 1, 2.2, true, "bar"},
			DataType: types.String,
			Out:      "foo\n1\n2.2\ntrue\nbar",
		},
	}

	testConvertGoType(t, tests)
}
