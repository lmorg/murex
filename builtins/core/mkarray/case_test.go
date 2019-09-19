package mkarray

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestGetCase(t *testing.T) {
	count.Tests(t, 4)

	if getCase("foobar") != caseLower {
		t.Error("`foobar` not being detected as lower case")
	}

	if getCase("Foobar") != caseFirst {
		t.Error("`Foobar` not being detected as first case")
	}

	if getCase("FOOBAR") != caseUpper {
		t.Error("`FOOBAR` not being detected as upper case")
	}

	if getCase("Foo Bar") != caseTitle {
		t.Error("`Foo Bar` not being detected as title case")
	}
}

func TestSetCase(t *testing.T) {
	count.Tests(t, 3)

	input := "foobar"

	expected := "foobar"
	output := setCase(input, caseLower)
	if output != expected {
		t.Error("setCase not lower casing correctly")
		t.Log("  input:    ", input)
		t.Log("  output:   ", output)
		t.Log("  expected: ", expected)
	}

	expected = "Foobar"
	output = setCase(input, caseTitle)
	if output != expected {
		t.Error("setCase not title casing correctly")
		t.Log("  input:    ", input)
		t.Log("  output:   ", output)
		t.Log("  expected: ", expected)
	}

	expected = "FOOBAR"
	output = setCase(input, caseUpper)
	if output != expected {
		t.Error("setCase not upper casing correctly")
		t.Log("  input:    ", input)
		t.Log("  output:   ", output)
		t.Log("  expected: ", expected)
	}
}

// TestOptimisedSetCase checks that nobody tries to "bug fix" the setCase()
// function with lowercasing already lowercased elements
func TestOptimisedSetCase(t *testing.T) {
	count.Tests(t, 2)

	input := "fooBar"

	expected := "foobar"
	output := setCase(input, caseLower)
	if output == expected {
		t.Error("setCase(s, caseLower) has been changed to lower case when shouldn't")
		t.Log("  input:    ", input)
		t.Log("  output:   ", output)
		t.Log("  expected: ", input)
		t.Log("All elements should be lowercase by default, so we don't need to waste time lowercasing the string")
	}

	expected = "Foobar"
	output = setCase(input, caseTitle)
	if output == expected {
		t.Error("setCase(s, caseTitle) has been changed to lower case when shouldn't")
		t.Log("  input:    ", input)
		t.Log("  output:   ", output)
		t.Log("  expected: ", "FooBar")
		t.Log("All elements should be lowercase by default, so we don't need to waste time lowercasing most of the string")
	}
}
