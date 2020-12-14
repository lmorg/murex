package mxjson_test

import "testing"

func TestQuoteBrace(t *testing.T) {
	tests := []testCase{
		{
			Json:     `{"foo": ({foobar}) }`,
			Expected: `{"foo":"{foobar}"}`,
		},
		{
			Json:     `{"foo": ({true}) }`,
			Expected: `{"foo":"{true}"}`,
		},
		{
			Json:     `{"foo": ({false}) }`,
			Expected: `{"foo":"{false}"}`,
		},
		{
			Json:     `{"foo": ({1}) }`,
			Expected: `{"foo":"{1}"}`,
		},
		{
			Json:     `{"foo": ({1.3}) }`,
			Expected: `{"foo":"{1.3}"}`,
		},
		{
			Json:     `{"foo": ({1 2 3}) }`,
			Expected: `{"foo":"{1 2 3}"}`,
		},
		{
			Json:     `{"foo": ({@[foobar]}) }`,
			Expected: `{"foo":"{@[foobar]}"}`,
		},
	}

	runTestCases(t, tests)
}
