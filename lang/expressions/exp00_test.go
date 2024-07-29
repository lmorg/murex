package expressions_test

import (
	"testing"

	"github.com/lmorg/murex/test"
)

// https://github.com/lmorg/murex/issues/845
func TestGenericsIssue845(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `tout * "" -> set TestGenericsIssue845_0; $TestGenericsIssue845_0 == "" `,
			Stdout: `true`,
		},
		{
			Block:  `tout randomType "foobar" -> set TestGenericsIssue845_1; $TestGenericsIssue845_1 == "foobar"`,
			Stdout: `true`,
		},
		{
			Block:  `tout json '{"foo": "bar"}' -> set TestGenericsIssue845_2; $TestGenericsIssue845_2 == %{foo:bar}`,
			Stdout: `true`,
		},
		{
			Block:  `tout json '{"c": 3, "b": 2, "a": 1}' -> set TestGenericsIssue845_3; $TestGenericsIssue845_3 == %{c:3, a:1, b:2}`,
			Stdout: `true`,
		},
		{
			Block:  `tout json '[5,3,1]' -> set TestGenericsIssue845_4; $TestGenericsIssue845_4 == %[5 3 1]`,
			Stdout: `true`,
		},
	}

	test.RunMurexTests(tests, t)
}
