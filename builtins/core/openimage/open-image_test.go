package openimage_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

func TestOpenImageError(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:   `trypipe { open-image -> out: foobar }`,
			Stdout:  `^$`,
			Stderr:  `expecting to output to the terminal`,
			ExitNum: 1,
		},
	}

	test.RunMurexTestsRx(tests, t)
}
