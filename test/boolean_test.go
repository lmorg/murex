package test

import (
	"testing"
)

// TestBool proves the boolean test framework works
func TestBool(t *testing.T) {
	tests := []BooleanTest{
		{
			Block:  "true",
			Result: true,
		},
		{
			Block:  "false",
			Result: false,
		},
	}

	RunBooleanTests(tests, t)
}
