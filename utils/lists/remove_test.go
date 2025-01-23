package lists_test

import (
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/lists"
)

func TestRemoveOrdered(t *testing.T) {
	tests := []struct {
		Slice    []any
		Index    int
		Expected []any
		Error    bool
	}{
		{
			Slice:    []any{"foo", "bar", "baz"},
			Index:    -1,
			Expected: nil,
			Error:    true,
		},
		{
			Slice:    []any{"foo", "bar", "baz"},
			Index:    0,
			Expected: []any{"bar", "baz"},
		},
		{
			Slice:    []any{"foo", "bar", "baz"},
			Index:    1,
			Expected: []any{"foo", "baz"},
		},
		{
			Slice:    []any{"foo", "bar", "baz"},
			Index:    2,
			Expected: []any{"foo", "bar"},
		},
		{
			Slice:    []any{"foo", "bar", "baz"},
			Index:    3,
			Expected: nil,
			Error:    true,
		},

		/////

		{
			Slice:    []any{5, 10, 15},
			Index:    -1,
			Expected: nil,
			Error:    true,
		},
		{
			Slice:    []any{5, 10, 15},
			Index:    0,
			Expected: []any{10, 15},
		},
		{
			Slice:    []any{5, 10, 15},
			Index:    1,
			Expected: []any{5, 15},
		},
		{
			Slice:    []any{5, 10, 15},
			Index:    2,
			Expected: []any{5, 10},
		},
		{
			Slice:    []any{5, 10, 15},
			Index:    3,
			Expected: nil,
			Error:    true,
		},

		/////

		{
			Slice:    []any{'b', 'a', 'r'},
			Index:    -1,
			Expected: nil,
			Error:    true,
		},
		{
			Slice:    []any{'b', 'a', 'r'},
			Index:    0,
			Expected: []any{'a', 'r'},
		},
		{
			Slice:    []any{'b', 'a', 'r'},
			Index:    1,
			Expected: []any{'b', 'r'},
		},
		{
			Slice:    []any{'b', 'a', 'r'},
			Index:    2,
			Expected: []any{'b', 'a'},
		},
		{
			Slice:    []any{'b', 'a', 'r'},
			Index:    3,
			Expected: nil,
			Error:    true,
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		expJson := json.LazyLogging(test.Expected)
		actual, err := lists.RemoveOrdered(test.Slice, test.Index)
		actJson := json.LazyLogging(actual)

		if expJson != actJson || ((err != nil) != test.Error) {
			t.Errorf("Unexpected return in test %d:", i)
			t.Logf("  Slice:    %s", json.LazyLogging(test.Slice))
			t.Logf("  Index:    %d", test.Index)
			t.Logf("  Expected: %s", expJson)
			t.Logf("  Actual:   %s", actJson)
			t.Logf("  exp err:  %v", test.Error)
			t.Logf("  act err:  %v", err)
		}

	}
}

func TestRemoveUnordered(t *testing.T) {
	tests := []struct {
		Slice    []any
		Index    int
		Expected []any
		Error    bool
	}{
		{
			Slice:    []any{"foo", "bar", "baz"},
			Index:    -1,
			Expected: nil,
			Error:    true,
		},
		{
			Slice:    []any{"foo", "bar", "baz"},
			Index:    0,
			Expected: []any{"baz", "bar"},
		},
		{
			Slice:    []any{"foo", "bar", "baz"},
			Index:    1,
			Expected: []any{"foo", "baz"},
		},
		{
			Slice:    []any{"foo", "bar", "baz"},
			Index:    2,
			Expected: []any{"foo", "bar"},
		},
		{
			Slice:    []any{"foo", "bar", "baz"},
			Index:    3,
			Expected: nil,
			Error:    true,
		},

		/////

		{
			Slice:    []any{5, 10, 15},
			Index:    -1,
			Expected: nil,
			Error:    true,
		},
		{
			Slice:    []any{5, 10, 15},
			Index:    0,
			Expected: []any{15, 10},
		},
		{
			Slice:    []any{5, 10, 15},
			Index:    1,
			Expected: []any{5, 15},
		},
		{
			Slice:    []any{5, 10, 15},
			Index:    2,
			Expected: []any{5, 10},
		},
		{
			Slice:    []any{5, 10, 15},
			Index:    3,
			Expected: nil,
			Error:    true,
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		expJson := json.LazyLogging(test.Expected)
		actual, err := lists.RemoveUnordered(test.Slice, test.Index)
		actJson := json.LazyLogging(actual)

		if expJson != actJson || ((err != nil) != test.Error) {
			t.Errorf("Unexpected return in test %d:", i)
			t.Logf("  Slice:    %s", json.LazyLogging(test.Slice))
			t.Logf("  Index:    %d", test.Index)
			t.Logf("  Expected: %s", expJson)
			t.Logf("  Actual:   %s", actJson)
			t.Logf("  exp err:  %v", test.Error)
			t.Logf("  act err:  %v", err)
		}

	}
}
