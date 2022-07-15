package term

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestAppendBytes(t *testing.T) {
	tests := []struct {
		Slice    string
		Data     string
		Expected string
	}{
		{
			Slice:    "",
			Data:     "",
			Expected: "",
		},
		{
			Slice:    "",
			Data:     "bar",
			Expected: "bar",
		},
		{
			Slice:    "foo",
			Data:     "",
			Expected: "foo",
		},
		{
			Slice:    "foo",
			Data:     "bar",
			Expected: "foobar",
		},
		{
			Slice:    "foo",
			Data:     "barbarbarbarbarbarbarbarbarbarbarbarbarbarbarbarbarbarbar",
			Expected: "foobarbarbarbarbarbarbarbarbarbarbarbarbarbarbarbarbarbarbar",
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		b := appendBytes([]byte(test.Slice), []byte(test.Data)...)
		actual := string(b)
		if actual != test.Expected {
			t.Errorf("Actual does not match expected in test %d", i)
			t.Logf("  Slice:    '%s'", test.Slice)
			t.Logf("  Data:     '%s'", test.Data)
			t.Logf("  Expected: '%s'", test.Expected)
			t.Logf("  Actual:   '%s'", actual)
		}
	}
}
