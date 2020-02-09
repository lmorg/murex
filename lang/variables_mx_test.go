package lang_test

import (
	"testing"

	"github.com/lmorg/murex/test"
)

// https://github.com/lmorg/murex/issues/138
func TestVariableMxLocal(t *testing.T) {
	mxt := test.MurexTest{
		Block: `
			function TestVariableMxLocal2 {
				set TestVariableMxLocal=redefined
			}
			function TestVariableMxLocal1 {
				set TestVariableMxLocal=original
				TestVariableMxLocal2
				$TestVariableMxLocal
			}
			TestVariableMxLocal1`,

		Stdout: `original`,
	}

	test.RunMurexTests([]test.MurexTest{mxt}, t)
}

func TestVariableMxGlobal(t *testing.T) {
	mxt := test.MurexTest{
		Block: `
			function TestVariableMxGlobal2 {
				global TestVariableMxGlobal=redefined
			}
			function TestVariableMxGlobal1 {
				global TestVariableMxGlobal=original
				TestVariableMxGlobal2
				$TestVariableMxGlobal
			}
			TestVariableMxGlobal1`,

		Stdout: `redefined`,
	}

	test.RunMurexTests([]test.MurexTest{mxt}, t)
}

func TestVariableMxLocalGlobal(t *testing.T) {
	mxt := test.MurexTest{
		Block: `
			function TestVariableMxLocalGlobal2 {
				global t=redefined
			}
			function TestVariableMxLocalGlobal1 {
				set TestVariableMxLocalGlobal=original
				TestVariableMxLocalGlobal2
				$TestVariableMxLocalGlobal
			}
			TestVariableMxLocalGlobal1`,

		Stdout: `original`,
	}

	test.RunMurexTests([]test.MurexTest{mxt}, t)
}

/*func TestVariableMxGlobalLocal(t *testing.T) {
	mxt := test.MurexTest{
		Block: `
			function test2 {
				set t=redefined
			}
			function test1 {
				global t=original
				test2
				$t
			}
			test1`,

		Stdout: `original`,
	}

	test.RunMurexTests([]test.MurexTest{mxt}, t)
}*/

func TestVariableMxForLoop(t *testing.T) {
	mxt := test.MurexTest{
		Block: `
			function TestVariableMxForLoop2 {
				for (TestVariableMxForLoop=0; TestVariableMxForLoop<5; TestVariableMxForLoop++) {
					$TestVariableMxForLoop
				}
				echo $TestVariableMxForLoop
			}
			function TestVariableMxForLoop1 {
				TestVariableMxForLoop2
				echo $TestVariableMxForLoop
			}
			TestVariableMxForLoop1`,

		Stdout: "01234\n\n",
	}

	test.RunMurexTests([]test.MurexTest{mxt}, t)
}

func TestVariableMxForLoopGlobal(t *testing.T) {
	mxt := test.MurexTest{
		Block: `
			function TestVariableMxForLoopGlobal2 {
				for (TestVariableMxForLoopGlobal=0; TestVariableMxForLoopGlobal<5; TestVariableMxForLoopGlobal++) {
					$TestVariableMxForLoopGlobal
				}
				echo $TestVariableMxForLoopGlobal
			}
			function TestVariableMxForLoopGlobal1 {
				global TestVariableMxForLoopGlobal=-1
				TestVariableMxForLoopGlobal2
				echo $TestVariableMxForLoopGlobal
			}
			TestVariableMxForLoopGlobal1`,

		Stdout: "012345\n5\n",
	}

	test.RunMurexTests([]test.MurexTest{mxt}, t)
}

func TestVariableMxForLoopLocal(t *testing.T) {
	mxt := test.MurexTest{
		Block: `
			function TestVariableMxForLoopLocal2 {
				for (TestVariableMxForLoopLocal=0; TestVariableMxForLoopLocal<5; TestVariableMxForLoopLocal++) {
					$TestVariableMxForLoopLocal
				}
				echo $TestVariableMxForLoopLocal
			}
			function TestVariableMxForLoopLocal1 {
				set TestVariableMxForLoopLocal=-1
				TestVariableMxForLoopLocal2
				echo $TestVariableMxForLoopLocal
			}
			TestVariableMxForLoopLocal1`,

		Stdout: "01234\n-1\n",
	}

	test.RunMurexTests([]test.MurexTest{mxt}, t)
}
