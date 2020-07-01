// +build ignore

package types_test

import (
	"testing"

	"github.com/lmorg/murex/lang/types"
)

func TestConvertGoTypeJsonLines(t *testing.T) {
	tests := []test{
		{
			In:       nil,
			DataType: types.JsonLines,
			Out:      `{}`,
		},
		{
			In:       "",
			DataType: types.JsonLines,
			Out:      `{ "Value": "" }`,
		},
		{
			In:       "foobar",
			DataType: types.JsonLines,
			Out:      `{ "Value": "foobar" }`,
		},
		{
			In:       "0",
			DataType: types.JsonLines,
			Out:      `{ "Value": "0" }`,
		},
		{
			In:       "true",
			DataType: types.JsonLines,
			Out:      `{ "Value": "true" }`,
		},
		{
			In:       "false",
			DataType: types.JsonLines,
			Out:      `{ "Value": "false" }`,
		},
		{
			In:       0,
			DataType: types.JsonLines,
			Out:      `{ "Value": 0 }`,
		},
		{
			In:       float64(0),
			DataType: types.JsonLines,
			Out:      `{ "Value": 0 }`,
		},
		{
			In:       "42",
			DataType: types.JsonLines,
			Out:      `{ "Value": "42" }`,
		},
		{
			In:       "42.12345",
			DataType: types.JsonLines,
			Out:      `{ "Value": "42.12345" }`,
		},
		{
			In:       42,
			DataType: types.JsonLines,
			Out:      `{ "Value": 42 }`,
		},
		{
			In:       42.12345,
			DataType: types.JsonLines,
			Out:      `{ "Value": 42.12345 }`,
		},
		{
			In:       true,
			DataType: types.JsonLines,
			Out:      `{ "Value": true }`,
		},
		{
			In:       false,
			DataType: types.JsonLines,
			Out:      `{ "Value": false }`,
		},
		{
			In:       `{ out: "testing" }`,
			DataType: types.JsonLines,
			Out:      `{ "Value": "{ out: \"testing\" }" }`,
		},
		{
			In:       `{ "foo": "bar" }`,
			DataType: types.JsonLines,
			Out:      `{ "Value": "{ \"foo\": \"bar\" }" }`,
		},
	}

	testConvertGoType(t, tests)
}
