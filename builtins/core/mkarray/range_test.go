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
		{
			Block:   `a: [..June]`,
			Stdout:  "",
			Stderr:  "Error in `a` ( 1,1): Unable to auto-detect range in `..June`\n",
			ExitNum: 1,
		},
		{
			Block:   `a: [June..]`,
			Stdout:  "",
			Stderr:  "Error in `a` ( 1,1): Unable to auto-detect range in `June..`\n",
			ExitNum: 1,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestRange(t *testing.T) {
	tests := []test.MurexTest{
		// single range
		{
			Block:  `a: [00..03]`,
			Stdout: "00\n01\n02\n03\n",
		},
		{
			Block:  `a: [01,03]`,
			Stdout: "01\n03\n",
		},
		// multiple ranges
		{
			Block:  `a: [01..02][01..02]`,
			Stdout: "0101\n0102\n0201\n0202\n",
		},
		{
			Block:  `a: [01,03][01..02]`,
			Stdout: "0101\n0102\n0301\n0302\n",
		},
		{
			Block:  `a: [01..02][01,03]`,
			Stdout: "0101\n0103\n0201\n0203\n",
		},
		{
			Block:  `a: [01,03][01,03]`,
			Stdout: "0101\n0103\n0301\n0303\n",
		},
		// multiple ranges with non-ranged data
		{
			Block:  `a: .[01..02].[01..02].`,
			Stdout: ".01.01.\n.01.02.\n.02.01.\n.02.02.\n",
		},
		{
			Block:  `a: .[01,03].[01..02].`,
			Stdout: ".01.01.\n.01.02.\n.03.01.\n.03.02.\n",
		},
		{
			Block:  `a: .[01..02].[01,03].`,
			Stdout: ".01.01.\n.01.03.\n.02.01.\n.02.03.\n",
		},
		{
			Block:  `a: .[01,03].[01,03].`,
			Stdout: ".01.01.\n.01.03.\n.03.01.\n.03.03.\n",
		},
	}

	test.RunMurexTests(tests, t)
}
