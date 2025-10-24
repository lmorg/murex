package yaml_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

func TestTableToMap(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `tout csv "a, b, c\n1, 2, 3" -> format yaml`,
			Stdout: "- a: \"1\"\n  b: \"2\"\n  c: \"3\"\n",
		},
		{
			Block:  `tout * "a, b, c\n1, 2, 3" -> format yaml`,
			Stdout: "- - a,\n  - b,\n  - c\n- - 1,\n  - 2,\n  - \"3\"\n",
		},
	}

	test.RunMurexTests(tests, t)
}
