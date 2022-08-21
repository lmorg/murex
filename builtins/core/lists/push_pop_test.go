package lists_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/core/mkarray"
	_ "github.com/lmorg/murex/builtins/types/json"
	_ "github.com/lmorg/murex/builtins/types/string"
	"github.com/lmorg/murex/test"
)

func TestLeft(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `ja: [Monday..Wednesday] -> left 2`,
			Stdout: `["Mo","Tu","We"]`,
		},
		{
			Block:  `a: [Monday..Wednesday] -> left 2`,
			Stdout: "Mo\nTu\nWe\n",
		},
		/////
		{
			Block:  `ja: [Monday..Wednesday] -> left -3`,
			Stdout: `["Mon","Tues","Wednes"]`,
		},
		{
			Block:  `a: [Monday..Wednesday] -> left -3`,
			Stdout: "Mon\nTues\nWednes\n",
		},
		/////
		{
			Block:  `ja: [Monday..Wednesday] -> left 10`,
			Stdout: `["Monday","Tuesday","Wednesday"]`,
		},
		{
			Block:  `a: [Monday..Wednesday] -> left 10`,
			Stdout: "Monday\nTuesday\nWednesday\n",
		},
		/////
		{
			Block:  `ja: [Monday..Wednesday] -> left 0`,
			Stdout: `["","",""]`,
		},
		{
			Block:  `a: [Monday..Wednesday] -> left 0`,
			Stdout: "\n\n\n",
		},
		/////
		{
			Block:  `ja: [Monday..Wednesday] -> left -10`,
			Stdout: `["","",""]`,
		},
		{
			Block:  `a: [Monday..Wednesday] -> left -10`,
			Stdout: "\n\n\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestRight(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `ja: [Monday..Wednesday] -> right 4`,
			Stdout: `["nday","sday","sday"]`,
		},
		{
			Block:  `a: [Monday..Wednesday] -> right 4`,
			Stdout: "nday\nsday\nsday\n",
		},
		/////
		{
			Block:  `ja: [Monday..Wednesday] -> right -3`,
			Stdout: `["day","sday","nesday"]`,
		},
		{
			Block:  `a: [Monday..Wednesday] -> right -3`,
			Stdout: "day\nsday\nnesday\n",
		},
		/////
		{
			Block:  `ja: [Monday..Wednesday] -> right 10`,
			Stdout: `["Monday","Tuesday","Wednesday"]`,
		},
		{
			Block:  `a: [Monday..Wednesday] -> right 10`,
			Stdout: "Monday\nTuesday\nWednesday\n",
		},
		/////
		{
			Block:  `ja: [Monday..Wednesday] -> right 0`,
			Stdout: `["","",""]`,
		},
		{
			Block:  `a: [Monday..Wednesday] -> right 0`,
			Stdout: "\n\n\n",
		},
		/////
		{
			Block:  `ja: [Monday..Wednesday] -> right -10`,
			Stdout: `["","",""]`,
		},
		{
			Block:  `a: [Monday..Wednesday] -> right -10`,
			Stdout: "\n\n\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestPrefix(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `ja: [Monday..Wednesday] -> prefix foobar`,
			Stdout: `["foobarMonday","foobarTuesday","foobarWednesday"]`,
		},
		{
			Block:  `a: [Monday..Wednesday] -> prefix foobar`,
			Stdout: "foobarMonday\nfoobarTuesday\nfoobarWednesday\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestSuffix(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `ja: [Monday..Wednesday] -> suffix foobar`,
			Stdout: `["Mondayfoobar","Tuesdayfoobar","Wednesdayfoobar"]`,
		},
		{
			Block:  `a: [Monday..Wednesday] -> suffix foobar`,
			Stdout: "Mondayfoobar\nTuesdayfoobar\nWednesdayfoobar\n",
		},
	}

	test.RunMurexTests(tests, t)
}
