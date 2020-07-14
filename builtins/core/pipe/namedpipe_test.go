package cmdpipe_test

import (
	"fmt"
	"sync/atomic"
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

func TestReadPipe(t *testing.T) {
	id := atomic.AddInt32(&uniqueID, 1)

	tests := []test.MurexTest{{
		Block: fmt.Sprintf(`
			function murex_test_readpipe_%d {
				echo <null> this is a dummy line
				if { true } then {
					<stdin> -> match 2
				}
			}

			out 1 -> murex_test_readpipe_%d
			out 2 -> murex_test_readpipe_%d
			out 3 -> murex_test_readpipe_%d
		`, id, id, id, id),
		ExitNum: 0,
		Stdout:  "2\n",
		Stderr:  ``,
	}}

	test.RunMurexTests(tests, t)
}
