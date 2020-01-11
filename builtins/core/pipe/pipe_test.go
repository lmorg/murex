package cmdpipe_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

func TestPipe(t *testing.T) {
	tests := []test.MurexTest{{
		Block: `
		pipe: murextest

		bg {
			<murextest> -> match 2
		}
		out 1 -> <murextest>
		out 2 -> <murextest>
		out 3 -> <murextest>

		!pipe: murextest
		`,
		ExitNum: 0,
		Stdout:  "2\n",
		Stderr:  ``,
	}}

	test.RunMurexTests(tests, t)
}
