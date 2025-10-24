package lang_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
)

type testArrayWithTypeTemplateT struct {
	Value    string
	ExpV     string // if different from value
	Expected []string
	Error    bool
}

func testArrayWithTypeTemplate(t *testing.T, dataType string, tests []testArrayWithTypeTemplateT) {
	t.Helper()
	count.Tests(t, len(tests))

	marshal := func(v any) ([]byte, error) {
		p := lang.NewTestProcess()
		return lang.MarshalData(p, dataType, v)
	}

	unmarshal := func(b []byte, v any) error {
		p := lang.NewTestProcess()
		p.Stdin = streams.NewStdin()
		_, err := p.Stdin.Write(b)
		if err != nil {
			return err
		}
		ret, err := lang.UnmarshalData(p, dataType)
		if err != nil {
			return err
		}
		ref := reflect.ValueOf(v).Elem()
		ref.Set(reflect.ValueOf(ret))
		return nil
	}

	for i, test := range tests {
		stream := streams.NewStdin()
		_, err := stream.Write([]byte(test.Value))
		if err != nil {
			panic(fmt.Sprintf("error writing to stream in test %d: %v", i, err))
		}

		var (
			values    []any
			dataTypes []string
		)
		callback := func(v any, dt string) {
			switch t := v.(type) {
			case []byte:
				values = append(values, string(t))
			case []rune:
				values = append(values, string(t))
			default:
				values = append(values, v)
			}
			dataTypes = append(dataTypes, dt)
		}

		err = lang.ArrayWithTypeTemplate(context.TODO(),
			types.Json, marshal, unmarshal, stream, callback)

		if (err != nil) != test.Error {
			t.Errorf("Error returned from lang.ArrayWithTypeTemplate() doesn't match expected:")
			t.Logf("  Test:      %d", i)
			t.Logf("  exp error: %v", test.Error)
			t.Logf("  act error: %v", err)
		}

		actVb, err := marshal(values)
		actV := string(actVb)
		actDt := json.LazyLogging(dataTypes)
		expV := test.ExpV
		if expV == "" {
			expV = test.Value
		}
		expDt := json.LazyLogging(test.Expected)

		if expV != actV || expDt != actDt || err != nil {
			t.Errorf("Mismatch in arrays:")
			t.Logf("  Test:      %d", i)
			t.Logf("  exp value: %s", expV)
			t.Logf("  act value: %s", actV)
			t.Logf("  exp types: %s", expDt)
			t.Logf("  act types: %s", actDt)
			t.Logf("  exp error: %v", test.Error)
			t.Logf("  act error: %v", err)
		}
	}
}

func TestArrayWithTypeTemplateJson(t *testing.T) {
	tests := []testArrayWithTypeTemplateT{
		{
			Value:    `[1,2,3]`,
			Expected: []string{types.Number, types.Number, types.Number},
		},
		{
			Value:    `[1.9,2.8,3.7]`,
			Expected: []string{types.Number, types.Number, types.Number},
		},
		{
			Value:    `["a","b","c"]`,
			Expected: []string{types.String, types.String, types.String},
		},
		{
			Value:    `[true,false,true]`,
			Expected: []string{types.Boolean, types.Boolean, types.Boolean},
		},
		{
			Value:    `[null,null,null]`,
			Expected: []string{types.Null, types.Null, types.Null},
		},
	}

	testArrayWithTypeTemplate(t, types.Json, tests)
}

func TestArrayWithTypeTemplateJsonStructs(t *testing.T) {
	tests := []testArrayWithTypeTemplateT{
		{
			Value:    `[[1,2,3],[4,5,6],[7,8,9]]`,
			ExpV:     `["[1,2,3]","[4,5,6]","[7,8,9]"]`,
			Expected: []string{types.Json, types.Json, types.Json},
		},
		{
			Value:    `["{\"a\":[1,2,3],\"b\":[4,5,6],\"c\":[7,8,9]}"]`,
			Expected: []string{types.String},
		},
	}

	testArrayWithTypeTemplate(t, types.Json, tests)
}

func TestArrayWithTypeTemplateString(t *testing.T) {
	tests := []testArrayWithTypeTemplateT{
		{
			Value:    "abc\ndef\nghi\n",
			Expected: []string{types.String, types.String, types.String},
		},
		{
			Value:    "true\nfalse\ntrue\n",
			Expected: []string{types.String, types.String, types.String},
		},
		{
			Value:    "1\n2\n3\n",
			Expected: []string{types.String, types.String, types.String},
		},
	}

	testArrayWithTypeTemplate(t, types.String, tests)
}
