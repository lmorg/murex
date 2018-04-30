package structs

import (
	"testing"
)

// TestOr tests the `or` builtin
func TestOr(t *testing.T) {
	tests := []test{
		// --- or ---
		{
			block:  "or { true } { true } { true }",
			result: true,
		},
		{
			block:  "or { false } { true } { true }",
			result: true,
		},
		{
			block:  "or { true } { false } { true }",
			result: true,
		},
		{
			block:  "or { true } { true } { false }",
			result: true,
		},
		{
			block:  "or { false } { true } { false }",
			result: true,
		},
		{
			block:  "or { false } { false } { false }",
			result: false,
		},
		// --- !or ---
		{
			block:  "!or { true } { true } { true }",
			result: false,
		},
		{
			block:  "!or { false } { true } { true }",
			result: true,
		},
		{
			block:  "!or { true } { false } { true }",
			result: true,
		},
		{
			block:  "!or { true } { true } { false }",
			result: true,
		},
		{
			block:  "!or { false } { true } { false }",
			result: true,
		},
		{
			block:  "!or { false } { false } { false }",
			result: true,
		},
	}

	runTests(tests, t)
}
