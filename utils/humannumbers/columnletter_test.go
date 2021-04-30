package humannumbers_test

import (
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/humannumbers"
)

func TestExcelColumnLetter(t *testing.T) {
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
		actual := humannumbers.ColumnLetter(i)

		if expected != actual {
			t.Error("Expected != Actual")
			t.Logf("  Integer:  %d", i)
			t.Logf("  Expected: %s", expected)
			t.Logf("  Actual:   %s", actual)
		}
	}
}
