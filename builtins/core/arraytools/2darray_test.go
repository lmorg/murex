package arraytools_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/core/mkarray"
	"github.com/lmorg/murex/test"
)

func Test2dArray(t *testing.T) {
	tests := []test.MurexTest{{
		Block:  `2darray { a: [1..3] } { ja: [Mon..Wed] }`,
		Stdout: `[["",""],["1","Mon"],["2","Tue"],["3","Wed"]]`,
	}}

	test.RunMurexTests(tests, t)
}
