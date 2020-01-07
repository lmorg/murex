package cmdruntime

import (
	"os"
	"testing"

	"github.com/lmorg/murex/test"
	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
)

func TestRangeByIndex(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `runtime --help`,
			Stdout: marshalHelp(t),
		},
	}

	test.RunMurexTests(tests, t)
}

func marshalHelp(t *testing.T) string {
	t.Helper()

	b, err := json.Marshal(help(), false)
	if err != nil {
		t.Errorf("Cannot marshal help(): %s", err)
	}
	return string(b)
}

func TestDumpExports(t *testing.T) {
	namespace := "MUREX_TEST_TestDumpExports_"

	tests := map[string]string{
		"nil":       "",
		"text":      "testing123",
		"s p a c e": "testing 123",
		"equals":    "testing=123",
	}

	count.Tests(t, len(tests))

	for k, v := range tests {
		if err := os.Setenv(namespace+k, v); err != nil {
			t.Error("Unable to set env var for test!")
			t.Logf("  Name:  '%s'", namespace+k)
			t.Logf("  Value: '%s'", v)
			t.Logf("  Error:  %s", err)
		}
	}

	m, err := dumpExports()
	if err != nil {
		t.Error(err)
	}

	for k, v := range tests {
		if m[namespace+k] != v {
			t.Error("Unexpected map value:")
			t.Logf("  Name:     '%s'", namespace+k)
			t.Logf("  Expected: '%s'", v)
			t.Logf("  Actual:   '%s'", m[namespace+k])
		}

		if err := os.Unsetenv(namespace + k); err != nil {
			t.Errorf("Unable to unset '%s': %s", namespace+k, err)
		}
	}
}
