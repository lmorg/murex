package structs

import (
	"testing"

	"github.com/lmorg/murex/test"
)

// TestAnd tests the `and` builtin
func TestAnd(t *testing.T) {
	tests := []test.BooleanTest{
		// --- and ---
		{
			Block:  "and { true } { true } { true }",
			Result: true,
		},
		{
			Block:  "and { false } { true } { true }",
			Result: false,
		},
		{
			Block:  "and { true } { false } { true }",
			Result: false,
		},
		{
			Block:  "and { true } { true } { false }",
			Result: false,
		},
		{
			Block:  "and { false } { true } { false }",
			Result: false,
		},
		{
			Block:  "and { false } { false } { false }",
			Result: false,
		},
		// --- !and ---
		{
			Block:  "!and { true } { true } { true }",
			Result: false,
		},
		{
			Block:  "!and { false } { true } { true }",
			Result: false,
		},
		{
			Block:  "!and { true } { false } { true }",
			Result: false,
		},
		{
			Block:  "!and { true } { true } { false }",
			Result: false,
		},
		{
			Block:  "!and { false } { true } { false }",
			Result: false,
		},
		{
			Block:  "!and { false } { false } { false }",
			Result: true,
		},
	}

	test.RunBooleanTests(tests, t)
}
