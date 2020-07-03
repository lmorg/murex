// +build ignore

package types_test

import (
	"testing"

	"github.com/lmorg/murex/lang/types"
)

func TestConvertGoTypeCodeBlock(t *testing.T) {
	tests := []test{
		{
			In:       nil,
			DataType: types.CodeBlock,
			Out:      "{}",
		},
		{
			In:       "",
			DataType: types.CodeBlock,
			Out:      "out: ''",
		},
		{
			In:       "foobar",
			DataType: types.CodeBlock,
			Out:      "out: 'foobar'",
		},
		{
			In:       "0",
			DataType: types.CodeBlock,
			Out:      "out: '0'",
		},
		{
			In:       "true",
			DataType: types.CodeBlock,
			Out:      "out: 'true'",
		},
		{
			In:       "false",
			DataType: types.CodeBlock,
			Out:      "out: 'false'",
		},
		{
			In:       0,
			DataType: types.CodeBlock,
			Out:      "out: 0",
		},
		{
			In:       float64(0),
			DataType: types.CodeBlock,
			Out:      "out: 0",
		},
		{
			In:       42,
			DataType: types.CodeBlock,
			Out:      "out: 42",
		},
		{
			In:       42.12345,
			DataType: types.CodeBlock,
			Out:      "out: 42.12345",
		},
		{
			In:       true,
			DataType: types.CodeBlock,
			Out:      "true",
		},
		{
			In:       false,
			DataType: types.CodeBlock,
			Out:      "false",
		},
		{
			In:       `{ out: "testing" }`,
			DataType: types.CodeBlock,
			Out:      ` out: "testing" `,
		},
		{
			In:       `{ "foo": "bar" }`,
			DataType: types.CodeBlock,
			Out:      ` "foo": "bar" `,
		},
	}

	testConvertGoType(t, tests)
}
