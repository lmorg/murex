package structs_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/core/io"
	_ "github.com/lmorg/murex/builtins/core/structs"
	_ "github.com/lmorg/murex/builtins/core/typemgmt"
	"github.com/lmorg/murex/test"
)

func TestSwitchCaseBlock1(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				switch {
					case { true } then { out: 1 }
					case { false } then { out: 2 }
					case { false } then { out: 3 }
				}`,
			Stdout: "1\n",
		},
		{
			Block: `
				switch {
					case { true } then { out: 1 }
					case { true } then { out: 2 }
					case { false } then { out: 3 }
				}`,
			Stdout: "1\n",
		},
		{
			Block: `
				switch {
					case { true } then { out: 1 }
					case { true } then { out: 2 }
					case { true } then { out: 3 }
				}`,
			Stdout: "1\n",
		},
		{
			Block: `
				switch {
					case { true } then { out: 1 }
					case { false } then { out: 2 }
					case { false } then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "1\n",
		},
		{
			Block: `
				switch {
					case { true } then { out: 1 }
					case { true } then { out: 2 }
					case { false } then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "1\n",
		},
		{
			Block: `
				switch {
					case { true } then { out: 1 }
					case { true } then { out: 2 }
					case { true } then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "1\n",
		},
		/////
		{
			Block: `
				switch {
					case { true }  { out: 1 }
					case { false }  { out: 2 }
					case { false }  { out: 3 }
				}`,
			Stdout: "1\n",
		},
		{
			Block: `
				switch {
					case { true }  { out: 1 }
					case { true }  { out: 2 }
					case { false }  { out: 3 }
				}`,
			Stdout: "1\n",
		},
		{
			Block: `
				switch {
					case { true }  { out: 1 }
					case { true }  { out: 2 }
					case { true }  { out: 3 }
				}`,
			Stdout: "1\n",
		},
		{
			Block: `
				switch {
					case { true }  { out: 1 }
					case { false }  { out: 2 }
					case { false }  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "1\n",
		},
		{
			Block: `
				switch {
					case { true }  { out: 1 }
					case { true }  { out: 2 }
					case { false }  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "1\n",
		},
		{
			Block: `
				switch {
					case { true }  { out: 1 }
					case { true }  { out: 2 }
					case { true }  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "1\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestSwitchCaseBlock2(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				switch {
					case { false } then { out: 1 }
					case { true } then { out: 2 }
					case { false } then { out: 3 }
				}`,
			Stdout: "2\n",
		},
		{
			Block: `
				switch {
					case { false } then { out: 1 }
					case { true } then { out: 2 }
					case { true } then { out: 3 }
				}`,
			Stdout: "2\n",
		},
		{
			Block: `
				switch {
					case { false } then { out: 1 }
					case { true } then { out: 2 }
					case { false } then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "2\n",
		},
		{
			Block: `
				switch {
					case { false } then { out: 1 }
					case { true } then { out: 2 }
					case { true } then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "2\n",
		},
		/////
		{
			Block: `
				switch {
					case { false }  { out: 1 }
					case { true }  { out: 2 }
					case { false }  { out: 3 }
				}`,
			Stdout: "2\n",
		},
		{
			Block: `
				switch {
					case { false }  { out: 1 }
					case { true }  { out: 2 }
					case { true }  { out: 3 }
				}`,
			Stdout: "2\n",
		},
		{
			Block: `
				switch {
					case { false }  { out: 1 }
					case { true }  { out: 2 }
					case { false }  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "2\n",
		},
		{
			Block: `
				switch {
					case { false }  { out: 1 }
					case { true }  { out: 2 }
					case { true }  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "2\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestSwitchCaseBlock3(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				switch {
					case { false } then { out: 1 }
					case { false } then { out: 2 }
					case { true } then { out: 3 }
				}`,
			Stdout: "3\n",
		},
		{
			Block: `
				switch {
					case { false } then { out: 1 }
					case { false } then { out: 2 }
					case { true } then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "3\n",
		},
		/////
		{
			Block: `
				switch {
					case { false }  { out: 1 }
					case { false }  { out: 2 }
					case { true }  { out: 3 }
				}`,
			Stdout: "3\n",
		},
		{
			Block: `
				switch {
					case { false }  { out: 1 }
					case { false }  { out: 2 }
					case { true }  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "3\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestSwitchCaseBlock4(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				switch {
					case { false } then { out: 1 }
					case { false } then { out: 2 }
					case { false } then { out: 3 }
				}`,
			Stdout:  "",
			ExitNum: 1,
		},
		{
			Block: `
				switch {
					case { false } then { out: 1 }
					case { false } then { out: 2 }
					case { false } then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "4\n",
		},
		/////
		{
			Block: `
				switch {
					case { false }  { out: 1 }
					case { false }  { out: 2 }
					case { false }  { out: 3 }
				}`,
			Stdout:  "",
			ExitNum: 1,
		},
		{
			Block: `
				switch {
					case { false }  { out: 1 }
					case { false }  { out: 2 }
					case { false }  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "4\n",
		},
	}

	test.RunMurexTests(tests, t)
}

/////

func TestSwitchIfBlock1(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				switch {
					if { true } then { out: 1 }
					if { false } then { out: 2 }
					if { false } then { out: 3 }
				}`,
			Stdout:  "1\n",
			ExitNum: 1,
		},
		{
			Block: `
				switch {
					if { true } then { out: 1 }
					if { true } then { out: 2 }
					if { false } then { out: 3 }
				}`,
			Stdout:  "1\n2\n",
			ExitNum: 1,
		},
		{
			Block: `
				switch {
					if { true } then { out: 1 }
					if { true } then { out: 2 }
					if { true } then { out: 3 }
				}`,
			Stdout:  "1\n2\n3\n",
			ExitNum: 1,
		},
		{
			Block: `
				switch {
					if { true } then { out: 1 }
					if { false } then { out: 2 }
					if { false } then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "1\n4\n",
		},
		{
			Block: `
				switch {
					if { true } then { out: 1 }
					if { true } then { out: 2 }
					if { false } then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "1\n2\n4\n",
		},
		{
			Block: `
				switch {
					if { true } then { out: 1 }
					if { true } then { out: 2 }
					if { true } then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "1\n2\n3\n4\n",
		},
		/////
		{
			Block: `
				switch {
					if { true }  { out: 1 }
					if { false }  { out: 2 }
					if { false }  { out: 3 }
				}`,
			Stdout:  "1\n",
			ExitNum: 1,
		},
		{
			Block: `
				switch {
					if { true }  { out: 1 }
					if { true }  { out: 2 }
					if { false }  { out: 3 }
				}`,
			Stdout:  "1\n2\n",
			ExitNum: 1,
		},
		{
			Block: `
				switch {
					if { true }  { out: 1 }
					if { true }  { out: 2 }
					if { true }  { out: 3 }
				}`,
			Stdout:  "1\n2\n3\n",
			ExitNum: 1,
		},
		{
			Block: `
				switch {
					if { true }  { out: 1 }
					if { false }  { out: 2 }
					if { false }  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "1\n4\n",
		},
		{
			Block: `
				switch {
					if { true }  { out: 1 }
					if { true }  { out: 2 }
					if { false }  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "1\n2\n4\n",
		},
		{
			Block: `
				switch {
					if { true }  { out: 1 }
					if { true }  { out: 2 }
					if { true }  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "1\n2\n3\n4\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestSwitchIfBlock2(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				switch {
					if { false } then { out: 1 }
					if { true } then { out: 2 }
					if { false } then { out: 3 }
				}`,
			Stdout:  "2\n",
			ExitNum: 1,
		},
		{
			Block: `
				switch {
					if { false } then { out: 1 }
					if { true } then { out: 2 }
					if { true } then { out: 3 }
				}`,
			Stdout:  "2\n3\n",
			ExitNum: 1,
		},
		{
			Block: `
				switch {
					if { false } then { out: 1 }
					if { true } then { out: 2 }
					if { false } then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "2\n4\n",
		},
		{
			Block: `
				switch {
					if { false } then { out: 1 }
					if { true } then { out: 2 }
					if { true } then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "2\n3\n4\n",
		},
		/////
		{
			Block: `
				switch {
					if { false }  { out: 1 }
					if { true }  { out: 2 }
					if { false }  { out: 3 }
				}`,
			Stdout:  "2\n",
			ExitNum: 1,
		},
		{
			Block: `
				switch {
					if { false }  { out: 1 }
					if { true }  { out: 2 }
					if { true }  { out: 3 }
				}`,
			Stdout:  "2\n3\n",
			ExitNum: 1,
		},
		{
			Block: `
				switch {
					if { false }  { out: 1 }
					if { true }  { out: 2 }
					if { false }  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "2\n4\n",
		},
		{
			Block: `
				switch {
					if { false }  { out: 1 }
					if { true }  { out: 2 }
					if { true }  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "2\n3\n4\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestSwitchIfBlock3(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				switch {
					if { false } then { out: 1 }
					if { false } then { out: 2 }
					if { true } then { out: 3 }
				}`,
			Stdout:  "3\n",
			ExitNum: 1,
		},
		{
			Block: `
				switch {
					if { false } then { out: 1 }
					if { false } then { out: 2 }
					if { true } then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "3\n4\n",
		},
		/////
		{
			Block: `
				switch {
					if { false }  { out: 1 }
					if { false }  { out: 2 }
					if { true }  { out: 3 }
				}`,
			Stdout:  "3\n",
			ExitNum: 1,
		},
		{
			Block: `
				switch {
					if { false }  { out: 1 }
					if { false }  { out: 2 }
					if { true }  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "3\n4\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestSwitchIfBlock4(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				switch {
					if { false } then { out: 1 }
					if { false } then { out: 2 }
					if { false } then { out: 3 }
				}`,
			Stdout:  "",
			ExitNum: 1,
		},
		{
			Block: `
				switch {
					if { false } then { out: 1 }
					if { false } then { out: 2 }
					if { false } then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "4\n",
		},
		/////
		{
			Block: `
				switch {
					if { false }  { out: 1 }
					if { false }  { out: 2 }
					if { false }  { out: 3 }
				}`,
			Stdout:  "",
			ExitNum: 1,
		},
		{
			Block: `
				switch {
					if { false }  { out: 1 }
					if { false }  { out: 2 }
					if { false }  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "4\n",
		},
	}

	test.RunMurexTests(tests, t)
}
