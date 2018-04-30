package structs

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/core/typemgmt"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang/proc"
)

// TestIf tests the `if` builtin
func TestIf(t *testing.T) {
	defaults.Defaults(proc.InitConf, false)
	proc.InitEnv()

	tests := []test{
		// --- if / then---
		{
			block:  "if { true } then { true }",
			result: true,
		},
		{
			block:  "if { true } then { false }",
			result: false,
		},
		{
			block:  "if { false } then { true }",
			result: false,
		},
		{
			block:  "if { false } then { false }",
			result: false,
		},

		{
			block:  "if { true }  { true }",
			result: true,
		},
		{
			block:  "if { true }  { false }",
			result: false,
		},
		{
			block:  "if { false }  { true }",
			result: false,
		},
		{
			block:  "if { false }  { false }",
			result: false,
		},
		// --- if / then / else ---
		{
			block:  "if { true } then { true } else { false }",
			result: true,
		},
		{
			block:  "if { true } then { false } else { false }",
			result: false,
		},
		{
			block:  "if { false } then { true } else { false }",
			result: false,
		},
		{
			block:  "if { false } then { false } else { false }",
			result: false,
		},

		{
			block:  "if { true }  { true }  { false }",
			result: true,
		},
		{
			block:  "if { true }  { false }  { false }",
			result: false,
		},
		{
			block:  "if { false }  { true }  { false }",
			result: false,
		},
		{
			block:  "if { false }  { false }  { false }",
			result: false,
		},
		// ---
		{
			block:  "if { true } then { true } else { true }",
			result: true,
		},
		{
			block:  "if { true } then { false } else { true }",
			result: false,
		},
		{
			block:  "if { false } then { true } else { true }",
			result: true,
		},
		{
			block:  "if { false } then { false } else { true }",
			result: true,
		},

		{
			block:  "if { true }  { true }  { true }",
			result: true,
		},
		{
			block:  "if { true }  { false }  { true }",
			result: false,
		},
		{
			block:  "if { false }  { true }  { true }",
			result: true,
		},
		{
			block:  "if { false }  { false }  { true }",
			result: true,
		},
	}

	runTests(tests, t)
}

// TestNotIf tests the `!if` builtin
func TestNotIf(t *testing.T) {
	defaults.Defaults(proc.InitConf, false)
	proc.InitEnv()

	tests := []test{
		// --- !if / then---
		{
			block:  "!if { true } then { true }",
			result: false,
		},
		{
			block:  "!if { true } then { false }",
			result: false,
		},
		{
			block:  "!if { false } then { true }",
			result: true,
		},
		{
			block:  "!if { false } then { false }",
			result: false,
		},

		{
			block:  "!if { true }  { true }",
			result: false,
		},
		{
			block:  "!if { true }  { false }",
			result: false,
		},
		{
			block:  "!if { false }  { true }",
			result: true,
		},
		{
			block:  "!if { false }  { false }",
			result: false,
		},
		// --- !if / then / else ---
		{
			block:  "!if { true } then { true } else { false }",
			result: false,
		},
		{
			block:  "!if { true } then { false } else { false }",
			result: false,
		},
		{
			block:  "!if { false } then { true } else { false }",
			result: true,
		},
		{
			block:  "!if { false } then { false } else { false }",
			result: false,
		},

		{
			block:  "!if { true }  { true }  { false }",
			result: false,
		},
		{
			block:  "!if { true }  { false }  { false }",
			result: false,
		},
		{
			block:  "!if { false }  { true }  { false }",
			result: true,
		},
		{
			block:  "!if { false }  { false }  { false }",
			result: false,
		},
		// ---
		{
			block:  "!if { true } then { true } else { true }",
			result: true,
		},
		{
			block:  "!if { true } then { false } else { true }",
			result: true,
		},
		{
			block:  "!if { false } then { true } else { true }",
			result: true,
		},
		{
			block:  "!if { false } then { false } else { true }",
			result: false,
		},

		{
			block:  "!if { true }  { true }  { true }",
			result: true,
		},
		{
			block:  "!if { true }  { false }  { true }",
			result: true,
		},
		{
			block:  "!if { false }  { true }  { true }",
			result: true,
		},
		{
			block:  "!if { false }  { false }  { true }",
			result: false,
		},
	}

	runTests(tests, t)
}
