package expressions_test

import (
	"testing"

	"github.com/lmorg/murex/test"
)

func TestParseFunction(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `out(hello world)`,
			Stdout: "hello world",
		},
		{
			Block:  `out("hello" 'world')`,
			Stdout: "hello world",
		},
		{
			Block:  `out(%(hello) %(world))`,
			Stdout: "hello world",
		},
		{
			Block:  `out({hello}{world})`,
			Stdout: "{hello}{world}",
		},
		{
			Block:  `if({true}{out hello}{out world})`,
			Stdout: "hello\n",
		},
		{
			Block:  `out(%({hello}{world}))`,
			Stdout: "{hello}{world}",
		},
	}

	test.RunMurexTests(tests, t)
}
