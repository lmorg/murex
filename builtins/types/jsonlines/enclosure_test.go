package jsonlines

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

type noTests struct {
	slice    string
	expected bool
}

func TestNoQuote(t *testing.T) {
	tests := []noTests{
		// empty
		{
			slice:    `foobar`,
			expected: true,
		},
		// single quote
		{
			slice:    `'foobar`,
			expected: true,
		},
		{
			slice:    `foobar'`,
			expected: true,
		},
		{
			slice:    `'foobar'`,
			expected: true,
		},
		// double quote
		{
			slice:    `"foobar`,
			expected: true,
		},
		{
			slice:    `foobar"`,
			expected: true,
		},
		{
			slice:    `"foobar"`,
			expected: false,
		},
		// square bracket
		{
			slice:    `[foobar`,
			expected: true,
		},
		{
			slice:    `foobar]`,
			expected: true,
		},
		{
			slice:    `[foobar]`,
			expected: true,
		},
		// curley bracket
		{
			slice:    `{foobar`,
			expected: true,
		},
		{
			slice:    `foobar}`,
			expected: true,
		},
		{
			slice:    `{foobar}`,
			expected: true,
		},
		// round bracket
		{
			slice:    `(foobar`,
			expected: true,
		},
		{
			slice:    `foobar)`,
			expected: true,
		},
		{
			slice:    `(foobar)`,
			expected: true,
		},
	}

	count.Tests(t, len(tests))
	for i := range tests {
		actual := noQuote([]byte(tests[i].slice))
		if actual != tests[i].expected {
			t.Errorf("%s doesn't return expected", t.Name())
			t.Logf("  slice:    %s", tests[i].slice)
			t.Logf("  expected: %v", tests[i].expected)
			t.Logf("  actual:   %v", actual)
		}

	}

}

func TestNoSquare(t *testing.T) {
	tests := []noTests{
		// empty
		{
			slice:    `foobar`,
			expected: true,
		},
		// single quote
		{
			slice:    `'foobar`,
			expected: true,
		},
		{
			slice:    `foobar'`,
			expected: true,
		},
		{
			slice:    `'foobar'`,
			expected: true,
		},
		// double quote
		{
			slice:    `"foobar`,
			expected: true,
		},
		{
			slice:    `foobar"`,
			expected: true,
		},
		{
			slice:    `"foobar"`,
			expected: true,
		},
		// square bracket
		{
			slice:    `[foobar`,
			expected: true,
		},
		{
			slice:    `foobar]`,
			expected: true,
		},
		{
			slice:    `[foobar]`,
			expected: false,
		},
		// curley bracket
		{
			slice:    `{foobar`,
			expected: true,
		},
		{
			slice:    `foobar}`,
			expected: true,
		},
		{
			slice:    `{foobar}`,
			expected: true,
		},
		// round bracket
		{
			slice:    `(foobar`,
			expected: true,
		},
		{
			slice:    `foobar)`,
			expected: true,
		},
		{
			slice:    `(foobar)`,
			expected: true,
		},
	}

	count.Tests(t, len(tests))
	for i := range tests {
		actual := noSquare([]byte(tests[i].slice))
		if actual != tests[i].expected {
			t.Errorf("%s doesn't return expected", t.Name())
			t.Logf("  slice:    %s", tests[i].slice)
			t.Logf("  expected: %v", tests[i].expected)
			t.Logf("  actual:   %v", actual)
		}

	}

}

func TestNoCurly(t *testing.T) {
	tests := []noTests{
		// empty
		{
			slice:    `foobar`,
			expected: true,
		},
		// single quote
		{
			slice:    `'foobar`,
			expected: true,
		},
		{
			slice:    `foobar'`,
			expected: true,
		},
		{
			slice:    `'foobar'`,
			expected: true,
		},
		// double quote
		{
			slice:    `"foobar`,
			expected: true,
		},
		{
			slice:    `foobar"`,
			expected: true,
		},
		{
			slice:    `"foobar"`,
			expected: true,
		},
		// square bracket
		{
			slice:    `[foobar`,
			expected: true,
		},
		{
			slice:    `foobar]`,
			expected: true,
		},
		{
			slice:    `[foobar]`,
			expected: true,
		},
		// curley bracket
		{
			slice:    `{foobar`,
			expected: true,
		},
		{
			slice:    `foobar}`,
			expected: true,
		},
		{
			slice:    `{foobar}`,
			expected: false,
		},
		// round bracket
		{
			slice:    `(foobar`,
			expected: true,
		},
		{
			slice:    `foobar)`,
			expected: true,
		},
		{
			slice:    `(foobar)`,
			expected: true,
		},
	}

	count.Tests(t, len(tests))
	for i := range tests {
		actual := noCurly([]byte(tests[i].slice))
		if actual != tests[i].expected {
			t.Errorf("%s doesn't return expected", t.Name())
			t.Logf("  slice:    %s", tests[i].slice)
			t.Logf("  expected: %v", tests[i].expected)
			t.Logf("  actual:   %v", actual)
		}

	}

}
