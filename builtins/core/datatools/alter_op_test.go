package datatools_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/core/io"
	"github.com/lmorg/murex/test"
)

func TestAlterOp(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `tout json %{a:1, b:2, c:3} ~> %{b:4}`,
			Stdout: `{"a":1,"b":4,"c":3}`,
		},
	}

	test.RunMurexTests(tests, t)
}
