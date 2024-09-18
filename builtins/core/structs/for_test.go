package structs_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

func TestFor(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `for { $i=1; $i<6; $i=$i+1 } { echo $i }`,
			Stdout: "1\n2\n3\n4\n5\n",
		},
	}

	test.RunMurexTests(tests, t)
}
