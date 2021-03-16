package envvars

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/lmorg/murex/test/count"
)

type envAllTestT struct {
	Slice    []string
	Expected map[string]string
	Error    bool
}

func TestEnvVarsAll(t *testing.T) {
	tests := []envAllTestT{
		{
			Slice:    []string{},
			Expected: map[string]string{},
			Error:    false,
		},
		{
			Slice: []string{
				`a=a`,
				`b=b`,
				`c=c`,
			},
			Expected: map[string]string{
				"a": "a",
				"b": "b",
				"c": "c",
			},
			Error: false,
		},
		{
			Slice: []string{
				`a=`,
				`b=b`,
				`c=c`,
			},
			Expected: map[string]string{
				"a": "",
				"b": "b",
				"c": "c",
			},
			Error: false,
		},
		{
			Slice: []string{
				`a`,
				`b=b`,
				`c=c`,
			},
			Expected: map[string]string{
				"a": "",
				"b": "b",
				"c": "c",
			},
			Error: false,
		},
		{
			Slice: []string{
				`a=a=a`,
				`b=b`,
				`c=c`,
			},
			Expected: map[string]string{
				"a": "a=a",
				"b": "b",
				"c": "c",
			},
			Error: false,
		},
		{
			Slice: []string{
				`a=a`,
				`foo=bar`,
				`c=c`,
			},
			Expected: map[string]string{
				"a":   "a",
				"foo": "bar",
				"c":   "c",
			},
			Error: false,
		},
		{
			Slice: []string{
				`a=a`,
				`=foobar`,
				`c=c`,
			},
			Expected: map[string]string{
				"a": "a",
				"":  "foobar",
				"c": "c",
			},
			Error: false,
		},
		{
			Slice: []string{
				`a=a`,
				`foobar=`,
				`c=c`,
			},
			Expected: map[string]string{
				"a":      "a",
				"foobar": "",
				"c":      "c",
			},
			Error: false,
		},
		{
			Slice: []string{
				`a=a`,
				`=`,
				`c=c`,
			},
			Expected: map[string]string{
				"a": "a",
				"":  "",
				"c": "c",
			},
			Error: false,
		},
		{
			Slice: []string{
				`a=a`,
				`foobar`,
				`c=c`,
			},
			Expected: map[string]string{
				"a":      "a",
				"foobar": "",
				"c":      "c",
			},
			Error: false,
		},
		{
			Slice: []string{
				`a=a`,
				``,
				`c=c`,
			},
			Expected: map[string]string{
				"a": "a",
				"":  "",
				"c": "c",
			},
			Error: false,
		},
		{
			Slice: []string{
				`a=a`,
				`hello=世界`,
				`c=c`,
			},
			Expected: map[string]string{
				"a":     "a",
				"hello": "世界",
				"c":     "c",
			},
			Error: false,
		},
		{
			Slice: []string{
				`a=a`,
				`世界=hello`,
				`c=c`,
			},
			Expected: map[string]string{
				"a":  "a",
				"世界": "hello",
				"c":  "c",
			},
			Error: false,
		},
		{
			Slice: []string{
				`a=a`,
				`世界=`,
				`c=c`,
			},
			Expected: map[string]string{
				"a":  "a",
				"世界": "",
				"c":  "c",
			},
			Error: false,
		},
		{
			Slice: []string{
				`a=a`,
				`世界`,
				`c=c`,
			},
			Expected: map[string]string{
				"a":  "a",
				"世界": "",
				"c":  "c",
			},
			Error: false,
		},
		{
			Slice: []string{
				`a=a`,
				`=世界`,
				`c=c`,
			},
			Expected: map[string]string{
				"a": "a",
				"":  "世界",
				"c": "c",
			},
			Error: false,
		},
		{
			Slice: []string{
				"a=a",
				"TIMEFMT=\n================\nCPU	%P\nuser	%*U\nsystem	%*S\ntotal	%*E",
				`c=c`,
			},
			Expected: map[string]string{
				"a": "a",
				"TIMEFMT": "\n================\nCPU	%P\nuser	%*U\nsystem	%*S\ntotal	%*E",
				"c": "c",
			},
			Error: false,
		},
	}

	testEnvVarsAll(t, tests)
}

func testEnvVarsAll(t *testing.T, tests []envAllTestT) {
	count.Tests(t, len(tests))
	t.Helper()

	for i, test := range tests {
		if test.Expected == nil {
			test.Expected = make(map[string]string)
		}

		actual := make(map[string]interface{})
		all(test.Slice, actual)

		/*if (err != nil) != test.Error {
			t.Errorf("Error expectation mismatch in test %d", i)
			t.Logf("  Expected: %v", test.Error)
			t.Logf("  Actual:   %v", err)
		}*/

		if len(test.Expected) != len(actual) {
			t.Errorf("Output count mistmatch in test %d", i)
			t.Logf("  Exp Count: %d", len(test.Expected))
			t.Logf("  Act Count: %d", len(actual))
			t.Logf("  Expected:\n%s", testJsonEncode(test.Expected))
			t.Logf("  Actual:\n%s", testJsonEncode(actual))
		}

		for k := range actual {
			if actual[k] != test.Expected[k] {
				t.Errorf("Key/value mistmatch in test %d", i)
				t.Logf("  Key:       `%s`", k)
				t.Logf("  Exp Value: `%s`", test.Expected[k])
				t.Logf("  Act Value: `%s`", actual[k].(string))
				t.Logf("  Expected:\n%s", testJsonEncode(test.Expected))
				t.Logf("  Actual:\n%s", testJsonEncode(actual))
			}
		}
	}
}

func testJsonEncode(v interface{}) string {
	b, _ := json.MarshalIndent(v, "    ", "    ")
	return string(b)
}

func TestEnvVarsE2E(t *testing.T) {
	namespace := "MUREX_TEST_TestDumpExports_"

	tests := map[string]string{
		"nil":       "",
		"text":      "testing123",
		"s p a c e": "testing 123",
		"equals":    "testing=123",
	}

	count.Tests(t, len(tests))

	for k, v := range tests {
		if err := os.Setenv(namespace+k, v); err != nil {
			t.Error("Unable to set env var for test!")
			t.Logf("  Name:  '%s'", namespace+k)
			t.Logf("  Value: '%s'", v)
			t.Logf("  Error:  %s", err)
		}
	}

	m := make(map[string]interface{})
	All(m)

	for k, v := range tests {
		if m[namespace+k] != v {
			t.Error("Unexpected map value:")
			t.Logf("  Name:     '%s'", namespace+k)
			t.Logf("  Expected: '%s'", v)
			t.Logf("  Actual:   '%s'", m[namespace+k].(string))
		}

		if err := os.Unsetenv(namespace + k); err != nil {
			t.Errorf("Unable to unset '%s': %s", namespace+k, err)
		}
	}
}
