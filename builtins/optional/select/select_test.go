package sqlselect

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

type Str2IfaceT struct {
	Input  []string
	Max    int
	Output []string
}

func TestStringToInterfaceTrim(t *testing.T) {
	tests := []Str2IfaceT{
		{
			Input:  []string{"a", "b", "c", "d", "e"},
			Max:    0,
			Output: []string{},
		},
		{
			Input:  []string{"a", "b", "c", "d", "e"},
			Max:    1,
			Output: []string{"a"},
		},
		{
			Input:  []string{"a", "b", "c", "d", "e"},
			Max:    2,
			Output: []string{"a", "b"},
		},
		{
			Input:  []string{"a", "b", "c", "d", "e"},
			Max:    3,
			Output: []string{"a", "b", "c"},
		},
		{
			Input:  []string{"a", "b", "c", "d", "e"},
			Max:    4,
			Output: []string{"a", "b", "c", "d"},
		},
		{
			Input:  []string{"a", "b", "c", "d", "e"},
			Max:    5,
			Output: []string{"a", "b", "c", "d", "e"},
		},
		{
			Input:  []string{"a", "b", "c", "d", "e"},
			Max:    6,
			Output: []string{"a", "b", "c", "d", "e", ""},
		},
		{
			Input:  []string{"a", "b", "c", "d", "e"},
			Max:    7,
			Output: []string{"a", "b", "c", "d", "e", "", ""},
		},
		{
			Input:  []string{"a", "b", "c", "d", "e"},
			Max:    8,
			Output: []string{"a", "b", "c", "d", "e", "", "", ""},
		},
		/////
		{
			Input:  []string{},
			Max:    5,
			Output: []string{"", "", "", "", ""},
		},
		{
			Input:  []string{"a"},
			Max:    5,
			Output: []string{"a", "", "", "", ""},
		},
		{
			Input:  []string{"a", "b"},
			Max:    5,
			Output: []string{"a", "b", "", "", ""},
		},
		{
			Input:  []string{"a", "b", "c"},
			Max:    5,
			Output: []string{"a", "b", "c", "", ""},
		},
		{
			Input:  []string{"a", "b", "c", "d"},
			Max:    5,
			Output: []string{"a", "b", "c", "d", ""},
		},
		{
			Input:  []string{"a", "b", "c", "d", "e"},
			Max:    5,
			Output: []string{"a", "b", "c", "d", "e"},
		},
		{
			Input:  []string{"a", "b", "c", "d", "e", "f"},
			Max:    5,
			Output: []string{"a", "b", "c", "d", "e"},
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		actual := stringToInterfaceTrim(test.Input, test.Max)

		if len(test.Output) != len(actual) {
			t.Errorf("Length mismatch in test %d", i)
			t.Logf("  Input:    %v", test.Input)
			t.Logf("  Max:      %d", test.Max)
			t.Logf("  Expected: %v", test.Output)
			t.Logf("  Actual:   %v", actual)
		}

		for j := range test.Output {
			if test.Output[j] != actual[j].(string) {
				t.Errorf("Value mismatch in test %d[%d]", i, j)
				t.Logf("  Input:    %v", test.Input)
				t.Logf("  Max:      %d", test.Max)
				t.Logf("  Expected: %v", test.Output)
				t.Logf("  Actual:   %v", actual)
				t.Logf("  Expected: '%s'", test.Output[j])
				t.Logf("  Actual:   '%s'", actual[j].(string))
			}
		}
	}
}

func TestIToColumnLetter(t *testing.T) {
	tests := map[int]string{
		0:  "A",
		1:  "B",
		2:  "C",
		3:  "D",
		4:  "E",
		5:  "F",
		6:  "G",
		7:  "H",
		8:  "I",
		9:  "J",
		10: "K",
		11: "L",
		12: "M",
		// ...
		20: "U",
		21: "V",
		22: "W",
		23: "X",
		24: "Y",
		25: "Z",
		26: "AA",
		27: "AB",
		28: "AC",
		29: "AD",
		30: "AE",
		// ...
		595: "VX",
		596: "VY",
		597: "VZ",
		598: "WA",
		599: "WB",
		600: "WC",
		// ...
		700: "ZY",
		701: "ZZ",
		702: "AAA",
		703: "AAB",
		704: "AAC",
		705: "AAD",
		// ...
		725: "AAX",
		726: "AAY",
		727: "AAZ",
		728: "ABA",
		729: "ABB",
		730: "ABC",
	}

	count.Tests(t, len(tests))

	for i, expected := range tests {
		actual := iToColumnLetter(i)

		if expected != actual {
			t.Error("Expected != Actual")
			t.Logf("  Integer:  %d", i)
			t.Logf("  Expected: %s", expected)
			t.Logf("  Actual:   %s", actual)
		}
	}
}
