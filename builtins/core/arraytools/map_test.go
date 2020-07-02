package arraytools_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/core/mkarray"
	"github.com/lmorg/murex/test"
)

func TestMap(t *testing.T) {
	tests := []test.MurexTest{{
		Block:  `map { ja: [1..5] } { a: [Mon..Fri] }`,
		Stdout: `{"1":"Mon","2":"Tue","3":"Wed","4":"Thu","5":"Fri"}`,
	}}

	test.RunMurexTests(tests, t)
}
