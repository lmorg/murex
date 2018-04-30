package structs

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/core/typemgmt"
)

// TestAnd tests the `and` builtin
func TestAnd(t *testing.T) {
	tests := []test{
		// --- and ---
		{
			block:  "and { true } { true } { true }",
			result: true,
		},
		{
			block:  "and { false } { true } { true }",
			result: false,
		},
		{
			block:  "and { true } { false } { true }",
			result: false,
		},
		{
			block:  "and { true } { true } { false }",
			result: false,
		},
		{
			block:  "and { false } { true } { false }",
			result: false,
		},
		{
			block:  "and { false } { false } { false }",
			result: false,
		},
		// --- !and ---
		{
			block:  "!and { true } { true } { true }",
			result: false,
		},
		{
			block:  "!and { false } { true } { true }",
			result: false,
		},
		{
			block:  "!and { true } { false } { true }",
			result: false,
		},
		{
			block:  "!and { true } { true } { false }",
			result: false,
		},
		{
			block:  "!and { false } { true } { false }",
			result: false,
		},
		{
			block:  "!and { false } { false } { false }",
			result: true,
		},
	}

	runTests(tests, t)
}
