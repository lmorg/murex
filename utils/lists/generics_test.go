package lists_test

import (
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/lists"
)

func TestGenericToString(t *testing.T) {
	tests := []struct {
		Source   []any
		Expected []string
		Error    bool
	}{
		{
			Source:   []interface{}{1, 2, 3},
			Expected: []string{"1", "2", "3"},
			Error:    false,
		},
		{
			Source:   []interface{}{"1", "2", "3"},
			Expected: []string{"1", "2", "3"},
			Error:    false,
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		actual, err := lists.GenericToString(test.Source)
		expJson := json.LazyLoggingPretty(test.Expected)
		actJson := json.LazyLoggingPretty(actual)

		if expJson != actJson || (err != nil) != test.Error {
			t.Errorf("Conversion failed in test %d", i)
			t.Logf("  Source:   %s", json.LazyLoggingPretty(test.Source))
			t.Logf("  Expected: %s", expJson)
			t.Logf("  Actual:   %s", actJson)
			t.Logf("  exp err:  %v", test.Error)
			t.Logf("  act err:  %v", err)
		}
	}
}
