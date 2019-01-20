package structs

import (
	"testing"

	"github.com/lmorg/murex/test"
)

// TestIf tests the `if` builtin
func TestIf(t *testing.T) {
	tests := []test.BooleanTest{
		// --- if / then---
		{
			Block:  "if { true } then { true }",
			Result: true,
		},
		{
			Block:  "if { true } then { false }",
			Result: false,
		},
		{
			Block:  "if { false } then { true }",
			Result: false,
		},
		{
			Block:  "if { false } then { false }",
			Result: false,
		},

		{
			Block:  "if { true }  { true }",
			Result: true,
		},
		{
			Block:  "if { true }  { false }",
			Result: false,
		},
		{
			Block:  "if { false }  { true }",
			Result: false,
		},
		{
			Block:  "if { false }  { false }",
			Result: false,
		},
		// --- if / then / else ---
		{
			Block:  "if { true } then { true } else { false }",
			Result: true,
		},
		{
			Block:  "if { true } then { false } else { false }",
			Result: false,
		},
		{
			Block:  "if { false } then { true } else { false }",
			Result: false,
		},
		{
			Block:  "if { false } then { false } else { false }",
			Result: false,
		},

		{
			Block:  "if { true }  { true }  { false }",
			Result: true,
		},
		{
			Block:  "if { true }  { false }  { false }",
			Result: false,
		},
		{
			Block:  "if { false }  { true }  { false }",
			Result: false,
		},
		{
			Block:  "if { false }  { false }  { false }",
			Result: false,
		},
		// ---
		{
			Block:  "if { true } then { true } else { true }",
			Result: true,
		},
		{
			Block:  "if { true } then { false } else { true }",
			Result: false,
		},
		{
			Block:  "if { false } then { true } else { true }",
			Result: true,
		},
		{
			Block:  "if { false } then { false } else { true }",
			Result: true,
		},

		{
			Block:  "if { true }  { true }  { true }",
			Result: true,
		},
		{
			Block:  "if { true }  { false }  { true }",
			Result: false,
		},
		{
			Block:  "if { false }  { true }  { true }",
			Result: true,
		},
		{
			Block:  "if { false }  { false }  { true }",
			Result: true,
		},
	}

	test.RunBooleanTests(tests, t)
}

// TestNotIf tests the `!if` builtin
func TestNotIf(t *testing.T) {
	tests := []test.BooleanTest{
		// --- !if / then---
		{
			Block:  "!if { true } then { true }",
			Result: false,
		},
		{
			Block:  "!if { true } then { false }",
			Result: false,
		},
		{
			Block:  "!if { false } then { true }",
			Result: true,
		},
		{
			Block:  "!if { false } then { false }",
			Result: false,
		},

		{
			Block:  "!if { true }  { true }",
			Result: false,
		},
		{
			Block:  "!if { true }  { false }",
			Result: false,
		},
		{
			Block:  "!if { false }  { true }",
			Result: true,
		},
		{
			Block:  "!if { false }  { false }",
			Result: false,
		},
		// --- !if / then / else ---
		{
			Block:  "!if { true } then { true } else { false }",
			Result: false,
		},
		{
			Block:  "!if { true } then { false } else { false }",
			Result: false,
		},
		{
			Block:  "!if { false } then { true } else { false }",
			Result: true,
		},
		{
			Block:  "!if { false } then { false } else { false }",
			Result: false,
		},

		{
			Block:  "!if { true }  { true }  { false }",
			Result: false,
		},
		{
			Block:  "!if { true }  { false }  { false }",
			Result: false,
		},
		{
			Block:  "!if { false }  { true }  { false }",
			Result: true,
		},
		{
			Block:  "!if { false }  { false }  { false }",
			Result: false,
		},
		// ---
		{
			Block:  "!if { true } then { true } else { true }",
			Result: true,
		},
		{
			Block:  "!if { true } then { false } else { true }",
			Result: true,
		},
		{
			Block:  "!if { false } then { true } else { true }",
			Result: true,
		},
		{
			Block:  "!if { false } then { false } else { true }",
			Result: false,
		},

		{
			Block:  "!if { true }  { true }  { true }",
			Result: true,
		},
		{
			Block:  "!if { true }  { false }  { true }",
			Result: true,
		},
		{
			Block:  "!if { false }  { true }  { true }",
			Result: true,
		},
		{
			Block:  "!if { false }  { false }  { true }",
			Result: false,
		},
	}

	test.RunBooleanTests(tests, t)
}
