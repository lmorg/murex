package structs_test

import (
	"testing"

	"github.com/lmorg/murex/test"
)

func TestReturn(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `out 0; return; out 2`,
			Stdout: "0\n",
		},
		{
			Block:  `out 1; return 0; out 2`,
			Stdout: "1\n",
		},
		{
			Block:   `out 2; return 3; out 2`,
			Stdout:  "2\n",
			ExitNum: 3,
		},
		{
			Block: `
				function TestReturn3 {
					out 3
					return 4
					out 5
				}
				TestReturn3`,
			Stdout:  "3\n",
			ExitNum: 4,
		},
		{
			Block: `
				function TestReturn4 {
					out 4
					return 5
					out 6
				}
				TestReturn4
				exitnum`,
			Stdout:  "4\n5\n",
			ExitNum: 0,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestBreak(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				function TestBreak0 {
					out 3
					break TestBreak0
					out 5
				}
				TestBreak0`,
			Stdout:  "3\n",
			ExitNum: 0,
		},
		{
			Block: `
				function TestBreak1 {
					out 4
					break TestBreak1
					out 6
				}
				TestBreak1
				exitnum`,
			Stdout:  "4\n0\n",
			ExitNum: 0,
		},
		{
			Block: `
				function TestBreak2 {
					%[1..5] -> foreach i {
						out $i
						if { $i == 3 } then {
							break foreach
						}
						out "foobar"
					}
				}
				TestBreak2`,
			Stdout:  "1\nfoobar\n2\nfoobar\n3\n",
			ExitNum: 0,
		},
		{
			Block: `
				function TestBreak3 {
					%[1..5] -> foreach i {
						out $i
						if { $i == 3 } then {
							break TestBreak3
						}
						out "foobar"
					}
				}
				TestBreak3`,
			Stdout:  "1\nfoobar\n2\nfoobar\n3\n",
			ExitNum: 0,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestContinue(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				function TestContinue0 {
					out 3
					continue TestContinue0
					out 5
				}
				TestContinue0`,
			Stdout:  "3\n5\n",
			ExitNum: 0,
		},
		{
			Block: `
				function TestContinue1 {
					out 4
					continue TestContinue1
					out 6
				}
				TestContinue1
				exitnum`,
			Stdout:  "4\n6\n0\n",
			ExitNum: 0,
		},
		{
			Block: `
				function TestContinue2 {
					%[1..5] -> foreach i {
						out $i
						if { $i == 3 } then {
							continue foreach
						}
						out "foobar"
					}
				}
				TestContinue2`,
			Stdout:  "1\nfoobar\n2\nfoobar\n3\n4\nfoobar\n5\nfoobar\n",
			ExitNum: 0,
		},
		{
			Block: `
				function TestContinue3 {
					%[1..5] -> foreach i {
						out $i
						if { $i == 3 } then {
							continue TestContinue3
						}
						out "foobar"
					}
				}
				TestContinue3`,
			Stdout:  "1\nfoobar\n2\nfoobar\n3\n",
			ExitNum: 0,
		},
	}

	test.RunMurexTests(tests, t)
}
