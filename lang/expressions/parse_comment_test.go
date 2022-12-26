package expressions

import (
	"testing"

	"github.com/lmorg/murex/test"
)

func TestParseCommentMultiLineStatement(t *testing.T) {
	tests := []testParseStatementT{
		{
			Statement: `echo /# this is an infixed comment #/ foo bar`,
			Args: []string{
				"echo", "foo", "bar",
			},
			Exec: false,
		},
		{
			Statement: `echo /# this is an infixed comment #/ foo bar`,
			Args: []string{
				"echo", "foo", "bar",
			},
			Exec: true,
		},
		///
		{
			Statement: `echo/# this is an infixed comment #/ foo bar`,
			Args: []string{
				"echo", "foo", "bar",
			},
			Exec: false,
		},
		{
			Statement: `echo/# this is an infixed comment #/ foo bar`,
			Args: []string{
				"echo", "foo", "bar",
			},
			Exec: true,
		},
		///
		{
			Statement: `echo /# this is an infixed comment #/foo bar`,
			Args: []string{
				"echo", "foo", "bar",
			},
			Exec: false,
		},
		{
			Statement: `echo /# this is an infixed comment #/foo bar`,
			Args: []string{
				"echo", "foo", "bar",
			},
			Exec: true,
		},
	}

	testParseStatement(t, tests)
}

func TestParseCommentMultiLine(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `3 /#* 4 #/+ 5`,
			Stdout: `8`,
		},
		{
			Block:  `/# foobar #/ out test`,
			Stdout: "test\n",
		},
		{
			Block:  `%[ 1 2 /#3#/ 4 ]`,
			Stdout: `[1,2,4]`,
		},
		{
			Block:  `%{ a:1, b:2, /#c:3#/ d:4 }`,
			Stdout: `{"a":1,"b":2,"d":4}`,
		},
	}

	test.RunMurexTests(tests, t)
}
