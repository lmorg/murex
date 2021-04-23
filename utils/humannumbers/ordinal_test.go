package humannumbers_test

import (
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/humannumbers"
)

func TestOrdinal(t *testing.T) {
	tests := map[int]string{
		-12512: "-12512th",
		-12511: "-12511th",
		-12510: "-12510th",
		-1251:  "-1251st",
		-123:   "-123rd",
		-12:    "-12th",
		-11:    "-11th",
		-10:    "-10th",
		-9:     "-9th",
		-8:     "-8th",
		-7:     "-7th",
		-6:     "-6th",
		-5:     "-5th",
		-4:     "-4th",
		-3:     "-3rd",
		-2:     "-2nd",
		-1:     "-1st",

		0: "0th",

		12512: "12512th",
		12511: "12511th",
		12510: "12510th",
		1251:  "1251st",
		123:   "123rd",
		12:    "12th",
		11:    "11th",
		10:    "10th",
		9:     "9th",
		8:     "8th",
		7:     "7th",
		6:     "6th",
		5:     "5th",
		4:     "4th",
		3:     "3rd",
		2:     "2nd",
		1:     "1st",
	}

	count.Tests(t, len(tests))

	for i, exp := range tests {
		act := humannumbers.Ordinal(i)

		if exp != act {
			t.Error("Expected != Actual:")
			t.Logf("  Integer:  %d", i)
			t.Logf("  Expected: %s", exp)
			t.Logf("  Actual:   %s", act)
		}
	}
}
