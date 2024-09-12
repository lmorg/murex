package lists_test

import (
	"testing"

	"github.com/lmorg/murex/test"
)

func TestMatchPositive(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `ja: [Monday..Wednesday] -> match s`,
			Stdout: `["Tuesday","Wednesday"]`,
		},
		{
			Block:  `ja: [Monday..Wednesday] -> !match s`,
			Stdout: `["Monday"]`,
		},
		{
			Block:  `ja: [Monday..Wednesday] -> !match S`,
			Stdout: `["Monday","Tuesday","Wednesday"]`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestMatchNegative(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:   `ja: [Monday..Wednesday] -> match S`,
			Stderr:  "no data returned\n",
			ExitNum: 1,
		},
		{
			Block:   `ja: [Monday..Wednesday] -> !match day`,
			Stderr:  "no data returned\n",
			ExitNum: 1,
		},
	}

	test.RunMurexTestsRx(tests, t)
}
