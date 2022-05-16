package onfilesystemchange_test

import (
	"testing"

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
