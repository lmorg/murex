package typemgmt_test

import (
	"testing"

	_ "github.com/lmorg/murex/lang/expressions"
	"github.com/lmorg/murex/test"
)

func TestScopingSet(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `set TestScopingSet0=1
					out $TestScopingSet0`,
			Stdout: "^1\n$",
		},
		{
			Block: `function TestScopingSet1 {
						set TestScopingSet1=1
						out $TestScopingSet1
					}
					TestScopingSet1`,
			Stdout: "^1\n$",
		},
		{
			Block: `function TestScopingSet2 {
						out $TestScopingSet2
					}
					set TestScopingSet2=1
					TestScopingSet2`,
			Stderr:  "variable 'TestScopingSet2' does not exist\n",
			ExitNum: 1,
		},
		{
			Block: `function TestScopingSet3 {
						set: TestScopingSet3=2
					}
					set: TestScopingSet3=1
					TestScopingSet3
					out: $TestScopingSet3`,
			Stdout: "^1\n$",
		},
		{
			Block: `function TestScopingSet4.0 {
						set TestScopingSet4=2
						$TestScopingSet4
					}
					function TestScopingSet4.1 {
						set TestScopingSet4=3
						$TestScopingSet4
					}
					set TestScopingSet4=1
					TestScopingSet4.0
					TestScopingSet4.1
					out $TestScopingSet4`,
			Stdout: "^231\n$",
		},
		{
			Block: `function TestScopingSet5.0 {
						set TestScopingSet5=2
						$TestScopingSet5
						TestScopingSet5.1
					}
					function TestScopingSet5.1 {
						set TestScopingSet5=3
						$TestScopingSet5
					}
					set TestScopingSet5=1
					TestScopingSet5.0
					out $TestScopingSet5`,
			Stdout: "^231\n$",
		},
		{
			Block: `function TestScopingSet6.0 {
						set TestScopingSet6=2
						$TestScopingSet6
						TestScopingSet6.1
					}
					function TestScopingSet6.1 {
						set TestScopingSet6=3
						$TestScopingSet6
					}
					TestScopingSet6.0
					out $TestScopingSet6`,
			Stdout:  "^23$",
			Stderr:  "variable 'TestScopingSet6' does not exist\n",
			ExitNum: 1,
		},
	}

	test.RunMurexTestsRx(tests, t)
}

func TestScopingGlobal(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `global TestScopingGlobal0=1
					out $TestScopingGlobal0`,
			Stdout: "1\n",
		},
		{
			Block: `function TestScopingGlobal1.0 {
						global TestScopingGlobal=1
						out $TestScopingGlobal
					}
					TestScopingGlobal1.0`,
			Stdout: "1\n",
		},
		{
			Block: `function TestScopingGlobal2.0 {
						out $TestScopingGlobal
					}
					global TestScopingGlobal=1
					TestScopingGlobal2.0`,
			Stdout: "1\n",
		},
		{
			Block: `function TestScopingGlobal3.0 {
						global TestScopingGlobal=2
					}
					global TestScopingGlobal=1
					TestScopingGlobal3.0
					out $TestScopingGlobal`,
			Stdout: "2\n",
		},
		{
			Block: `function TestScopingGlobal4.0 {
						global TestScopingGlobal=2
						$TestScopingGlobal
					}
					function TestScopingGlobal4.1 {
						global TestScopingGlobal=3
						$TestScopingGlobal
					}
					global TestScopingGlobal=1
					TestScopingGlobal4.0
					TestScopingGlobal4.1
					out $TestScopingGlobal`,
			Stdout: "233\n",
		},
		{
			Block: `function TestScopingGlobal5.0 {
						global TestScopingGlobal=2
						$TestScopingGlobal
						TestScopingGlobal5.1
					}
					function TestScopingGlobal5.1 {
						global TestScopingGlobal=3
						$TestScopingGlobal
					}
					global TestScopingGlobal=1
					TestScopingGlobal5.0
					out $TestScopingGlobal`,
			Stdout: "233\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestScopingMixed(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `function TestScopingGlobal0.0 {
						set TestScopingGlobal=2
					}
					global TestScopingGlobal=1
					TestScopingGlobal0.0
					out $TestScopingGlobal`,
			Stdout: "1\n",
		},
		{
			Block: `function TestScopingGlobal1.0 {
						set TestScopingGlobal=2
						$TestScopingGlobal
					}
					function TestScopingGlobal1.1 {
						set TestScopingGlobal=3
						$TestScopingGlobal
					}
					global TestScopingGlobal=1
					TestScopingGlobal1.0
					TestScopingGlobal1.1
					out $TestScopingGlobal`,
			Stdout: "231\n",
		},
		{
			Block: `function TestScopingGlobal2.0 {
						set TestScopingGlobal=2
						$TestScopingGlobal
						TestScopingGlobal2.1
					}
					function TestScopingGlobal2.1 {
						set TestScopingGlobal=3
						$TestScopingGlobal
					}
					global TestScopingGlobal=1
					TestScopingGlobal2.0
					out $TestScopingGlobal`,
			Stdout: "231\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestScopingNonFuncBlock(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `if { true } then {
						set TestScopingNonFuncBlock0=1
						$TestScopingNonFuncBlock0
					}
					set TestScopingNonFuncBlock0=2
					out $TestScopingNonFuncBlock0`,
			Stdout: "12\n",
		},
		{
			Block: `set TestScopingNonFuncBlock1=1
					$TestScopingNonFuncBlock1
					if { true } then {
						set TestScopingNonFuncBlock1=2
						$TestScopingNonFuncBlock1
					}
					set TestScopingNonFuncBlock1=3
					out $TestScopingNonFuncBlock1`,
			Stdout: "123\n",
		},
		{
			Block: `set TestScopingNonFuncBlock2=1
					$TestScopingNonFuncBlock2
					if { true } then {
						out $TestScopingNonFuncBlock2
					}`,
			Stdout: "11\n",
		},
		{
			Block: `if { true } then {
						set TestScopingNonFuncBlock3=1
						$TestScopingNonFuncBlock3
					}
					out $TestScopingNonFuncBlock3`,
			Stdout: "11\n",
		},
	}

	test.RunMurexTests(tests, t)
}
