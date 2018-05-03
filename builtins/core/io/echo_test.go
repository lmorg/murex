package io

import (
	"testing"

	"github.com/lmorg/murex/test"
)

// TestOut tests the `out` builtin
func TestOut(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  "out: test",
			Stdout: "test\n",
		},
		{
			Block:  "out: t e s t",
			Stdout: "t e s t\n",
		},
		{
			Block:  "echo: test",
			Stdout: "test\n",
		},
		{
			Block:  "echo: t e s t",
			Stdout: "t e s t\n",
		},
	}
	test.RunMurexTests(tests, t)
}

// TestTout tests the `tout` builtin
func TestTout(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  "tout: str test",
			Stdout: "test",
		},
		{
			Block:  "tout: str t e s t",
			Stdout: "t e s t",
		},
	}
	test.RunMurexTests(tests, t)
}

// TestBrace tests the `(` builtin
func TestBrace(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  "(test)",
			Stdout: "test",
		},
		{
			Block:  "(t e s t)",
			Stdout: "t e s t",
		},
	}
	test.RunMurexTests(tests, t)
}

// TestErr tests the `err` builtin
func TestErr(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:   "err: test",
			Stderr:  "test\n",
			ExitNum: 1,
		},
		{
			Block:   "err: t e s t",
			Stderr:  "t e s t\n",
			ExitNum: 1,
		},
	}
	test.RunMurexTests(tests, t)
}
