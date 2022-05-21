package pathsplit_test

import (
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/pathsplit"
)

type testSplitT struct {
	Path  string
	Split []string
	Error bool
}

func TestSplit(t *testing.T) {
	tests := []testSplitT{
		{
			Path:  "",
			Error: true,
		},
		{
			Path:  "/",
			Error: true,
		},
		{
			Path:  "//",
			Split: []string{""},
		},
		{
			Path:  "##",
			Split: []string{""},
		},
		{
			Path:  ".0",
			Split: []string{"0"},
		},
		{
			Path:  ":hello:world!",
			Split: []string{"hello", "world!"},
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		actual, err := pathsplit.Split(test.Path)

		if (err != nil) != test.Error {
			t.Errorf("Error mismatch in test %d", i)
			t.Logf("Path:     '%s'", test.Path)
			t.Logf("Expected:  %s", json.LazyLogging(test.Split))
			t.Logf("Actual:    %s", json.LazyLogging(actual))
			t.Logf("exp err:   %v", test.Error)
			t.Logf("act err:   %v", err)
		}

		if json.LazyLogging(actual) != json.LazyLogging(test.Split) {
			t.Errorf("Output value doesn't match expected in test %d", i)
			t.Logf("Path:     '%s'", test.Path)
			t.Logf("Expected:  %s", json.LazyLogging(test.Split))
			t.Logf("Actual:    %s", json.LazyLogging(actual))
			t.Logf("exp err:   %v", test.Error)
			t.Logf("act err:   %v", err)
		}
	}
}
