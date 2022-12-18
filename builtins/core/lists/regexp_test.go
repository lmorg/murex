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

func TestRegexpMatchPositive(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `ja: [Monday..Wednesday] -> regexp (m/s/)`,
			Stdout: `["Tuesday","Wednesday"]`,
		},
		{
			Block:  `ja: [Monday..Wednesday] -> regexp (m.s.)`,
			Stdout: `["Tuesday","Wednesday"]`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestRegexpMatchNegative(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:   `ja: [Monday..Wednesday] -> regexp (m/S/)`,
			Stderr:  "no data returned\n",
			ExitNum: 1,
		},
		{
			Block:   `ja: [Monday..Wednesday] -> regexp (m.S.)`,
			Stderr:  "no data returned\n",
			ExitNum: 1,
		},
	}

	test.RunMurexTestsRx(tests, t)
}

func TestRegexpMatchBangPositive(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `ja: [Monday..Wednesday] -> !regexp (m/s/)`,
			Stdout: `["Monday"]`,
		},
		{
			Block:  `ja: [Monday..Wednesday] -> !regexp (m/S/)`,
			Stdout: `["Monday","Tuesday","Wednesday"]`,
		},
		{
			Block:  `ja: [Monday..Wednesday] -> !regexp (m.s.)`,
			Stdout: `["Monday"]`,
		},
		{
			Block:  `ja: [Monday..Wednesday] -> !regexp (m.S.)`,
			Stdout: `["Monday","Tuesday","Wednesday"]`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestRegexpMatchBangNegative(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:   `ja: [Monday..Wednesday] -> !regexp (m.day.)`,
			Stderr:  "no data returned\n",
			ExitNum: 1,
		},
	}

	test.RunMurexTestsRx(tests, t)
}

func TestRegexpSubstitutePositive(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `ja: [Monday..Wednesday] -> regexp (s/s/5/)`,
			Stdout: `["Monday","Tue5day","Wedne5day"]`,
		},
		{
			Block:  `ja: [Monday..Wednesday] -> regexp (s/S/5/)`,
			Stdout: `["Monday","Tuesday","Wednesday"]`,
		},
		{
			Block:  `ja: [Monday..Wednesday] -> regexp (s.day.night.)`,
			Stdout: `["Monnight","Tuesnight","Wednesnight"]`,
		},
		{
			Block:  `ja: [Monday..Wednesday] -> regexp s/day/`,
			Stdout: `["Mon","Tues","Wednes"]`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestRegexpSubstituteNegative(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:   `ja: [Monday..Wednesday] -> regexp s/day`,
			Stderr:  "invalid regex: too few parameters",
			ExitNum: 1,
		},
	}

	test.RunMurexTestsRx(tests, t)
}

func TestRegexpSubstituteBangNegative(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:   `ja: [Monday..Wednesday] -> !regexp (s/s/5/)`,
			Stderr:  "Error in `!regexp`",
			ExitNum: 1,
		},
	}

	test.RunMurexTestsRx(tests, t)
}

func TestRegexpFindSubmatchPositive(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `ja: [Monday..Wednesday] -> regexp (f/(day)/)`,
			Stdout: `["day","day","day"]`,
		},
		{
			Block:  `ja: [Monday..Wednesday] -> regexp (f.(day).)`,
			Stdout: `["day","day","day"]`,
		},
		{
			Block:  `ja: [Monday..Wednesday] -> regexp (f/(s)/)`,
			Stdout: `["s","s"]`,
		},
		{
			Block:  `ja: [Monday..Wednesday] -> regexp (f/Tue(s)/)`,
			Stdout: `["s"]`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestRegexpFindSubmatchNegative(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:   `ja: [Monday..Wednesday] -> regexp (f/S/)`,
			Stderr:  "no data returned\n",
			ExitNum: 1,
		},
		{
			Block:   `ja: [Monday..Wednesday] -> regexp (f/s/)`,
			Stderr:  "no data returned\n",
			ExitNum: 1,
		},
	}

	test.RunMurexTestsRx(tests, t)
}

func TestRegexpFindSubmatchBang(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:   `ja: [Monday..Wednesday] -> !regexp (f/(day)/)`,
			Stderr:  "Error in `!regexp`",
			ExitNum: 1,
		},
	}

	test.RunMurexTestsRx(tests, t)
}

func TestRegexpErrors(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:   `ja: [Monday..Wednesday] -> regexp`,
			Stderr:  "Error",
			ExitNum: 1,
		},
		{
			Block:   `ja: [Monday..Wednesday] -> regexp s`,
			Stderr:  "Error",
			ExitNum: 1,
		},
		{
			Block:   `ja: [Monday..Wednesday] -> regexp f`,
			Stderr:  "Error",
			ExitNum: 1,
		},
		{
			Block:   `ja: [Monday..Wednesday] -> regexp m`,
			Stderr:  "Error",
			ExitNum: 1,
		},
		{
			Block:   `ja: [Monday..Wednesday] -> regexp b/bob/`,
			Stderr:  "Error",
			ExitNum: 1,
		},
		{
			Block:   `ja: [Monday..Wednesday] -> regexp m\\bob\\`,
			Stderr:  "Error",
			ExitNum: 1,
		},
		{
			Block:   `ja: [Monday..Wednesday] -> regexp m{bob}`,
			Stderr:  "Error",
			ExitNum: 1,
		},
	}

	test.RunMurexTestsRx(tests, t)
}
