package io_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

func TestRead(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `bg { read: q "?" }; sleep 2`,
			Stderr: `background processes cannot read from stdin`,
		},
	}

	test.RunMurexTestsRx(tests, t)
}
