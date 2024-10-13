package semver_test

import (
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/semver"
)

func TestParse(t *testing.T) {
	tests := []struct {
		String   string
		Expected semver.Version
		Error    bool
	}{
		{
			String: "1",
			Expected: semver.Version{
				Major: 1,
			},
		},
		{
			String: "1.2",
			Expected: semver.Version{
				Major: 1,
				Minor: 2,
			},
		},
		{
			String: "1.2.3",
			Expected: semver.Version{
				Major: 1,
				Minor: 2,
				Patch: 3,
			},
		},
		{
			String: "1.2.3.4",
			Error:  true,
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		actual, err := semver.Parse(test.String)
		if (err != nil) != test.Error {
			t.Errorf("Unexpected error in test %d:", i)
			t.Logf("  Expected: %v", test.Expected)
			t.Logf("  Actual:   %v", *actual)
			t.Logf("  exp err:  %v", test.Error)
			t.Logf("  act err:  %v", err)
		}

		if actual == nil {
			continue
		}

		expS := json.LazyLogging(test.Expected)
		actS := json.LazyLogging(*actual)

		if expS != actS {
			t.Errorf("Expected doesn't match actual in test %d:", i)
			t.Logf("  Expected: %v", expS)
			t.Logf("  Actual:   %v", actS)
			t.Logf("  exp err:  %v", test.Error)
			t.Logf("  act err:  %v", err)
		}
	}
}

func TestCompareFunc(t *testing.T) {
	tests := []struct {
		Version    string
		Comparison string
		Expected   bool
		Error      bool
	}{
		{
			Version:    "1.2.3",
			Comparison: "> 0",
			Expected:   true,
		},
		{
			Version:    "1.2.3",
			Comparison: ">= 0",
			Expected:   true,
		},
		{
			Version:    "1.2.3",
			Comparison: "= 0",
			Expected:   false,
		},
		{
			Version:    "1.2.3",
			Comparison: "<= 0",
			Expected:   false,
		},
		{
			Version:    "1.2.3",
			Comparison: "< 0",
			Expected:   false,
		},
		{
			Version:    "1.2.3",
			Comparison: "> 1",
			Expected:   false,
		},
		{
			Version:    "1.2.3",
			Comparison: ">= 1",
			Expected:   true,
		},
		{
			Version:    "1.2.3",
			Comparison: "= 1",
			Expected:   true,
		},
		{
			Version:    "1.2.3",
			Comparison: "<= 1",
			Expected:   true,
		},
		{
			Version:    "1.2.3",
			Comparison: "< 1",
			Expected:   false,
		},
		{
			Version:    "1.2.3",
			Comparison: "> 2",
			Expected:   false,
		},
		{
			Version:    "1.2.3",
			Comparison: ">= 2",
			Expected:   false,
		},
		{
			Version:    "1.2.3",
			Comparison: "= 2",
			Expected:   false,
		},
		{
			Version:    "1.2.3",
			Comparison: "<= 2",
			Expected:   true,
		},
		{
			Version:    "1.2.3",
			Comparison: "< 2",
			Expected:   true,
		},
		///
		{
			Version:    "1.2.3",
			Comparison: "> 1.1",
			Expected:   true,
		},
		{
			Version:    "1.2.3",
			Comparison: ">= 1.1",
			Expected:   true,
		},
		{
			Version:    "1.2.3",
			Comparison: "= 1.1",
			Expected:   false,
		},
		{
			Version:    "1.2.3",
			Comparison: "<= 1.1",
			Expected:   false,
		},
		{
			Version:    "1.2.3",
			Comparison: "< 1.1",
			Expected:   false,
		},
		{
			Version:    "1.2.3",
			Comparison: "> 1.2",
			Expected:   false,
		},
		{
			Version:    "1.2.3",
			Comparison: ">= 1.2",
			Expected:   true,
		},
		{
			Version:    "1.2.3",
			Comparison: "= 1.2",
			Expected:   true,
		},
		{
			Version:    "1.2.3",
			Comparison: "<= 1.2",
			Expected:   true,
		},
		{
			Version:    "1.2.3",
			Comparison: "< 1.2",
			Expected:   false,
		},
		{
			Version:    "1.2.3",
			Comparison: "> 1.3",
			Expected:   false,
		},
		{
			Version:    "1.2.3",
			Comparison: ">= 2.3",
			Expected:   false,
		},
		{
			Version:    "1.2.3",
			Comparison: "= 2.3",
			Expected:   false,
		},
		{
			Version:    "1.2.3",
			Comparison: "<= 2.3",
			Expected:   true,
		},
		{
			Version:    "1.2.3",
			Comparison: "< 2.3",
			Expected:   true,
		},
		///
		{
			Version:    "1.2.3",
			Comparison: "> 1.2.2",
			Expected:   true,
		},
		{
			Version:    "1.2.3",
			Comparison: ">= 1.2.2",
			Expected:   true,
		},
		{
			Version:    "1.2.3",
			Comparison: "= 1.2.2",
			Expected:   false,
		},
		{
			Version:    "1.2.3",
			Comparison: "<= 1.2.2",
			Expected:   false,
		},
		{
			Version:    "1.2.3",
			Comparison: "< 1.2.2",
			Expected:   false,
		},
		{
			Version:    "1.2.3",
			Comparison: "> 1.2.3",
			Expected:   false,
		},
		{
			Version:    "1.2.3",
			Comparison: ">= 1.2.3",
			Expected:   true,
		},
		{
			Version:    "1.2.3",
			Comparison: "= 1.2.3",
			Expected:   true,
		},
		{
			Version:    "1.2.3",
			Comparison: "<= 1.2.3",
			Expected:   true,
		},
		{
			Version:    "1.2.3",
			Comparison: "< 1.2.3",
			Expected:   false,
		},
		{
			Version:    "1.2.3",
			Comparison: "> 1.2.4",
			Expected:   false,
		},
		{
			Version:    "1.2.3",
			Comparison: ">= 1.2.4",
			Expected:   false,
		},
		{
			Version:    "1.2.3",
			Comparison: "= 1.2.4",
			Expected:   false,
		},
		{
			Version:    "1.2.3",
			Comparison: "<= 1.2.4",
			Expected:   true,
		},
		{
			Version:    "1.2.3",
			Comparison: "< 1.2.4",
			Expected:   true,
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		actual, err := semver.Compare(test.Version, test.Comparison)
		if (err != nil) != test.Error || test.Expected != actual {
			t.Errorf("Unexpected result from test %d:", i)
			t.Logf("  Version:    '%s'", test.Version)
			t.Logf("  Comparison: '%s'", test.Comparison)
			t.Logf("  Expected:   %v", test.Expected)
			t.Logf("  Actual:     %v", actual)
			t.Logf("  exp err:    %v", test.Error)
			t.Logf("  act err:    %v", err)
		}
	}
}

func TestVersionCompare(t *testing.T) {
	tests := []struct {
		Version    string
		Comparison string
		Expected   []bool
	}{
		{
			Version:    "1.2.3",
			Comparison: "2.0.0",
			Expected:   []bool{true, true, false, false, false},
		},
		{
			Version:    "1.2.3",
			Comparison: "1.2.3",
			Expected:   []bool{false, true, true, true, false},
		},
		{
			Version:    "2.1.0",
			Comparison: "1.2.3",
			Expected:   []bool{false, false, false, true, true},
		},
	}

	count.Tests(t, 5*len(tests))

	for i, test := range tests {
		version, err := semver.Parse(test.Version)
		if err != nil {
			t.Fatal(err)
		}

		comparison, err := semver.Parse(test.Comparison)
		if err != nil {
			t.Fatal(err)
		}

		result := version.Compare(comparison)

		expected := json.LazyLogging(test.Expected)
		actual := json.LazyLogging([]bool{
			result.IsLessThan(),
			result.IsLessOrEqual(),
			result.IsEqualTo(),
			result.IsGreaterOrEqual(),
			result.IsGreaterThan(),
		})

		if expected != actual {
			t.Errorf("One or more result functions are incorrect in test %d", i)
			t.Logf("  Version:    %s", test.Version)
			t.Logf("  Comparison: %s", test.Comparison)
			t.Logf("  Expected:   %s", expected)
			t.Logf("  Actual:     %s", actual)
		}
	}
}
