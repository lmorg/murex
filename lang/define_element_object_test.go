package lang

import (
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
)

type testElementLookupT struct {
	Object   interface{}
	Path     string
	Expected interface{}
	Error    bool
}

func TestElementLookup(t *testing.T) {
	tests := []testElementLookupT{
		{
			Object: map[string]any{
				"one": map[string]any{"two": map[string]any{"three": "four"}},
			},
			Path:     "/one/two/three",
			Expected: "four",
		},
		{
			Object: map[string]any{
				"one": map[string]any{"two": map[string]any{"three": "four"}},
			},
			Path:     "😅one😅two😅three",
			Expected: "four",
		},

		{
			Object: []string{
				"foo", "bar",
			},
			Path:     "/1",
			Expected: "bar",
		},
		{
			Object: []string{
				"foo", "bar",
			},
			Path:     ".1",
			Expected: "bar",
		},
		{
			Object: []string{
				"foo", "bar",
			},
			Path:     ".-1",
			Expected: "bar",
		},

		/////

		{
			Object: []interface{}{
				"foo", "bar",
			},
			Path:     "/1",
			Expected: "bar",
		},
		{
			Object: []interface{}{
				"foo", "bar",
			},
			Path:     ".1",
			Expected: "bar",
		},
		{
			Object: []interface{}{
				"foo", "bar",
			},
			Path:     ".-1",
			Expected: "bar",
		},

		/////

		{
			Object: map[string]string{
				"Simpsons":   "Homer",
				"Futurama":   "Bender",
				"South Park": "Cartman",
			},
			Path:     "/Simpsons",
			Expected: "Homer",
		},
		{
			Object: map[string]string{
				"Simpsons":   "Homer",
				"Futurama":   "Bender",
				"South Park": "Cartman",
			},
			Path:     ",futurama",
			Expected: "Bender",
		},
		{
			Object: map[string]string{
				"Simpsons":   "Homer",
				"Futurama":   "Bender",
				"SOUTH PARK": "Cartman",
			},
			Path:     ".South Park",
			Expected: "Cartman",
		},
		{
			Object: map[string]string{
				"Simpsons":   "Homer",
				"Futurama":   "Bender",
				"south park": "Cartman",
			},
			Path:     ".South Park",
			Expected: "Cartman",
		},

		/////

		{
			Object: map[string]interface{}{
				"Simpsons":   "Homer",
				"Futurama":   "Bender",
				"South Park": "Cartman",
			},
			Path:     "/Simpsons",
			Expected: "Homer",
		},
		{
			Object: map[string]interface{}{
				"Simpsons":   "Homer",
				"Futurama":   "Bender",
				"South Park": "Cartman",
			},
			Path:     ",futurama",
			Expected: "Bender",
		},
		{
			Object: map[string]interface{}{
				"Simpsons":   "Homer",
				"Futurama":   "Bender",
				"SOUTH PARK": "Cartman",
			},
			Path:     ".South Park",
			Expected: "Cartman",
		},
		{
			Object: map[string]interface{}{
				"Simpsons":   "Homer",
				"Futurama":   "Bender",
				"south park": "Cartman",
			},
			Path:     ".South Park",
			Expected: "Cartman",
		},

		/////

		{
			Object: map[interface{}]interface{}{
				"Simpsons":   "Homer",
				"Futurama":   "Bender",
				"South Park": "Cartman",
			},
			Path:     "/Simpsons",
			Expected: "Homer",
		},
		{
			Object: map[interface{}]interface{}{
				"Simpsons":   "Homer",
				"Futurama":   "Bender",
				"South Park": "Cartman",
			},
			Path:     ",futurama",
			Expected: "Bender",
		},
		{
			Object: map[interface{}]interface{}{
				"Simpsons":   "Homer",
				"Futurama":   "Bender",
				"SOUTH PARK": "Cartman",
			},
			Path:     ".South Park",
			Expected: "Cartman",
		},
		{
			Object: map[interface{}]interface{}{
				"Simpsons":   "Homer",
				"Futurama":   "Bender",
				"south park": "Cartman",
			},
			Path:     ".South Park",
			Expected: "Cartman",
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		expected := json.LazyLogging(test.Expected)

		v, err := ElementLookup(test.Object, test.Path)
		actual := json.LazyLogging(v)

		if (err != nil) != test.Error || actual != expected {
			t.Errorf("Test %d failed", i)
			t.Logf("  Object:   %s", json.LazyLogging(test.Object))
			t.Logf("  Path:     '%s'", test.Path)
			t.Logf("  Expected: %s", expected)
			t.Logf("  Actual:   %s", actual)
			t.Logf("  err exp:  %v", test.Error)
			t.Logf("  err act:  %v", err)
		}
	}
}

func TestIsValidElementIndex(t *testing.T) {
	tests := []struct {
		Key    string
		Length int
		Index  int
		Error  bool
	}{
		{
			Key:    "-",
			Length: 0,
			Index:  0,
			Error:  true,
		},
		{
			Key:    "0",
			Length: 0,
			Index:  0,
			Error:  true,
		},
		{
			Key:    "0",
			Length: 1,
			Index:  0,
			Error:  false,
		},
		{
			Key:    "3",
			Length: 2,
			Index:  0,
			Error:  true,
		},
		{
			Key:    "-1",
			Length: 7,
			Index:  6,
			Error:  false,
		},
		{
			Key:    "-10",
			Length: 7,
			Index:  0,
			Error:  true,
		},
		{
			Key:    "-5",
			Length: 5,
			Index:  0,
			Error:  false,
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		index, err := isValidElementIndex(test.Key, test.Length)

		if index != test.Index || (err != nil) != test.Error {
			t.Errorf("Unexpected return in test %d: ", i)
			t.Logf("  Key:      '%s'", test.Key)
			t.Logf("  Length:    %d", test.Length)
			t.Logf("  exp index: %d", test.Index)
			t.Logf("  act index: %d", index)
			t.Logf("  exp err:   %v", test.Error)
			t.Logf("  act err:   %v", err)
		}

	}
}
