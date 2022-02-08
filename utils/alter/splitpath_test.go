package alter

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

type test struct {
	path     string
	expected []string
}

func TestSplitPath(t *testing.T) {
	tests := []test{
		{
			path:     "/",
			expected: []string{""},
		},

		{
			path:     "/a/b/c",
			expected: []string{"a", "b", "c"},
		},
		{
			path:     "/aaa/bbb/ccc",
			expected: []string{"aaa", "bbb", "ccc"},
		},
		{
			path:     "/1/2/3",
			expected: []string{"1", "2", "3"},
		},

		{
			path:     ".a.b.c",
			expected: []string{"a", "b", "c"},
		},
		{
			path:     ".aaa.bbb.ccc",
			expected: []string{"aaa", "bbb", "ccc"},
		},
		{
			path:     ".1.2.3",
			expected: []string{"1", "2", "3"},
		},

		{
			path:     "-a-b-c",
			expected: []string{"a", "b", "c"},
		},
		{
			path:     "-aaa-bbb-ccc",
			expected: []string{"aaa", "bbb", "ccc"},
		},
		{
			path:     "-1-2-3",
			expected: []string{"1", "2", "3"},
		},

		{
			path:     "1a1b1c",
			expected: []string{"a", "b", "c"},
		},
		{
			path:     "1aaa1bbb1ccc",
			expected: []string{"aaa", "bbb", "ccc"},
		},
		{
			path:     "111213",
			expected: []string{"", "", "2", "3"},
		},
	}

	count.Tests(t, len(tests))

	for i := range tests {
		split, err := SplitPath(tests[i].path)
		if err != nil {
			t.Error("SplitPath raised an error")
			t.Logf("  index:    %d", i)
			t.Logf("  path:     %s", tests[i].path)
			t.Logf("  expected: %v (%d)", tests[i].expected, len(tests[i].expected))
			t.Logf("  actual:   %v (%d)", split, len(split))
		}

		if len(split) != len(tests[i].expected) {
			t.Error("SplitPath returned a different length slice to expected")
			t.Logf("  index:    %d", i)
			t.Logf("  path:     %s", tests[i].path)
			t.Logf("  expected: %v (%d)", tests[i].expected, len(tests[i].expected))
			t.Logf("  actual:   %v (%d)", split, len(split))
		}

		for j := range split {
			if split[j] != tests[i].expected[j] {
				t.Error("SplitPath returned a different slice to expected")
				t.Logf("  index:      %d", i)
				t.Logf("  path:       %s", tests[i].path)
				t.Logf("  expected:   %v (%d)", tests[i].expected, len(tests[i].expected))
				t.Logf("  actual:     %v (%d)", split, len(split))
				t.Logf(`  diff value: "%s" != "%s"`, split[j], tests[i].expected[j])
				t.Logf(`  expected:   "%s"`, tests[i].expected[j])
				t.Logf(`  actual:     "%s"`, split[j])
				t.Logf("  diff index: %d", j)
			}
		}

	}
}
