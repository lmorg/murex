package arraytools_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

/*
	ADDHEADING
*/

func TestAddheadingJsonl(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `tout: jsonl '["Bob", 23, true]' -> addheading name age active`,
			Stdout: "[\"name\",\"age\",\"active\"]\n[\"Bob\",\"23\",\"true\"]\n",
		},
	}

	test.RunMurexTests(tests, t)
}
