package ranges_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/core/index"
	_ "github.com/lmorg/murex/builtins/core/mkarray"
	_ "github.com/lmorg/murex/builtins/core/ranges"
	_ "github.com/lmorg/murex/builtins/types/string"
	"github.com/lmorg/murex/test"
)

func TestRangeLegacyByString(t *testing.T) {
	tests := []test.MurexTest{
		// FLAGGED

		// Include
		{
			Block:  `a: [January..December] -> @[June..October]s`,
			Stdout: "June\nJuly\nAugust\nSeptember\nOctober\n",
		},
		{
			Block:  `a: [January..December] -> @[June..]s`,
			Stdout: "June\nJuly\nAugust\nSeptember\nOctober\nNovember\nDecember\n",
		},
		{
			Block:  `a: [January..December] -> @[..June]s`,
			Stdout: "January\nFebruary\nMarch\nApril\nMay\nJune\n",
		},

		// Exclude
		{
			Block:  `a: [January..December] -> @[June..October]se`,
			Stdout: "July\nAugust\nSeptember\n",
		},
		{
			Block:  `a: [January..December] -> @[June..]se`,
			Stdout: "July\nAugust\nSeptember\nOctober\nNovember\nDecember\n",
		},
		{
			Block:  `a: [January..December] -> @[..June]se`,
			Stdout: "January\nFebruary\nMarch\nApril\nMay\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestRangeIndexByString(t *testing.T) {
	tests := []test.MurexTest{
		// FLAGGED

		// Include
		{
			Block:  `a: [January..December] -> [June..October]s`,
			Stdout: "June\nJuly\nAugust\nSeptember\nOctober\n",
		},
		{
			Block:  `a: [January..December] -> [June..]s`,
			Stdout: "June\nJuly\nAugust\nSeptember\nOctober\nNovember\nDecember\n",
		},
		{
			Block:  `a: [January..December] -> [..June]s`,
			Stdout: "January\nFebruary\nMarch\nApril\nMay\nJune\n",
		},

		// Exclude
		{
			Block:  `a: [January..December] -> [June..October]se`,
			Stdout: "July\nAugust\nSeptember\n",
		},
		{
			Block:  `a: [January..December] -> [June..]se`,
			Stdout: "July\nAugust\nSeptember\nOctober\nNovember\nDecember\n",
		},
		{
			Block:  `a: [January..December] -> [..June]se`,
			Stdout: "January\nFebruary\nMarch\nApril\nMay\n",
		},
	}

	test.RunMurexTests(tests, t)
}
