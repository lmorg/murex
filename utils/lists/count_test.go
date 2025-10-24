package lists_test

import (
	"testing"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/lists"
)

func TestCountInt(t *testing.T) {
	type testCountT struct {
		List   []int
		Output map[string]int
		Error  bool
	}

	tests := []testCountT{
		{
			List: []int{0, 1, 2, 3, 2, 1},
			Output: map[string]int{
				"0": 1,
				"1": 2,
				"2": 2,
				"3": 1,
			},
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		jExp := json.LazyLogging(test.Output)
		actual, err := lists.Count(test.List)
		jAct := json.LazyLogging(actual)

		if jExp != jAct || (err == nil) == test.Error {
			t.Errorf("Test %d failed", i)
			t.Logf("  List:     %s", json.LazyLogging(test.List))
			t.Logf("  Expected: %s", jExp)
			t.Logf("  Actual:   %s", jAct)
			t.Logf("  exp err:  %v", test.Error)
			t.Logf("  err msg:  %s", err)
		}
	}
}

func TestCountFloat(t *testing.T) {
	type testCountT struct {
		List   []float64
		Output map[string]int
		Error  bool
	}

	tests := []testCountT{
		{
			List: []float64{0, 1, 2, 3.3, 2, 1.1},
			Output: map[string]int{
				"0":   1,
				"1":   1,
				"1.1": 1,
				"2":   2,
				"3.3": 1,
			},
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		jExp := json.LazyLogging(test.Output)
		actual, err := lists.Count(test.List)
		jAct := json.LazyLogging(actual)

		if jExp != jAct || (err == nil) == test.Error {
			t.Errorf("Test %d failed", i)
			t.Logf("  List:     %s", json.LazyLogging(test.List))
			t.Logf("  Expected: %s", jExp)
			t.Logf("  Actual:   %s", jAct)
			t.Logf("  exp err:  %v", test.Error)
			t.Logf("  err msg:  %s", err)
		}
	}
}

func TestCountString(t *testing.T) {
	type testCountT struct {
		List   []string
		Output map[string]int
		Error  bool
	}

	tests := []testCountT{
		{
			List: []string{"0", "a", "b", "c", "b", "a"},
			Output: map[string]int{
				"0": 1,
				"a": 2,
				"b": 2,
				"c": 1,
			},
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		jExp := json.LazyLogging(test.Output)
		actual, err := lists.Count(test.List)
		jAct := json.LazyLogging(actual)

		if jExp != jAct || (err == nil) == test.Error {
			t.Errorf("Test %d failed", i)
			t.Logf("  List:     %s", json.LazyLogging(test.List))
			t.Logf("  Expected: %s", jExp)
			t.Logf("  Actual:   %s", jAct)
			t.Logf("  exp err:  %v", test.Error)
			t.Logf("  err msg:  %s", err)
		}
	}
}

func TestCountBool(t *testing.T) {
	type testCountT struct {
		List   []bool
		Output map[string]int
		Error  bool
	}

	tests := []testCountT{
		{
			List: []bool{true, false, true},
			Output: map[string]int{
				types.TrueString:  2,
				types.FalseString: 1,
			},
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		jExp := json.LazyLogging(test.Output)
		actual, err := lists.Count(test.List)
		jAct := json.LazyLogging(actual)

		if jExp != jAct || (err == nil) == test.Error {
			t.Errorf("Test %d failed", i)
			t.Logf("  List:     %s", json.LazyLogging(test.List))
			t.Logf("  Expected: %s", jExp)
			t.Logf("  Actual:   %s", jAct)
			t.Logf("  exp err:  %v", test.Error)
			t.Logf("  err msg:  %s", err)
		}
	}
}

func TestCountInterface(t *testing.T) {
	type testCountT struct {
		List   []any
		Output map[string]int
		Error  bool
	}

	tests := []testCountT{
		{
			List: []any{1.11, 2.2, 3, "a", "b", "c", true, 2.2, 1.11, 1.11, "b", "c", "c"},
			Output: map[string]int{
				"1.11":           3,
				"2.2":            2,
				"3":              1,
				"a":              1,
				"b":              2,
				"c":              3,
				types.TrueString: 1,
			},
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		jExp := json.LazyLogging(test.Output)
		actual, err := lists.Count(test.List)
		jAct := json.LazyLogging(actual)

		if jExp != jAct || (err == nil) == test.Error {
			t.Errorf("Test %d failed", i)
			t.Logf("  List:     %s", json.LazyLogging(test.List))
			t.Logf("  Expected: %s", jExp)
			t.Logf("  Actual:   %s", jAct)
			t.Logf("  exp err:  %v", test.Error)
			t.Logf("  err msg:  %s", err)
		}
	}
}
