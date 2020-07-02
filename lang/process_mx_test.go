package lang_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

func TestMxPipes(t *testing.T) {
	tests := []test.MurexTest{{
		Block: `pipe: foobar
				bg { <foobar> }
				out: "Hello, world!" -> <foobar>
				!pipe: foobar`,
		Stdout: "Hello, world!\n",
	}}

	test.RunMurexTests(tests, t)
}
