package structs

import (
	"testing"

	"github.com/lmorg/murex/test"
)

// TestTry tests the `try` builtin
func TestTry(t *testing.T) {
	tests := []test.BooleanTest{
		{
			Block:  "try { true -> true }",
			Result: true,
		},
		{
			Block:  "try { false -> true }",
			Result: true,
		},
		{
			Block:  "try { true -> false }",
			Result: false,
		},
		{
			Block:  "try { true; true }",
			Result: true,
		},
		{
			Block:  "try { false; true }",
			Result: false,
		},
		{
			Block:  "try { true; false }",
			Result: false,
		},
	}

	test.RunBooleanTests(tests, t)
}

// TestTryPipe tests the `trypipe` builtin
func TestTryPipe(t *testing.T) {
	tests := []test.BooleanTest{
		{
			Block:  "trypipe { true -> true }",
			Result: true,
		},
		{
			Block:  "trypipe { false -> true }",
			Result: false,
		},
		{
			Block:  "trypipe { true -> false }",
			Result: false,
		},
		{
			Block:  "trypipe { true; true }",
			Result: true,
		},
		{
			Block:  "trypipe { false; true }",
			Result: false,
		},
		{
			Block:  "trypipe { true; false }",
			Result: false,
		},
	}

	test.RunBooleanTests(tests, t)
}

// TestCatch tests the `catch` builtin
func TestCatch(t *testing.T) {
	tests := []test.BooleanTest{
		// --- catch ---
		{
			Block:  "try <null> { true }; catch { true }",
			Result: false,
		},
		{
			Block:  "try <null> { false }; catch { true }",
			Result: false,
		},
		{
			Block:  "try <null> { true }; catch { false }",
			Result: false,
		},
		{
			Block:  "try <null> { false }; catch { false }",
			Result: false,
		},
		// --- !catch ---
		{
			Block:  "try <null> { true }; !catch { true }",
			Result: true,
		},
		{
			Block:  "try <null> { false }; !catch { true }",
			Result: false,
		},
		{
			Block:  "try <null> { true }; !catch { false }",
			Result: false,
		},
		{
			Block:  "try <null> { false }; !catch { false }",
			Result: false,
		},
	}

	test.RunBooleanTests(tests, t)
}

func TestUnsafe(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `try {
						unsafe { err "foo" }
						out "bar"
					}`,
			Stdout: "bar\n",
			Stderr: "foo\n",
		},
	}

	test.RunMurexTests(tests, t)
}
