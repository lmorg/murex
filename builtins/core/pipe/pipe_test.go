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

func TestPipe2(t *testing.T) {
	id := atomic.AddInt32(&uniqueID, 1)

	tests := []test.MurexTest{{
		Block: fmt.Sprintf(`
			pipe: murextest%d

			bg {
				<murextest%d>
			}
			out 1 -> <murextest%d>
			out 2 -> <murextest%d>
			out 3 -> <murextest%d>

			!pipe: murextest%d
		`, id, id, id, id, id, id),
		ExitNum: 0,
		Stdout:  "1\n2\n3\n",
		Stderr:  ``,
	}}

	test.RunMurexTests(tests, t)
}

func TestPipeOrderOfExecution(t *testing.T) {
	// This test might seem counterintuitive with the `pipe` command at the end
	// but `pipe` and `test` builtins in any given block are run before any
	// other function to enable murexes interpreter to create murex named pipes
	// before compiling the pipe stream into any other routine. The purpose of
	// this test is to ensure that behavior still holds true rather than to
	// demonstrate good practice of having `pipe` calls at the end of routines.
	id := atomic.AddInt32(&uniqueID, 1)

	tests := []test.MurexTest{{
		Block: fmt.Sprintf(`
			bg {
				<murextest%d>
			}
			out 1 -> <murextest%d>
			out 2 -> <murextest%d>
			out 3 -> <murextest%d>

			!pipe: murextest%d
			pipe: murextest%d
		`, id, id, id, id, id, id),
		ExitNum: 0,
		Stdout:  "1\n2\n3\n",
		Stderr:  ``,
	}}

	test.RunMurexTests(tests, t)
}
