package mxjson_test

import "testing"

func TestBasicMapsBoolean(t *testing.T) {
	tests := []testCase{
		{
			Json:     `{"foo": true}`,
			Expected: `{"foo":true}`,
		},
		{
			Json:     `{'foo': false}`,
			Expected: `{"foo":false}`,
		},
		{
			Json:     `{"foo":true}`,
			Expected: `{"foo":true}`,
		},

		{
			Json:     `{"1foo": true, "2foo": false}`,
			Expected: `{"1foo":true,"2foo":false}`,
		},
		{
			Json:     `{"1foo": false, "2foo": true, "3foo": false}`,
			Expected: `{"1foo":false,"2foo":true,"3foo":false}`,
		},

		{
			Json:     `{"1": {"11foo": true} }`,
			Expected: `{"1":{"11foo":true}}`,
		},
		{
			Json:     `{"1": {"11foo": true, "12foo": false} }`,
			Expected: `{"1":{"11foo":true,"12foo":false}}`,
		},
		{
			Json:     `{"1": {"11foo": false, "12foo": true}, "2": {"21foo": true, "22foo": false} }`,
			Expected: `{"1":{"11foo":false,"12foo":true},"2":{"21foo":true,"22foo":false}}`,
		},
	}

	runTestCases(t, tests)
}

func TestBasicMapsNumbers(t *testing.T) {
	tests := []testCase{
		{
			Json:     `{"foo": 1}`,
			Expected: `{"foo":1}`,
		},
		{
			Json:     `{'foo': 1 }`,
			Expected: `{"foo":1}`,
		},
		{
			Json:     `{"foo":1}`,
			Expected: `{"foo":1}`,
		},

		{
			Json:     `{"1foo": 1, "2foo": 2}`,
			Expected: `{"1foo":1,"2foo":2}`,
		},
		{
			Json:     `{"1foo": 1, "2foo": 2, "3foo": 3}`,
			Expected: `{"1foo":1,"2foo":2,"3foo":3}`,
		},

		{
			Json:     `{"1": {"11foo": 11} }`,
			Expected: `{"1":{"11foo":11}}`,
		},
		{
			Json:     `{"1": {"11foo": 11, "12foo": 12} }`,
			Expected: `{"1":{"11foo":11,"12foo":12}}`,
		},
		{
			Json:     `{"1": {"11foo": 11, "12foo": 12}, "2": {"21foo": 21, "22foo": 22} }`,
			Expected: `{"1":{"11foo":11,"12foo":12},"2":{"21foo":21,"22foo":22}}`,
		},
	}

	runTestCases(t, tests)
}

func TestBasicMaps(t *testing.T) {
	tests := []testCase{
		{
			Json:     `{"foo": "bar"}`,
			Expected: `{"foo":"bar"}`,
		},
		{
			Json:     `{'foo': 'bar'}`,
			Expected: `{"foo":"bar"}`,
		},
		{
			Json:     `{"foo": (bar)}`,
			Expected: `{"foo":"bar"}`,
		},

		{
			Json:     `{"1foo": "1bar", "2foo": "2bar"}`,
			Expected: `{"1foo":"1bar","2foo":"2bar"}`,
		},
		{
			Json:     `{"1foo": "1bar", "2foo": "2bar", "3foo": "3bar"}`,
			Expected: `{"1foo":"1bar","2foo":"2bar","3foo":"3bar"}`,
		},

		{
			Json:     `{"1": {"11foo": "11bar"} }`,
			Expected: `{"1":{"11foo":"11bar"}}`,
		},
		{
			Json:     `{"1": {"11foo": "11bar", "12foo": "12bar"} }`,
			Expected: `{"1":{"11foo":"11bar","12foo":"12bar"}}`,
		},
		{
			Json:     `{"1": {"11foo": "11bar", "12foo": "12bar"}, "2": {"21foo": "21bar", "22foo": "22bar"} }`,
			Expected: `{"1":{"11foo":"11bar","12foo":"12bar"},"2":{"21foo":"21bar","22foo":"22bar"}}`,
		},
	}

	runTestCases(t, tests)
}

func TestBasicArrayBoolean(t *testing.T) {
	tests := []testCase{
		{
			Json:     `[true, false, false, true]`,
			Expected: `[true,false,false,true]`,
		},
	}

	runTestCases(t, tests)
}

func TestBasicArrayNumber(t *testing.T) {
	tests := []testCase{
		{
			Json:     `[1, 3, 2, 4]`,
			Expected: `[1,3,2,4]`,
		},
	}

	runTestCases(t, tests)
}

func TestBasicArrayString(t *testing.T) {
	tests := []testCase{
		{
			Json:     `["1one", "2two", "3three", "4four"]`,
			Expected: `["1one","2two","3three","4four"]`,
		},
	}

	runTestCases(t, tests)
}
