package cmdpipe_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

func TestReadPipe(t *testing.T) {
	tests := []test.MurexTest{{
		Block: `
		function murex_test_readpipe {
			echo <null> this is a dummy line
			if { true } then {
				<stdin> -> match 2
			}
		}

		out 1 -> murex_test_readpipe
		out 2 -> murex_test_readpipe
		out 3 -> murex_test_readpipe
		`,
		ExitNum: 0,
		Stdout:  "2\n",
		Stderr:  ``,
	}}

	test.RunMurexTests(tests, t)
}
