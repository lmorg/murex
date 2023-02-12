package lang_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

func TestBug507(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:   `%{thing:{stuff:chow}} -> ![ thing ] -> set: TestBug507`,
			ExitNum: 0,
		},
	}

	test.RunMurexTestsRx(tests, t)
}
