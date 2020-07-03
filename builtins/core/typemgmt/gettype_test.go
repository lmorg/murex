package typemgmt_test

import (
	"os"
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

const envVarPrefix = "MUREX_TEST_VAR_"

func TestGetType(t *testing.T) {
	// these tests don't support multiple counts
	if os.Getenv(envVarPrefix+t.Name()) == "1" {
		return
	}

	err := os.Setenv(envVarPrefix+t.Name(), "1")
	if err != nil {
		t.Fatalf("Aborting test because unable to set env: %s", err)
	}

	tests := []test.MurexTest{
		{
			Block: `
				global: testtype foo=bar
				get-type: \$foo
			`,
			ExitNum: 0,
			Stdout:  "testtype",
			Stderr:  ``,
		},

		{
			Block: `
				function murex_test_gettype {
					get-type: stdin
				}
				tout: testtype foobar -> murex_test_gettype
			`,
			ExitNum: 0,
			Stdout:  "testtype",
			Stderr:  ``,
		},

		{
			Block: `
				pipe: testpipe
				tout testtype foobar -> <testpipe>
				get-type: testpipe
			`,
			ExitNum: 0,
			Stdout:  "testtype",
			Stderr:  ``,
		},
	}

	test.RunMurexTests(tests, t)
}
