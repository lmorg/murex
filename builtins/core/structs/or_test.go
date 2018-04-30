package structs

import (
	"testing"

	"github.com/lmorg/murex/test"
)

// TestOr tests the `or` builtin
func TestOr(t *testing.T) {
	tests := []test.BooleanTest{
		// --- or ---
		{
			Block:  "or { true } { true } { true }",
			Result: true,
		},
		{
			Block:  "or { false } { true } { true }",
			Result: true,
		},
		{
			Block:  "or { true } { false } { true }",
			Result: true,
		},
		{
			Block:  "or { true } { true } { false }",
			Result: true,
		},
		{
			Block:  "or { false } { true } { false }",
			Result: true,
		},
		{
			Block:  "or { false } { false } { false }",
			Result: false,
		},
		// --- !or ---
		{
			Block:  "!or { true } { true } { true }",
			Result: false,
		},
		{
			Block:  "!or { false } { true } { true }",
			Result: true,
		},
		{
			Block:  "!or { true } { false } { true }",
			Result: true,
		},
		{
			Block:  "!or { true } { true } { false }",
			Result: true,
		},
		{
			Block:  "!or { false } { true } { false }",
			Result: true,
		},
		{
			Block:  "!or { false } { false } { false }",
			Result: true,
		},
	}

	test.RunBooleanTests(tests, t)
}
