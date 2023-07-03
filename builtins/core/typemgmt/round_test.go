package typemgmt

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

type roundTest struct {
	Start    float64
	Round    float64
	Expected float64
}

func TestRoundUpInteger(t *testing.T) {
	tests := []roundTest{
		{Start: 0, Expected: 0},
		{Start: 0.1, Expected: 1},
		{Start: 0.5, Expected: 1},
		{Start: 0.9, Expected: 1},
		{Start: 1, Expected: 1},
		{Start: 1.1, Expected: 2},
		{Start: 1.5, Expected: 2},
		{Start: 1.9, Expected: 2},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		actual := roundUpInteger(test.Start)
		if actual != int(test.Expected) {
			t.Errorf("Expected does not match actual in test %d", i)
			t.Logf("  Start:    %f", test.Start)
			t.Logf("  Round:    %f", test.Round)
			t.Logf("  Expected: %f", test.Expected)
			t.Logf("  Actual:   %d", actual)
		}
	}
}

func TestRoundNearestMultiple(t *testing.T) {
	tests := []roundTest{
		{Start: 0, Round: 5, Expected: 0},
		{Start: 1, Round: 5, Expected: 0},
		{Start: 2, Round: 5, Expected: 0},
		{Start: 3, Round: 5, Expected: 5},
		{Start: 4, Round: 5, Expected: 5},
		{Start: 5, Round: 5, Expected: 5},
		{Start: 6, Round: 5, Expected: 5},
		{Start: 7, Round: 5, Expected: 5},
		{Start: 8, Round: 5, Expected: 10},
		{Start: 9, Round: 5, Expected: 10},
		{Start: 10, Round: 5, Expected: 10},
		{Start: 11, Round: 5, Expected: 10},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		actual := roundNearestMultiple(int(test.Start), int(test.Round))
		if actual != int(test.Expected) {
			t.Errorf("Expected does not match actual in test %d", i)
			t.Logf("  Start:    %f", test.Start)
			t.Logf("  Round:    %f", test.Round)
			t.Logf("  Expected: %f", test.Expected)
			t.Logf("  Actual:   %d", actual)
		}
	}
}

func TestRoundDownMultiple(t *testing.T) {
	tests := []roundTest{
		{Start: 0, Round: 5, Expected: 0},
		{Start: 1, Round: 5, Expected: 0},
		{Start: 2, Round: 5, Expected: 0},
		{Start: 3, Round: 5, Expected: 0},
		{Start: 4, Round: 5, Expected: 0},
		{Start: 5, Round: 5, Expected: 5},
		{Start: 6, Round: 5, Expected: 5},
		{Start: 7, Round: 5, Expected: 5},
		{Start: 8, Round: 5, Expected: 5},
		{Start: 9, Round: 5, Expected: 5},
		{Start: 10, Round: 5, Expected: 10},
		{Start: 11, Round: 5, Expected: 10},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		actual := roundDownMultiple(int(test.Start), int(test.Round))
		if actual != int(test.Expected) {
			t.Errorf("Expected does not match actual in test %d", i)
			t.Logf("  Start:    %f", test.Start)
			t.Logf("  Round:    %f", test.Round)
			t.Logf("  Expected: %f", test.Expected)
			t.Logf("  Actual:   %d", actual)
		}
	}
}

func TestRoundUpMultiple(t *testing.T) {
	tests := []roundTest{
		{Start: 0, Round: 5, Expected: 0},
		{Start: 1, Round: 5, Expected: 5},
		{Start: 2, Round: 5, Expected: 5},
		{Start: 3, Round: 5, Expected: 5},
		{Start: 4, Round: 5, Expected: 5},
		{Start: 5, Round: 5, Expected: 5},
		{Start: 6, Round: 5, Expected: 10},
		{Start: 7, Round: 5, Expected: 10},
		{Start: 8, Round: 5, Expected: 10},
		{Start: 9, Round: 5, Expected: 10},
		{Start: 10, Round: 5, Expected: 10},
		{Start: 11, Round: 5, Expected: 15},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		actual := roundUpMultiple(int(test.Start), int(test.Round))
		if actual != int(test.Expected) {
			t.Errorf("Expected does not match actual in test %d", i)
			t.Logf("  Start:    %f", test.Start)
			t.Logf("  Round:    %f", test.Round)
			t.Logf("  Expected: %f", test.Expected)
			t.Logf("  Actual:   %d", actual)
		}
	}
}

func TestRoundNearestDecimalPlace(t *testing.T) {
	tests := []roundTest{
		{Start: 0, Round: 2, Expected: 0},
		{Start: 1, Round: 2, Expected: 1},
		{Start: 1.1, Round: 2, Expected: 1.1},
		{Start: 1.12, Round: 2, Expected: 1.12},
		{Start: 1.123, Round: 2, Expected: 1.12},
		{Start: 1.1234, Round: 2, Expected: 1.12},
		{Start: 1.1234, Round: 3, Expected: 1.123},
		{Start: 1.1234, Round: 4, Expected: 1.1234},
		{Start: 1.12345, Round: 5, Expected: 1.12345},
		{Start: 1.123456, Round: 5, Expected: 1.12346},
		{Start: 99.99, Round: 1, Expected: 100},
		{Start: 99.99, Round: 2, Expected: 99.99},
		{Start: 99.99, Round: 3, Expected: 99.99},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		actual := roundNearestDecimalPlace(test.Start, int(test.Round))
		if actual != test.Expected {
			t.Errorf("Expected does not match actual in test %d", i)
			t.Logf("  Start:    %f", test.Start)
			t.Logf("  Round:    %f", test.Round)
			t.Logf("  Expected: %f", test.Expected)
			t.Logf("  Actual:   %f", actual)
		}
	}
}
