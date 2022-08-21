package arraytools

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/types/generic"
	_ "github.com/lmorg/murex/builtins/types/json"
	_ "github.com/lmorg/murex/builtins/types/jsonlines"
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
