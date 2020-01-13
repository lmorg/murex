package cmdpipe_test

import (
	"fmt"
	"sync/atomic"
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

var uniqueID int32

func TestPipe(t *testing.T) {
	id := atomic.AddInt32(&uniqueID, 1)

	tests := []test.MurexTest{{
		Block: fmt.Sprintf(`
			pipe: murextest%d

			bg {
				<murextest%d> -> match 2
			}
			out 1 -> <murextest%d>
			out 2 -> <murextest%d>
			out 3 -> <murextest%d>

			sleep 1
			!pipe: murextest%d
		`, id, id, id, id, id, id),
		ExitNum: 0,
		Stdout:  "2\n",
		Stderr:  ``,
	}}

	test.RunMurexTests(tests, t)
}
