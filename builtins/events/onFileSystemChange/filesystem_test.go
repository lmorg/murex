package onfilesystemchange_test

import (
	"testing"

	_ "github.com/lmorg/murex/lang/expressions"
	"github.com/lmorg/murex/test"
)

// https://github.com/lmorg/murex/issues/418
func TestIssue418(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:   `event onFileSystemChange foobar= {}`,
			Stderr:  `no path to watch supplied`,
			ExitNum: 1,
		},
	}

	test.RunMurexTestsRx(tests, t)
}
