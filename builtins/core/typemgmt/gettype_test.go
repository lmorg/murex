package typemgmt_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

func TestGetType(t *testing.T) {
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
