package test_test

import (
	"testing"

	_ "github.com/lmorg/murex/lang/expressions"
	"github.com/lmorg/murex/test"
)

// TestBool proves the boolean test framework works
func TestBool(t *testing.T) {
	tests := []test.BooleanTest{
		{
			Block:  "true",
			Result: true,
		},
		{
			Block:  "false",
			Result: false,
		},
	}

	test.RunBooleanTests(tests, t)
}
