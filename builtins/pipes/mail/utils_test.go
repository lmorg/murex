package mail

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

type testGetDomainT struct {
	Email    string
	Expected string
	Error    bool
}

func TestGetDomain(t *testing.T) {
	tests := []testGetDomainT{
		{
			Email:    `foo@bar`,
			Expected: `bar`,
		},
		{
			Email:    `foo@bar.com`,
			Expected: `bar.com`,
		},
		{
			Email:    `foo@192.168.1.1`,
			Expected: `192.168.1.1`,
		},
		{
			Email:    `@foo`,
			Expected: `foo`,
			Error:    false,
		},
		{
			Email:    `foo@`,
			Expected: ``,
			Error:    true,
		},
		{
			Email:    `foo`,
			Expected: ``,
			Error:    true,
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		actual, err := getDomain(test.Email)

		if (err == nil) == test.Error {
			t.Errorf("Error mismatch on test %d:", i)
			t.Logf("  Email: '%s'", test.Email)
			t.Logf("  Expected: '%s'", test.Expected)
			t.Logf("  Actual:   '%s'", actual)
			t.Logf("  err exp:   %v", test.Error)
			t.Logf("  err act:   %v", err)
		}

		if test.Expected != actual {
			t.Errorf("domain does not match expected on test %d:", i)
			t.Logf("  Email: '%s'", test.Email)
			t.Logf("  Expected: '%s'", test.Expected)
			t.Logf("  Actual:   '%s'", actual)
			t.Logf("  err exp:   %v", test.Error)
			t.Logf("  err act:   %v", err)
		}
	}
}
