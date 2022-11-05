package ranges

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/core/mkarray"
	_ "github.com/lmorg/murex/builtins/types/string"
	"github.com/lmorg/murex/test"
)

func TestRangeLegacyByIndex(t *testing.T) {
	tests := []test.MurexTest{
		// FLAGGED

		// Include
		{
			Block:  `a: [January..December] -> @[5..9]n`,
			Stdout: "June\nJuly\nAugust\nSeptember\nOctober\n",
		},
		{
			Block:  `a: [January..December] -> @[5..]n`,
			Stdout: "June\nJuly\nAugust\nSeptember\nOctober\nNovember\nDecember\n",
		},
		{
			Block:  `a: [January..December] -> @[..5]n`,
			Stdout: "January\nFebruary\nMarch\nApril\nMay\nJune\n",
		},

		// Exclude
		{
			Block:  `a: [January..December] -> @[5..9]ne`,
			Stdout: "July\nAugust\nSeptember\n",
		},
		{
			Block:  `a: [January..December] -> @[5..]ne`,
			Stdout: "July\nAugust\nSeptember\nOctober\nNovember\nDecember\n",
		},
		{
			Block:  `a: [January..December] -> @[..5]ne`,
			Stdout: "January\nFebruary\nMarch\nApril\nMay\n",
		},

		// DEFAULTED

		// Include
		{
			Block:  `a: [January..December] -> @[5..9]`,
			Stdout: "June\nJuly\nAugust\nSeptember\nOctober\n",
		},
		{
			Block:  `a: [January..December] -> @[5..]`,
			Stdout: "June\nJuly\nAugust\nSeptember\nOctober\nNovember\nDecember\n",
		},
		{
			Block:  `a: [January..December] -> @[..5]`,
			Stdout: "January\nFebruary\nMarch\nApril\nMay\nJune\n",
		},

		// Exclude
		{
			Block:  `a: [January..December] -> @[5..9]e`,
			Stdout: "July\nAugust\nSeptember\n",
		},
		{
			Block:  `a: [January..December] -> @[5..]e`,
			Stdout: "July\nAugust\nSeptember\nOctober\nNovember\nDecember\n",
		},
		{
			Block:  `a: [January..December] -> @[..5]e`,
			Stdout: "January\nFebruary\nMarch\nApril\nMay\n",
		},
	}

	test.RunMurexTests(tests, t)
}
