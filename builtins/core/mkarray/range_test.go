package mkarray

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/types/string"
	"github.com/lmorg/murex/test"
)

func TestRangeMonth(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `a: [June..October]`,
			Stdout: "June\nJuly\nAugust\nSeptember\nOctober\n",
		},
		{
			Block:  `a: [June..January]`,
			Stdout: "June\nJuly\nAugust\nSeptember\nOctober\nNovember\nDecember\nJanuary\n",
		},
		{
			Block:  `a: [December..June]`,
			Stdout: "December\nJanuary\nFebruary\nMarch\nApril\nMay\nJune\n",
		},
	}

	test.RunMurexTests(tests, t)
}
