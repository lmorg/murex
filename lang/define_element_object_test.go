package lang_test

import (
	"testing"

	"github.com/lmorg/murex/lang"
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
			Path:  ".-1",
			Error: true,
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
			Path:  ".-1",
			Error: true,
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

		v, err := lang.ElementLookup(test.Object, test.Path, "")
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
