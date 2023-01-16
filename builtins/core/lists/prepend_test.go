package lists_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

func TestPrependRegressionBug(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `set v = foo; $v -> prepend bar`,
			Stdout: "bar\nfoo\n",
		},
		{
			Block:  `set v = foo; $v -> append bar`,
			Stdout: "foo\nbar\n",
		},
	}

	test.RunMurexTests(tests, t)
}
