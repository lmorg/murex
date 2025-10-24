package types_test

import (
	"testing"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test/count"
)

type test struct {
	In       any
	DataType string
	Out      any
	Error    bool
}

func testConvertGoType(t *testing.T, tests []test) {
	count.Tests(t, len(tests))

	for i := range tests {
		var failed bool

		v, err := types.ConvertGoType(tests[i].In, tests[i].DataType)
		if (err != nil && !tests[i].Error) ||
			(err == nil && tests[i].Error) {

			t.Error("ConvertGoType failed with an error:")
			failed = true

		} else if tests[i].Out != v {
			t.Error("ConvertGoType out mismatch:")
			failed = true
		}

		if failed {
			t.Logf("  Test #:    %d", i)
			t.Logf("  Mx Type:   %s", tests[i].DataType)
			t.Logf("  In Type:   %T", tests[i].In)
			t.Logf("  Exp Type:  %T", tests[i].Out)
			t.Logf("  Act Type:  %T", v)
			t.Log("  In Value: ", tests[i].In)
			t.Log("  Exp Value:", tests[i].Out)
			t.Log("  Act Value:", v)
			t.Log("  Error:    ", err)
			t.Log("  Exp Err:  ", tests[i].Error)
		}
	}
}

func TestConvertGoTypeBaseline(t *testing.T) {
	tests := []test{
		{
			In:       nil,
			DataType: types.Null,
			Out:      "",
		},
		{
			In:       "foobar",
			DataType: types.String,
			Out:      "foobar",
		},
		{
			In:       42,
			DataType: types.Integer,
			Out:      42,
		},
		{
			In:       42.12345,
			DataType: types.Float,
			Out:      float64(42.12345),
		},
		{
			In:       42,
			DataType: types.Number,
			Out:      float64(42),
		},
		{
			In:       42.12345,
			DataType: types.Number,
			Out:      float64(42.12345),
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
			DataType: types.CodeBlock,
			Out:      ` out: "testing" `,
		},
		{
			In:       `foobar`,
			DataType: types.Json,
			Out:      `foobar`,
		},
	}

	testConvertGoType(t, tests)
}

// https://github.com/lmorg/murex/issues/829
func TestConvertGoTypeJsonToNumber(t *testing.T) {
	tests := []test{
		{
			In:       []any{1, 2, 3},
			DataType: types.Integer,
			Out:      0,
			Error:    true,
		},
		{
			In:       []any{1, 2, 3},
			DataType: types.Float,
			Out:      0,
			Error:    true,
		},
		{
			In:       []any{1, 2, 3},
			DataType: types.Number,
			Out:      0,
			Error:    true,
		},
	}

	testConvertGoType(t, tests)
}
