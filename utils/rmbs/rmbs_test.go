package rmbs_test

import (
	"strings"
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/rmbs"
)

func TestRmBS(t *testing.T) {
	tests := []struct {
		Source   string
		Expected string
	}{
		{
			Source:   "",
			Expected: "",
		},
		{
			Source:   "--ca-bundle",
			Expected: "--ca-bundle",
		},
		{
			Source:   "--88-8-8cc8aa8--8bb8uu8nn8dd8ll8ee8",
			Expected: "ca-bundle",
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		source := strings.ReplaceAll(test.Source, "8", string([]byte{8}))

		actual := rmbs.Remove(source)

		if actual != test.Expected {
			t.Errorf("Unexpected result in test %d", i)
			t.Logf("  Source:    '%s'", test.Source)
			t.Logf("  Expected:  '%s'", test.Expected)
			t.Logf("  Actual:    '%s'", actual)
			t.Log("  exp bytes:", []byte(test.Expected))
			t.Log("  act bytes:", []byte(actual))
		}
	}
}
