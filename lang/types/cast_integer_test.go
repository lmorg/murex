package types_test

import (
	"testing"

	"github.com/lmorg/murex/lang/types"
)

func TestConvertGoTypeInteger(t *testing.T) {
	tests := []test{
		{
			In:       nil,
			DataType: types.Integer,
			Out:      0,
		},
		{
			In:       "",
			DataType: types.Integer,
			Out:      0,
		},
		{
			In:       "foobar",
			DataType: types.Integer,
			Out:      0,
			Error:    true,
		},
		{
			In:       "0",
			DataType: types.Integer,
			Out:      0,
		},
		{
			In:       "true",
			DataType: types.Integer,
			Out:      0,
			Error:    true,
		},
		{
			In:       "false",
			DataType: types.Integer,
			Out:      0,
			Error:    true,
		},
		{
			In:       "42",
			DataType: types.Integer,
			Out:      42,
		},
		{
			In:       "42.12345",
			DataType: types.Integer,
			Out:      42,
		},
		{
			In:       0,
			DataType: types.Integer,
			Out:      0,
		},
		{
			In:       float64(0),
			DataType: types.Integer,
			Out:      0,
		},
		{
			In:       42,
			DataType: types.Integer,
			Out:      42,
		},
		{
			In:       42.12345,
			DataType: types.Integer,
			Out:      42,
		},
		{
			In:       true,
			DataType: types.Integer,
			Out:      1,
		},
		{
			In:       false,
			DataType: types.Integer,
			Out:      0,
		},
		{
			In:       `{ out: "testing" }`,
			DataType: types.Integer,
			Out:      0,
			Error:    true,
		},
		{
			In:       `{ "foo": "bar" }`,
			DataType: types.Integer,
			Out:      0,
			Error:    true,
		},
	}

	testConvertGoType(t, tests)
}
