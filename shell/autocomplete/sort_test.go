package autocomplete

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestSortCompletions(t *testing.T) {
	count.Tests(t, 1)

	test := []string{
		"cc", "-b:", "c", "bb:", "-a", "b:", "dd", "-bb", "d", "aa:", "a:",
	}

	expected := []string{
		"a:", "aa:", "b:", "bb:", "c", "cc", "d", "dd", "-a", "-b:", "-bb",
	}

	sortCompletions(test)

	passed := true
	for i := range test {
		passed = passed && test[i] == expected[i]
	}

	if !passed {
		t.Error("Expected splice does not match actual splice")
		t.Log("  Expected:", expected)
		t.Log("  Actual:  ", test)
	}
}

type testIsLtT struct {
	a string
	b string
	r bool
}

func TestIsLt(t *testing.T) {
	tests := []testIsLtT{
		{
			a: "foobara",
			b: "foobar:",
			r: false,
		},
		{
			a: "foobar:",
			b: "foobara",
			r: true,
		},
		{
			a: "-foobara",
			b: "foobar:",
			r: false,
		},
		{
			a: "-foobar:",
			b: "foobara",
			r: false,
		},
		{
			a: "foobara",
			b: "-foobar:",
			r: true,
		},
		{
			a: "foobar:",
			b: "-foobara",
			r: true,
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		if isLt(test.a, test.b) != test.r {
			t.Errorf("Test %d failed:", i)
			t.Logf("  a:  `%s`", test.a)
			t.Logf("  b:  `%s`", test.b)
			t.Logf("  exp: %v", test.r)
			t.Logf("  act: %v", isLt(test.a, test.b))
		}
	}
}

func TestNoColon(t *testing.T) {
	count.Tests(t, 2)

	if noColon("foobar") != "foobar" {
		t.Errorf(`noColon("foobar") != "foobar": %s`, noColon("foobar"))
	}

	if noColon("foobar:") != "foobar" {
		t.Errorf(`noColon("foobar:") != "foobar": %s`, noColon("foobar:"))
	}
}
