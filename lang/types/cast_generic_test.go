package types_test

import (
	"testing"

	"github.com/lmorg/murex/lang/types"
)

func TestConvertGoTypeGeneric(t *testing.T) {
	tests := []test{
		{
			In:       nil,
			DataType: types.Generic,
			Out:      "",
		},
		{
			In:       "",
			DataType: types.Generic,
			Out:      "",
		},
		{
			In:       "foobar",
			DataType: types.Generic,
			Out:      "foobar",
		},
		{
			In:       "0",
			DataType: types.Generic,
			Out:      "0",
		},
		{
			In:       "true",
			DataType: types.Generic,
			Out:      "true",
		},
		{
			In:       "false",
			DataType: types.Generic,
			Out:      "false",
		},
		{
			In:       0,
			DataType: types.Generic,
			Out:      0,
		},
		{
			In:       float64(0),
			DataType: types.Generic,
			Out:      float64(0),
		},
		{
			In:       42,
			DataType: types.Generic,
			Out:      42,
		},
		{
			In:       42.12345,
			DataType: types.Generic,
			Out:      42.12345,
		},
		{
			In:       true,
			DataType: types.Generic,
			Out:      true,
		},
		{
			In:       false,
			DataType: types.Generic,
			Out:      false,
		},
		{
			In:       `{ out: "testing" }`,
			DataType: types.Generic,
			Out:      `{ out: "testing" }`,
		},
		{
			In:       `{ "foo": "bar" }`,
			DataType: types.Generic,
			Out:      `{ "foo": "bar" }`,
		},
	}

	testConvertGoType(t, tests)
}
