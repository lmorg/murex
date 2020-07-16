package io_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/core/open"
	_ "github.com/lmorg/murex/builtins/core/typemgmt"
	"github.com/lmorg/murex/test"
)

func TestTmp(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				out: "foobar" -> tmp -> set: MUREX_TEST_tmp_cmd
				open: $MUREX_TEST_tmp_cmd`,
			Stdout: "foobar\n",
		},
	}

	test.RunMurexTests(tests, t)
}
