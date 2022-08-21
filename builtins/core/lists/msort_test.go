package lists_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/core/io"
	"github.com/lmorg/murex/test"
)

func TestMsort(t *testing.T) {
	tests := []test.MurexTest{{
		Block:  `tout json ([ "b", "c", "a" ]) -> msort`,
		Stdout: `["a","b","c"]`,
	}}

	test.RunMurexTests(tests, t)
}
