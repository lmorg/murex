package structs_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
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
			ExitNum: 0,
		},
		{
			Block: `
				switch {
					if { true } then { out: 1 }
					if { true } then { out: 2 }
					if { false } then { out: 3 }
				}`,
			Stdout:  "1\n2\n",
			ExitNum: 0,
		},
		{
			Block: `
				switch {
					if { true } then { out: 1 }
					if { true } then { out: 2 }
					if { true } then { out: 3 }
				}`,
			Stdout:  "1\n2\n3\n",
			ExitNum: 0,
		},
		{
			Block: `
				switch {
					if { true } then { out: 1 }
					if { false } then { out: 2 }
					if { false } then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "1\n",
		},
		{
			Block: `
				switch {
					if { true } then { out: 1 }
					if { true } then { out: 2 }
					if { false } then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "1\n2\n",
		},
		{
			Block: `
				switch {
					if { true } then { out: 1 }
					if { true } then { out: 2 }
					if { true } then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "1\n2\n3\n",
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
			ExitNum: 0,
		},
		{
			Block: `
				switch {
					if { true }  { out: 1 }
					if { true }  { out: 2 }
					if { false }  { out: 3 }
				}`,
			Stdout:  "1\n2\n",
			ExitNum: 0,
		},
		{
			Block: `
				switch {
					if { true }  { out: 1 }
					if { true }  { out: 2 }
					if { true }  { out: 3 }
				}`,
			Stdout:  "1\n2\n3\n",
			ExitNum: 0,
		},
		{
			Block: `
				switch {
					if { true }  { out: 1 }
					if { false }  { out: 2 }
					if { false }  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "1\n",
		},
		{
			Block: `
				switch {
					if { true }  { out: 1 }
					if { true }  { out: 2 }
					if { false }  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "1\n2\n",
		},
		{
			Block: `
				switch {
					if { true }  { out: 1 }
					if { true }  { out: 2 }
					if { true }  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "1\n2\n3\n",
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
			ExitNum: 0,
		},
		{
			Block: `
				switch {
					if { false } then { out: 1 }
					if { true } then { out: 2 }
					if { true } then { out: 3 }
				}`,
			Stdout:  "2\n3\n",
			ExitNum: 0,
		},
		{
			Block: `
				switch {
					if { false } then { out: 1 }
					if { true } then { out: 2 }
					if { false } then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "2\n",
		},
		{
			Block: `
				switch {
					if { false } then { out: 1 }
					if { true } then { out: 2 }
					if { true } then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "2\n3\n",
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
			ExitNum: 0,
		},
		{
			Block: `
				switch {
					if { false }  { out: 1 }
					if { true }  { out: 2 }
					if { true }  { out: 3 }
				}`,
			Stdout:  "2\n3\n",
			ExitNum: 0,
		},
		{
			Block: `
				switch {
					if { false }  { out: 1 }
					if { true }  { out: 2 }
					if { false }  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "2\n",
		},
		{
			Block: `
				switch {
					if { false }  { out: 1 }
					if { true }  { out: 2 }
					if { true }  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "2\n3\n",
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
			ExitNum: 0,
		},
		{
			Block: `
				switch {
					if { false } then { out: 1 }
					if { false } then { out: 2 }
					if { true } then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "3\n",
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
			ExitNum: 0,
		},
		{
			Block: `
				switch {
					if { false }  { out: 1 }
					if { false }  { out: 2 }
					if { true }  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "3\n",
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

func TestSwitchCaseIfBlock(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				switch {
					if { false } then { out: 1 }
					if { false } then { out: 2 }
					case { false } then { out: 3 }
				}`,
			Stdout:  "",
			ExitNum: 1,
		},
		{
			Block: `
				switch {
					if { false } then { out: 1 }
					if { false } then { out: 2 }
					case { false } then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "4\n",
		},
		{
			Block: `
				switch {
					if { false } then { out: 1 }
					if { true } then { out: 2 }
					case { false } then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "2\n",
		},
		{
			Block: `
				switch {
					if { false } then { out: 1 }
					if { true } then { out: 2 }
					case { true } then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "2\n3\n",
		},
		{
			Block: `
				switch {
					if { false } then { out: 1 }
					case { true } then { out: 2 }
					if { true } then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "2\n",
		},
		{
			Block: `
				switch {
					if { false } then { out: 1 }
					case { false } then { out: 2 }
					if { true } then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "3\n",
		},
		/////
		{
			Block: `
				switch {
					if { false }  { out: 1 }
					if { false }  { out: 2 }
					case { false }  { out: 3 }
				}`,
			Stdout:  "",
			ExitNum: 1,
		},
		{
			Block: `
				switch {
					if { false }  { out: 1 }
					if { false }  { out: 2 }
					case { false }  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "4\n",
		},
		{
			Block: `
				switch {
					if { false }  { out: 1 }
					if { true }  { out: 2 }
					case { false } then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "2\n",
		},
		{
			Block: `
				switch {
					if { false }  { out: 1 }
					if { true }  { out: 2 }
					case { true }  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "2\n3\n",
		},
		{
			Block: `
				switch {
					if { false }  { out: 1 }
					case { true }  { out: 2 }
					if { true }  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "2\n",
		},
		{
			Block: `
				switch {
					if { false }  { out: 1 }
					case { false }  { out: 2 }
					if { true }  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout: "3\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestSwitchByValCaseIfBlock1(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				switch foobar {
					if foo then { out: 1 }
					if bar then { out: 2 }
					if oof then { out: 3 }
				}`,
			Stdout:  "",
			ExitNum: 1,
		},
		{
			Block: `
				switch foobar {
					if foo then { out: 1 }
					if bar then { out: 2 }
					case oof then { out: 3 }
				}`,
			Stdout:  "",
			ExitNum: 1,
		},
		{
			Block: `
				switch foobar {
					if foo then { out: 1 }
					case bar then { out: 2 }
					if oof then { out: 3 }
				}`,
			Stdout:  "",
			ExitNum: 1,
		},
		{
			Block: `
				switch foobar {
					case foo then { out: 1 }
					if bar then { out: 2 }
					if oof then { out: 3 }
				}`,
			Stdout:  "",
			ExitNum: 1,
		},
		/////
		{
			Block: `
				switch foobar {
					if foo  { out: 1 }
					if bar  { out: 2 }
					if oof  { out: 3 }
				}`,
			Stdout:  "",
			ExitNum: 1,
		},
		{
			Block: `
				switch foobar {
					if foo  { out: 1 }
					if bar  { out: 2 }
					case oof  { out: 3 }
				}`,
			Stdout:  "",
			ExitNum: 1,
		},
		{
			Block: `
				switch foobar {
					if foo  { out: 1 }
					case bar  { out: 2 }
					if oof  { out: 3 }
				}`,
			Stdout:  "",
			ExitNum: 1,
		},
		{
			Block: `
				switch foobar {
					case foo  { out: 1 }
					if bar  { out: 2 }
					if oof  { out: 3 }
				}`,
			Stdout:  "",
			ExitNum: 1,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestSwitchByValCaseIfBlock2(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				switch foobar {
					if foo then { out: 1 }
					if bar then { out: 2 }
					if oof then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout:  "4\n",
			ExitNum: 0,
		},
		{
			Block: `
				switch foobar {
					if foo then { out: 1 }
					if bar then { out: 2 }
					case oof then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout:  "4\n",
			ExitNum: 0,
		},
		{
			Block: `
				switch foobar {
					if foo then { out: 1 }
					case bar then { out: 2 }
					if oof then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout:  "4\n",
			ExitNum: 0,
		},
		{
			Block: `
				switch foobar {
					case foo then { out: 1 }
					if bar then { out: 2 }
					if oof then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout:  "4\n",
			ExitNum: 0,
		},
		/////
		{
			Block: `
				switch foobar {
					if foo  { out: 1 }
					if bar  { out: 2 }
					if oof  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout:  "4\n",
			ExitNum: 0,
		},
		{
			Block: `
				switch foobar {
					if foo  { out: 1 }
					if bar  { out: 2 }
					case oof  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout:  "4\n",
			ExitNum: 0,
		},
		{
			Block: `
				switch foobar {
					if foo  { out: 1 }
					case bar  { out: 2 }
					if oof  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout:  "4\n",
			ExitNum: 0,
		},
		{
			Block: `
				switch foobar {
					case foo  { out: 1 }
					if bar  { out: 2 }
					if oof  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout:  "4\n",
			ExitNum: 0,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestSwitchByValCaseIfBlock3(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				switch foobar {
					if foo then { out: 1 }
					if foobar then { out: 2 }
					if oof then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout:  "2\n",
			ExitNum: 0,
		},
		{
			Block: `
				switch foobar {
					if foo then { out: 1 }
					if foobar then { out: 2 }
					case oof then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout:  "2\n",
			ExitNum: 0,
		},
		{
			Block: `
				switch foobar {
					if foo then { out: 1 }
					case foobar then { out: 2 }
					if foobar then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout:  "2\n",
			ExitNum: 0,
		},
		{
			Block: `
				switch foobar {
					case foo then { out: 1 }
					if foobar then { out: 2 }
					if oof then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout:  "2\n",
			ExitNum: 0,
		},
		/////
		{
			Block: `
				switch foobar {
					if foo  { out: 1 }
					if foobar  { out: 2 }
					if oof  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout:  "2\n",
			ExitNum: 0,
		},
		{
			Block: `
				switch foobar {
					if foo  { out: 1 }
					if foobar  { out: 2 }
					case oof  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout:  "2\n",
			ExitNum: 0,
		},
		{
			Block: `
				switch foobar {
					if foo  { out: 1 }
					case foobar  { out: 2 }
					if foobar  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout:  "2\n",
			ExitNum: 0,
		},
		{
			Block: `
				switch foobar {
					case foo  { out: 1 }
					if foobar  { out: 2 }
					if oof  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout:  "2\n",
			ExitNum: 0,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestSwitchByValCaseIfBlock4(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				switch foobar {
					if foo then { out: 1 }
					if { out: foobar } then { out: 2 }
					if oof then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout:  "2\n",
			ExitNum: 0,
		},
		{
			Block: `
				switch foobar {
					if foo then { out: 1 }
					if { out: foobar } then { out: 2 }
					case oof then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout:  "2\n",
			ExitNum: 0,
		},
		{
			Block: `
				switch foobar {
					if foo then { out: 1 }
					case { out: foobar } then { out: 2 }
					if foobar then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout:  "2\n",
			ExitNum: 0,
		},
		{
			Block: `
				switch foobar {
					case foo then { out: 1 }
					if { out: foobar } then { out: 2 }
					if oof then { out: 3 }
					catch { out: 4 }
				}`,
			Stdout:  "2\n",
			ExitNum: 0,
		},
		/////
		{
			Block: `
				switch foobar {
					if foo  { out: 1 }
					if { out: foobar }  { out: 2 }
					if oof  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout:  "2\n",
			ExitNum: 0,
		},
		{
			Block: `
				switch foobar {
					if foo  { out: 1 }
					if { out: foobar }  { out: 2 }
					case oof  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout:  "2\n",
			ExitNum: 0,
		},
		{
			Block: `
				switch foobar {
					if foo  { out: 1 }
					case { out: foobar }  { out: 2 }
					if foobar  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout:  "2\n",
			ExitNum: 0,
		},
		{
			Block: `
				switch foobar {
					case foo  { out: 1 }
					if { out: foobar }  { out: 2 }
					if oof  { out: 3 }
					catch { out: 4 }
				}`,
			Stdout:  "2\n",
			ExitNum: 0,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestSwitchErrors(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:   `switch`,
			Stderr:  `no parameters`,
			ExitNum: 1,
		},
		{
			Block:   `switch foobar`,
			Stderr:  `not a valid code block`,
			ExitNum: 1,
		},
		{
			Block:   `switch foo bar`,
			Stderr:  `not a valid code block`,
			ExitNum: 1,
		},
		{
			Block:   `switch foo bar baz`,
			Stderr:  `too many`,
			ExitNum: 1,
		},
		{
			Block:   `switch {"}`,
			Stderr:  `missing closing quote`,
			ExitNum: 1,
		},
		{
			Block:   `switch foo {"}`,
			Stderr:  `missing closing quote`,
			ExitNum: 1,
		},
		{
			Block:   `switch foo { foobar }`,
			Stderr:  `not a valid statement`,
			ExitNum: 1,
		},
		{
			Block:   "switch foo { if {false} }",
			Stderr:  `too few parameters`,
			ExitNum: 1,
		},
		/*{
			Block:   "switch foo { if {)} { out yes } }",
			Stderr:  `if conditional:\n\s+> syntax error at`,
			ExitNum: 1,
		},*/
		/*{
			Block:   "switch foo { if {true} {)} }",
			Stderr:  `if conditional:\n\s+> syntax error at`,
			ExitNum: 1,
		},*/
		/*{
			Block:   "switch foo { case {)} { out yes } }",
			Stderr:  `case conditional:\n\s+> syntax error at`,
			ExitNum: 1,
		},*/
		/*{
			Block:   "switch foo { case {true} {)} }",
			Stderr:  `if conditional:\n\s+> syntax error at`,
			ExitNum: 1,
		},*/
		/*{
			Block:   "switch foo { catch {)} }",
			Stderr:  `catch block:\n\s+> syntax error at`,
			ExitNum: 1,
		},*/
		{ // nothing matched but no error
			Block:   `switch foo { if "bob" { out "yes" } }`,
			Stdout:  `^$`,
			Stderr:  `^$`,
			ExitNum: 1,
		},
		/*{ // nothing matched but no error // TODO: investigate
			Block:   `switch foo { if "bob" { out "yes" }; catch {} }`,
			Stdout:  `^$`,
			Stderr:  `^$`,
			ExitNum: 0,
		},*/
	}

	test.RunMurexTestsRx(tests, t)
}
