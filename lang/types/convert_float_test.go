package types_test

import (
	"testing"

	"github.com/lmorg/murex/lang/types"
)

func TestConvertGoTypeFloat(t *testing.T) {
	tests := []test{
		{
			In:       nil,
			DataType: types.Float,
			Out:      float64(0),
		},
		{
			In:       "",
			DataType: types.Float,
			Out:      float64(0),
		},
		{
			In:       "foobar",
			DataType: types.Float,
			Out:      float64(0),
			Error:    true,
		},
		{
			In:       "0",
			DataType: types.Float,
			Out:      float64(0),
		},
		{
			In:       "true",
			DataType: types.Float,
			Out:      float64(0),
			Error:    true,
		},
		{
			In:       "false",
			DataType: types.Float,
			Out:      float64(0),
			Error:    true,
		},
		{
			In:       "42",
			DataType: types.Float,
			Out:      float64(42),
		},
		{
			In:       "42.12345",
			DataType: types.Float,
			Out:      float64(42.12345),
		},
		{
			In:       0,
			DataType: types.Float,
			Out:      float64(0),
		},
		{
			In:       float64(0),
			DataType: types.Float,
			Out:      float64(0),
		},
		{
			In:       42,
			DataType: types.Float,
			Out:      float64(42),
		},
		{
			In:       42.12345,
			DataType: types.Float,
			Out:      float64(42.12345),
		},
		{
			In:       true,
			DataType: types.Float,
			Out:      float64(1),
		},
		{
			In:       false,
			DataType: types.Float,
			Out:      float64(0),
		},
		{
			In:       `{ out: "testing" }`,
			DataType: types.Float,
			Out:      float64(0),
			Error:    true,
		},
		{
			In:       `{ "foo": "bar" }`,
			DataType: types.Float,
			Out:      float64(0),
			Error:    true,
		},
	}

	testConvertGoType(t, tests)
}
