package json_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

func TestTableToMap(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `tout csv "a, b, c\n1, 2, 3" -> format json`,
			Stdout: `[{"a":"1","b":"2","c":"3"}]`,
		},
		{
			Block:  `tout * "a, b, c\n1, 2, 3" -> format json`,
			Stdout: `[["a,","b,","c"],["1,","2,","3"]]`,
		},
		{
			Block:  `tout csv "a, b, c\n1, 2, 3" -> format jsonl -> format json`,
			Stdout: `[["a","b","c"],["1","2","3"]]`,
		},
	}

	test.RunMurexTests(tests, t)
}
