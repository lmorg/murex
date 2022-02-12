package lists_test

import (
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/lists"
)

func TestCropPartial(t *testing.T) {
	type testCropPartialT struct {
		List    []string
		Partial string
		Output  []string
	}

	tests := []testCropPartialT{
		{
			List:    []string{"Foo", "Bar"},
			Partial: "",
			Output:  []string{"Foo", "Bar"},
		},
		{
			List:    []string{"Foo", "Bar"},
			Partial: "F",
			Output:  []string{"oo"},
		},
		{
			List:    []string{"Foo", "Bar"},
			Partial: "B",
			Output:  []string{"ar"},
		},
		{
			List:    []string{"Foo", "Bar", "Foobar"},
			Partial: "Fo",
			Output:  []string{"o", "obar"},
		},
		{
			List:    []string{"Foo", "Bar", "Foobar"},
			Partial: "Foo",
			Output:  []string{"", "bar"},
		},
		{
			List:    []string{"Foo", "Bar", "Foobar"},
			Partial: "Foob",
			Output:  []string{"ar"},
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		jExp := json.LazyLogging(test.Output)
		actual := lists.CropPartial(test.List, test.Partial)
		jAct := json.LazyLogging(actual)

		if jExp != jAct {
			t.Errorf("Test %d failed", i)
			t.Logf("  List:     %s", json.LazyLogging(test.List))
			t.Logf("  Partial:  '%s'", test.Partial)
			t.Logf("  Expected: %s", jExp)
			t.Logf("  Actual:   %s", jAct)
		}
	}
}
