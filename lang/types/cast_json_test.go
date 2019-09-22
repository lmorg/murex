package types_test

import (
	"testing"

	"github.com/lmorg/murex/lang/types"
)

func TestConvertGoTypeJson(t *testing.T) {
	tests := []test{
		{
			In:       nil,
			DataType: types.Json,
			Out:      `{}`,
		},
		{
			In:       "",
			DataType: types.Json,
			Out:      `{ "Value": "" }`,
		},
		{
			In:       "foobar",
			DataType: types.Json,
			Out:      `{ "Value": "foobar" }`,
		},
		{
			In:       "0",
			DataType: types.Json,
			Out:      `{ "Value": "0" }`,
		},
		{
			In:       "true",
			DataType: types.Json,
			Out:      `{ "Value": "true" }`,
		},
		{
			In:       "false",
			DataType: types.Json,
			Out:      `{ "Value": "false" }`,
		},
		{
			In:       0,
			DataType: types.Json,
			Out:      `{ "Value": 0 }`,
		},
		{
			In:       float64(0),
			DataType: types.Json,
			Out:      `{ "Value": 0 }`,
		},
		{
			In:       "42",
			DataType: types.Json,
			Out:      `{ "Value": "42" }`,
		},
		{
			In:       "42.12345",
			DataType: types.Json,
			Out:      `{ "Value": "42.12345" }`,
		},
		{
			In:       42,
			DataType: types.Json,
			Out:      `{ "Value": 42 }`,
		},
		{
			In:       42.12345,
			DataType: types.Json,
			Out:      `{ "Value": 42.12345 }`,
		},
		{
			In:       true,
			DataType: types.Json,
			Out:      `{ "Value": true }`,
		},
		{
			In:       false,
			DataType: types.Json,
			Out:      `{ "Value": false }`,
		},
		{
			In:       `{ out: "testing" }`,
			DataType: types.Json,
			Out:      `{ "Value": "{ out: \"testing\" }" }`,
		},
		{
			In:       `{ "foo": "bar" }`,
			DataType: types.Json,
			Out:      `{ "Value": "{ \"foo\": \"bar\" }" }`,
		},
	}

	testConvertGoType(t, tests)
}
