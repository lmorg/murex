package lists_test

import (
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/lists"
)

func TestSumInt(t *testing.T) {
	type testSumIntT struct {
		Source      map[string]int
		Destination map[string]int
		Output      map[string]int
	}

	tests := []testSumIntT{
		{
			Source: map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			Destination: map[string]int{
				"a": 9,
				"b": 7,
				"c": 5,
			},
			Output: map[string]int{
				"a": 10,
				"b": 9,
				"c": 8,
			},
		},
		{
			Source: map[string]int{
				"a": 1,
			},
			Destination: map[string]int{},
			Output: map[string]int{
				"a": 1,
			},
		},
		{
			Source: map[string]int{},
			Destination: map[string]int{
				"c": 5,
			},
			Output: map[string]int{
				"c": 5,
			},
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		jExp := json.LazyLogging(test.Output)

		actual := make(map[string]int)
		for k, v := range test.Destination {
			actual[k] = v
		}
		lists.SumInt(actual, test.Source)
		jAct := json.LazyLogging(actual)

		if jExp != jAct {
			t.Errorf("Test %d failed", i)
			t.Logf("  Source:      %s", json.LazyLogging(test.Source))
			t.Logf("  Destination: %s", json.LazyLogging(test.Destination))
			t.Logf("  Expected:    %s", jExp)
			t.Logf("  Actual:      %s", jAct)
		}
	}
}

func TestSumFloat64(t *testing.T) {
	type testSumFloat64T struct {
		Source      map[string]float64
		Destination map[string]float64
		Output      map[string]float64
	}

	tests := []testSumFloat64T{
		{
			Source: map[string]float64{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			Destination: map[string]float64{
				"a": 9,
				"b": 7,
				"c": 5,
			},
			Output: map[string]float64{
				"a": 10,
				"b": 9,
				"c": 8,
			},
		},
		{
			Source: map[string]float64{
				"a": 1,
			},
			Destination: map[string]float64{},
			Output: map[string]float64{
				"a": 1,
			},
		},
		{
			Source: map[string]float64{},
			Destination: map[string]float64{
				"c": 5,
			},
			Output: map[string]float64{
				"c": 5,
			},
		},
		{
			Source: map[string]float64{
				"a": 1.1,
				"b": 2.5,
				"c": 3.9,
			},
			Destination: map[string]float64{
				"a": 9.1,
				"b": 7.5,
				"c": 5.9,
			},
			Output: map[string]float64{
				"a": 10.2,
				"b": 10,
				"c": 9.8,
			},
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		jExp := json.LazyLogging(test.Output)

		actual := make(map[string]float64)
		for k, v := range test.Destination {
			actual[k] = v
		}
		lists.SumFloat64(actual, test.Source)
		jAct := json.LazyLogging(actual)

		if jExp != jAct {
			t.Errorf("Test %d failed", i)
			t.Logf("  Source:      %s", json.LazyLogging(test.Source))
			t.Logf("  Destination: %s", json.LazyLogging(test.Destination))
			t.Logf("  Expected:    %s", jExp)
			t.Logf("  Actual:      %s", jAct)
		}
	}
}

func TestSumInterface(t *testing.T) {
	type testSumInterfaceT struct {
		Source      map[string]any
		Destination map[string]any
		Output      map[string]any
		Error       bool
	}

	tests := []testSumInterfaceT{
		{
			Source: map[string]any{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			Destination: map[string]any{
				"a": 9,
				"b": 7,
				"c": 5,
			},
			Output: map[string]any{
				"a": 10,
				"b": 9,
				"c": 8,
			},
		},
		{
			Source: map[string]any{
				"a": "1",
				"b": "2",
				"c": "3",
			},
			Destination: map[string]any{
				"a": "9",
				"b": "7",
				"c": "5",
			},
			Output: map[string]any{
				"a": 10,
				"b": 9,
				"c": 8,
			},
		},
		{
			Source: map[string]any{
				"a": 1,
			},
			Destination: map[string]any{},
			Output: map[string]any{
				"a": 1,
			},
		},
		{
			Source: map[string]any{},
			Destination: map[string]any{
				"c": 5,
			},
			Output: map[string]any{
				"c": 5,
			},
		},
		{
			Source: map[string]any{
				"a": "1",
			},
			Destination: map[string]any{
				"c": 5,
			},
			Output: map[string]any{
				"a": 1,
				"c": 5,
			},
		},
		{
			Source: map[string]any{
				"a": 1,
			},
			Destination: map[string]any{
				"c": "5",
			},
			Output: map[string]any{
				"a": 1,
				"c": "5",
			},
		},
		{
			Source: map[string]any{
				"a": 1.1,
				"b": 2.5,
				"c": 3.9,
			},
			Destination: map[string]any{
				"a": 9.1,
				"b": 7.5,
				"c": 5.9,
			},
			Output: map[string]any{
				"a": 10.2,
				"b": 10,
				"c": 9.8,
			},
		},
		{
			Source: map[string]any{
				"a": 1.1,
				"b": "2.5",
				"c": 3.9,
			},
			Destination: map[string]any{
				"a": "9.1",
				"b": 7.5,
				"c": "5.9",
			},
			Output: map[string]any{
				"a": 10.2,
				"b": 10,
				"c": 9.8,
			},
		},
		{
			Source: map[string]any{
				"a":   1.1,
				"b":   "2.5",
				"c":   3.9,
				"foo": "bar",
			},
			Destination: map[string]any{
				"a":   "9.1",
				"b":   7.5,
				"c":   "5.9",
				"bar": "foo",
			},
			Output: map[string]any{
				"a":   10.2,
				"b":   10,
				"c":   9.8,
				"bar": "foo",
			},
			Error: true,
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		jExp := json.LazyLogging(test.Output)

		actual := make(map[string]any)
		for k, v := range test.Destination {
			actual[k] = v
		}
		err := lists.SumInterface(actual, test.Source)
		jAct := json.LazyLogging(actual)

		var errStr string

		if jExp != jAct && !test.Error {
			errStr += ". Expected != Actual"
		}

		if (err == nil) == test.Error {
			errStr += ". Error mismatch"
		}

		if errStr != "" {
			t.Errorf("Test %d failed%s", i, errStr)
			t.Logf("  Source:      %s", json.LazyLogging(test.Source))
			t.Logf("  Destination: %s", json.LazyLogging(test.Destination))
			t.Logf("  Expected:    %s", jExp)
			t.Logf("  Actual:      %s", jAct)
			t.Logf("  err exp:     %v", test.Error)
			t.Logf("  err message: %s", err)
		}
	}
}
