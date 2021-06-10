package dedup_test

import (
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/dedup"
)

type testSortT struct {
	Input     []string
	Expected  []string
	ExpLength int
}

func TestSortAndDedupString(t *testing.T) {
	tests := []testSortT{
		{
			Input:     []string{},
			Expected:  []string{},
			ExpLength: 0,
		},
		{
			Input:     []string{"a", "a"},
			Expected:  []string{"a"},
			ExpLength: 1,
		},
		{
			Input:     []string{"a", "a", "a"},
			Expected:  []string{"a"},
			ExpLength: 1,
		},
		{
			Input:     []string{"a", "a", "b"},
			Expected:  []string{"a", "b"},
			ExpLength: 2,
		},
		{
			Input:     []string{"a", "a", "b", "c"},
			Expected:  []string{"a", "b", "c"},
			ExpLength: 3,
		},
		{
			Input:     []string{"a", "a", "b", "c", "c"},
			Expected:  []string{"a", "b", "c"},
			ExpLength: 3,
		},
		{
			Input:     []string{"c", "a", "a", "b", "c"},
			Expected:  []string{"a", "b", "c"},
			ExpLength: 3,
		},
		{
			Input:     []string{"a", "f", "f", "c", "g", "d", "a", "b", "e", "a", "b", "b"},
			Expected:  []string{"a", "b", "c", "d", "e", "f", "g"},
			ExpLength: 7,
		},
		{
			Input:     []string{"bee", "cee", "aee", "cee", "dee", "gee", "eff", "eee", "cee", "bee", "cee"},
			Expected:  []string{"aee", "bee", "cee", "dee", "eee", "eff", "gee"},
			ExpLength: 7,
		},
		{
			Input:     []string{"f:", "foo:", "fo:", "b:", "bar:", "ba:"},
			Expected:  []string{"b:", "ba:", "bar:", "f:", "fo:", "foo:"},
			ExpLength: 6,
		},
		{
			Input:     []string{"foobar", "foo", "bar"},
			Expected:  []string{"bar", "foo", "foobar"},
			ExpLength: 3,
		},
		{
			Input:     []string{"bar", "foo", "foobar"},
			Expected:  []string{"bar", "foo", "foobar"},
			ExpLength: 3,
		},
	}

	count.Tests(t, len(tests))
	for i, test := range tests {
		s := make([]string, len(test.Input))
		copy(s, test.Input)
		ActLength := dedup.SortAndDedupString(s)

		if test.ExpLength != ActLength {
			t.Errorf("Return integer doesn't match expected in test %d", i)
			t.Logf("  Input:    %v", test.Input)
			t.Logf("  Expected: %v", test.Expected)
			t.Logf("  Actual:   %v", s[:ActLength])
			t.Logf("  Uncropped:%v", s)
			t.Logf("  Exp Len:  %d", test.ExpLength)
			t.Logf("  Act Len:  %d", ActLength)
			continue
		}

		for j := 0; j < ActLength; j++ {
			if s[j] != test.Expected[j] {
				t.Errorf("Slice element %d doesn't match expected in test %d", j, i)
				t.Logf("  Input:    %v", test.Input)
				t.Logf("  Expected: %v", test.Expected)
				t.Logf("  Actual:   %v", s[:ActLength])
				t.Logf("  Uncropped:%v", s)
				t.Logf("  Exp Len:  %d", test.ExpLength)
				t.Logf("  Act Len:  %d", ActLength)
				t.Logf("  Exp Str: %v", test.Expected[j])
				t.Logf("  Act Str:   %v", s[j])
				continue
			}
		}
	}
}
