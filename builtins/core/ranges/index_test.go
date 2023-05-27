package ranges_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/core/index"
	_ "github.com/lmorg/murex/builtins/core/mkarray"
	_ "github.com/lmorg/murex/builtins/core/ranges"
	_ "github.com/lmorg/murex/builtins/types/string"
	_ "github.com/lmorg/murex/lang/expressions"
	"github.com/lmorg/murex/test"
)

func TestRangeLegacyByIndex(t *testing.T) {
	tests := []test.MurexTest{
		// FLAGGED

		// Include
		{
			Block:  `a: [January..December] -> @[6..10]i`,
			Stdout: "June\nJuly\nAugust\nSeptember\nOctober\n",
		},
		{
			Block:  `a: [January..December] -> @[6..]i`,
			Stdout: "June\nJuly\nAugust\nSeptember\nOctober\nNovember\nDecember\n",
		},
		{
			Block:  `a: [January..December] -> @[..6]i`,
			Stdout: "January\nFebruary\nMarch\nApril\nMay\nJune\n",
		},

		// Exclude
		{
			Block:  `a: [January..December] -> @[6..10]ie`,
			Stdout: "July\nAugust\nSeptember\n",
		},
		{
			Block:  `a: [January..December] -> @[6..]ie`,
			Stdout: "July\nAugust\nSeptember\nOctober\nNovember\nDecember\n",
		},
		{
			Block:  `a: [January..December] -> @[..6]ie`,
			Stdout: "January\nFebruary\nMarch\nApril\nMay\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestRangeIndexByIndex(t *testing.T) {
	tests := []test.MurexTest{
		// FLAGGED

		// Include
		{
			Block:  `a: [January..December] -> [6..10]i`,
			Stdout: "June\nJuly\nAugust\nSeptember\nOctober\n",
		},
		{
			Block:  `a: [January..December] -> [6..]i`,
			Stdout: "June\nJuly\nAugust\nSeptember\nOctober\nNovember\nDecember\n",
		},
		{
			Block:  `a: [January..December] -> [..6]i`,
			Stdout: "January\nFebruary\nMarch\nApril\nMay\nJune\n",
		},

		// Exclude
		{
			Block:  `a: [January..December] -> [6..10]ie`,
			Stdout: "July\nAugust\nSeptember\n",
		},
		{
			Block:  `a: [January..December] -> [6..]ie`,
			Stdout: "July\nAugust\nSeptember\nOctober\nNovember\nDecember\n",
		},
		{
			Block:  `a: [January..December] -> [..6]ie`,
			Stdout: "January\nFebruary\nMarch\nApril\nMay\n",
		},

		// DEFAULTED

		// Include
		{
			Block:  `a: [January..December] -> [6..10]`,
			Stdout: "June\nJuly\nAugust\nSeptember\nOctober\n",
		},
		{
			Block:  `a: [January..December] -> [6..]`,
			Stdout: "June\nJuly\nAugust\nSeptember\nOctober\nNovember\nDecember\n",
		},
		{
			Block:  `a: [January..December] -> [..6]`,
			Stdout: "January\nFebruary\nMarch\nApril\nMay\nJune\n",
		},

		// Exclude
		{
			Block:  `a: [January..December] -> [6..10]e`,
			Stdout: "July\nAugust\nSeptember\n",
		},
		{
			Block:  `a: [January..December] -> [6..]e`,
			Stdout: "July\nAugust\nSeptember\nOctober\nNovember\nDecember\n",
		},
		{
			Block:  `a: [January..December] -> [..6]e`,
			Stdout: "January\nFebruary\nMarch\nApril\nMay\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestRangeIndexByNegativeIndex(t *testing.T) {
	tests := []test.MurexTest{
		// FLAGGED

		// Include
		{
			Block:  `a: [January..December] -> [-6..-3]i`,
			Stdout: "June\nJuly\nAugust\n",
		},
		{
			Block:  `a: [January..December] -> [-6..]i`,
			Stdout: "June\nJuly\nAugust\nSeptember\nOctober\nNovember\nDecember\n",
		},
		/*{
			Block:  `a: [January..December] -> [..-6]i`,
			Stdout: "January\nFebruary\nMarch\nApril\nMay\nJune\n",
		},*/

		// Exclude
		{
			Block:  `a: [January..December] -> [-6..-3]ie`,
			Stdout: "July\nAugust\n",
		},
		{
			Block:  `a: [January..December] -> [-6..]ie`,
			Stdout: "July\nAugust\nSeptember\nOctober\nNovember\nDecember\n",
		},
		/*{
			Block:  `a: [January..December] -> [..-6]ie`,
			Stdout: "January\nFebruary\nMarch\nApril\nMay\nJune\n",
		},*/

		// DEFAULTED

		// Include
		{
			Block:  `a: [January..December] -> [-6..-3]`,
			Stdout: "June\nJuly\nAugust\n",
		},
		{
			Block:  `a: [January..December] -> [-6..]`,
			Stdout: "June\nJuly\nAugust\nSeptember\nOctober\nNovember\nDecember\n",
		},
		/*{
			Block:  `a: [January..December] -> [..-6]`,
			Stdout: "January\nFebruary\nMarch\nApril\nMay\nJune\n",
		},*/

		// Exclude
		{
			Block:  `a: [January..December] -> [-6..-3]e`,
			Stdout: "July\nAugust\n",
		},
		{
			Block:  `a: [January..December] -> [-6..]e`,
			Stdout: "July\nAugust\nSeptember\nOctober\nNovember\nDecember\n",
		},
		/*{
			Block:  `a: [January..December] -> [..-6]e`,
			Stdout: "January\nFebruary\nMarch\nApril\nMay\nJune\n",
		},*/
	}

	test.RunMurexTests(tests, t)
}
